package sessions

import (
	"context"
	"errors"
	"fmt"
	"time"
	"washbonus/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

var OrgColumns = []string{"org.id", "org.name", "org.display_name", "org.description", "org.is_default", "org.processing_delay", "org.bonus_percentage", "org.deleted"}

func (r *repo) CreateSession(ctx context.Context, serverID uuid.UUID, postID int64) (dbmodels.Session, error) {
	op := "failed to create session: %w"

	var session dbmodels.Session
	err := r.db.NewSession(nil).
		InsertInto("sessions").
		Columns("wash_server", "post_id").
		Record(dbmodels.CreateSession{
			WashServer: uuid.NullUUID{
				UUID:  serverID,
				Valid: true,
			},
			PostID: postID,
		}).
		Returning("id", "wash_server", "user", "post_id", "started", "finished").
		LoadContext(ctx, &session)

	if err != nil {
		return dbmodels.Session{}, fmt.Errorf(op, err)
	}

	return session, err
}

func (r *repo) GetSession(ctx context.Context, sessionID uuid.UUID) (dbmodels.Session, error) {
	op := "failed to get session: %w"

	var session dbmodels.Session
	err := r.db.NewSession(nil).
		Select("*").
		From("sessions").
		Where("id = ?", uuid.NullUUID{
			UUID:  sessionID,
			Valid: true,
		}).
		LoadOneContext(ctx, &session)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return session, dbmodels.ErrNotFound
		}

		return session, fmt.Errorf(op, err)
	}

	return session, nil
}

func (r *repo) UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state dbmodels.SessionState) error {
	op := "failed to get session: %w"

	updateStmt := r.db.NewSession(nil).
		Update("sessions").
		Where("id = ?", uuid.NullUUID{
			UUID:  sessionID,
			Valid: true,
		})

	switch state {
	case dbmodels.SessionStateStart:
		updateStmt.Set("started", true)
	case dbmodels.SessionStateFinish:
		updateStmt.Set("finished", true)
	}

	_, err := updateStmt.ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return nil
}

func (r *repo) SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) error {
	op := "failed to set session user: %w"

	_, err := r.db.NewSession(nil).
		Update("sessions").
		Set("user", userID).
		Where("id = ?", uuid.NullUUID{
			UUID:  sessionID,
			Valid: true,
		}).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	return nil
}

func (r *repo) SaveMoneyReport(ctx context.Context, report dbmodels.MoneyReport) (err error) {
	op := "failed to save money report: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	var org dbmodels.Organization
	err = tx.Select(OrgColumns...).
		From(dbr.I("organizations").As("org")).
		Join(dbr.I("server_groups").As("gr"), "gr.organization_id = org.id").
		Join(dbr.I("wash_servers").As("ser"), "ser.group_id = gr.id").
		Join(dbr.I("sessions").As("s"), "s.wash_server = ser.id").
		Where("s.id = ?", report.SessionID).
		LoadOneContext(ctx, &org)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	report.OrganizationID = org.ID

	_, err = tx.
		InsertInto("session_money_report").
		Columns("station_id", "banknotes", "cars_total", "coins", "electronical", "service", "session_id", "organization_id", "bonuses", "processed", "uuid").
		Record(report).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func (r *repo) ProcessAndChargeMoneyReports(ctx context.Context, lastId int64) ([]dbmodels.UserMoneyReport, error) {
	op := "failed to process money reports: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.RollbackUnlessCommitted()

	dbReports, err := getUnprocessedReports(ctx, tx, lastId)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	var balanceEvents []dbmodels.UpdateBalanceEvent

	for _, report := range dbReports {
		var wallet dbmodels.Wallet
		err = tx.SelectBySql("SELECT * FROM wallets WHERE user_id = ? AND organization_id = ? FOR UPDATE", report.User, report.OrganizationID).
			LoadOneContext(ctx, &wallet)
		if err != nil {
			return nil, fmt.Errorf(op, err)
		}

		updatedBalance := wallet.Balance.Add(report.PendingBonuses)

		_, err = tx.Update("wallets").
			Set("balance", updatedBalance).
			Where("id = ?", wallet.ID).
			ExecContext(ctx)
		if err != nil {
			return nil, fmt.Errorf(op, err)
		}

		balanceEvents = append(balanceEvents, dbmodels.UpdateBalanceEvent{
			UserID:    report.User,
			WalletID:  wallet.ID,
			OldAmount: wallet.Balance,
			NewAmount: updatedBalance,
		})
	}

	err = insertBalanceEvents(ctx, tx, balanceEvents)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	err = markMoneyReportsAsProcessed(ctx, tx, dbReports)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return dbReports, tx.Commit()
}

func getUnprocessedReports(ctx context.Context, tx *dbr.Tx, lastId int64) ([]dbmodels.UserMoneyReport, error) {
	var dbReports []dbmodels.UserMoneyReport
	_, err := tx.SelectBySql(`
		SELECT 
			r.id, r.station_id, r.banknotes, r.cars_total, r.coins, r.electronical, r.service, r.bonuses, r.session_id, r.organization_id, r.processed, r.uuid, s.user,
			COALESCE(
				(r.banknotes + r.coins + r.electronical) * (COALESCE(sg.bonus_percentage, o.bonus_percentage) / 100.0), 0
			) as pending_bonuses
		FROM 
			session_money_report r
			LEFT JOIN sessions s ON r.session_id = s.id
			JOIN wash_servers w ON s.wash_server = w.id
			JOIN server_groups sg ON w.group_id = sg.id
			JOIN organizations o ON sg.organization_id = o.id OR r.organization_id = o.id
		WHERE 
			r.processed = FALSE 
			AND r.session_id IS NOT NULL
			AND s.user IS NOT NULL
			AND r.id > ?
			AND r.ctime < CURRENT_TIMESTAMP - make_interval(mins => COALESCE(sg.processing_delay, o.processing_delay))
		ORDER BY 
			r.id
		LIMIT 100
	`, lastId).
		LoadContext(ctx, &dbReports)

	if err != nil {
		return nil, err
	}

	return dbReports, nil
}

func insertBalanceEvents(ctx context.Context, tx *dbr.Tx, events []dbmodels.UpdateBalanceEvent) error {
	sqlInsert := `
		INSERT INTO balance_events ("user", wallet_id, old_amount, new_amount, date)
		VALUES ($1, $2, $3, $4, NOW())
	`

	for _, event := range events {
		_, err := tx.ExecContext(ctx, sqlInsert, event.UserID, event.WalletID, event.OldAmount, event.NewAmount)
		if err != nil {
			return err
		}
	}
	return nil
}

func markMoneyReportsAsProcessed(ctx context.Context, tx *dbr.Tx, reports []dbmodels.UserMoneyReport) error {
	if len(reports) < 1 {
		return nil
	}

	reportIDs := make([]int64, len(reports))
	for i, report := range reports {
		reportIDs[i] = report.ID
	}

	_, err := tx.
		Update("session_money_report").
		Set("processed", true).
		Where("id IN ?", reportIDs).
		ExecContext(ctx)

	return err
}

func (r *repo) GetUserPendingBalanceByOrganization(ctx context.Context, userID string, organizationID uuid.UUID) (decimal.Decimal, error) {
	op := "failed to get user pending balance: %w"

	session := r.db.NewSession(nil)
	var pendingBalance decimal.Decimal

	_, err := session.SelectBySql(`
			SELECT COALESCE(SUM((r.banknotes + r.coins + r.electronical) * (COALESCE(sg.bonus_percentage, o.bonus_percentage) / 100.0)), 0) as pending_balance
			FROM session_money_report r
			LEFT JOIN sessions s on r.session_id = s.id
			JOIN wash_servers w ON s.wash_server = w.id
			JOIN server_groups sg ON w.group_id = sg.id
			JOIN organizations o ON sg.organization_id = o.id OR r.organization_id = o.id
			WHERE r.processed = FALSE 
				AND r.session_id is NOT NULL 
				AND s.user = ?
				AND r.organization_id = ?
		`, userID, organizationID).
		LoadContext(ctx, &pendingBalance)

	if err != nil {
		return decimal.Zero, fmt.Errorf(op, err)
	}

	return pendingBalance, nil
}

func (r *repo) GetUserPendingBalances(ctx context.Context, userID string) ([]dbmodels.UserPendingBalance, error) {
	op := "failed to get user pending balances: %w"

	session := r.db.NewSession(nil)
	var pendingBalances []dbmodels.UserPendingBalance

	_, err := session.SelectBySql(`
			SELECT r.organization_id, COALESCE(SUM((r.banknotes + r.coins + r.electronical) * (COALESCE(sg.bonus_percentage, o.bonus_percentage) / 100.0)), 0) as pending_balance
			FROM session_money_report r
			LEFT JOIN sessions s ON r.session_id = s.id
			JOIN wash_servers w ON s.wash_server = w.id
			JOIN server_groups sg ON w.group_id = sg.id
			JOIN organizations o ON sg.organization_id = o.id OR r.organization_id = o.id
			WHERE r.processed = FALSE 
				AND r.session_id IS NOT NULL 
				AND s.user = ?
			GROUP BY r.organization_id
		`, userID).
		LoadContext(ctx, &pendingBalances)

	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return pendingBalances, nil
}

func (r *repo) ChargeBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID, userID string) error {
	op := "failed to charge bonuses: %w"

	if amount.LessThan(decimal.Zero) {
		return dbmodels.ErrBadRequest
	}

	if amount.LessThan(decimal.Zero) {
		return dbmodels.ErrBadRequest
	}

	var (
		sessionBalance decimal.NullDecimal
		orgWallet      dbmodels.Wallet
	)
	dbSessionUUID := uuid.NullUUID{UUID: sessionID, Valid: true}

	date := time.Now()

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	var org dbmodels.Organization
	err = tx.Select(OrgColumns...).
		From(dbr.I("organizations").As("org")).
		Join(dbr.I("server_groups").As("gr"), "gr.organization_id = org.id").
		Join(dbr.I("wash_servers").As("ser"), "ser.group_id = gr.id").
		Join(dbr.I("sessions").As("s"), "s.wash_server = ser.id").
		Where("s.id = ?", sessionID).
		LoadOneContext(ctx, &org)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	//User balance bonus consumption
	err = tx.SelectBySql("SELECT * from wallets WHERE user_id = ? AND organization_id = ? FOR UPDATE", userID, org.ID).
		LoadOneContext(ctx, &orgWallet)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	updatedOrgBalance := orgWallet.Balance.Sub(amount)
	if amount.LessThan(decimal.Zero) || updatedOrgBalance.LessThan(decimal.Zero) {
		return dbmodels.ErrNotEnoughMoney
	}

	_, err = tx.Update("wallets").
		Where("id = ?", orgWallet.ID).
		Set("balance", updatedOrgBalance).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	_, err = tx.InsertInto("balance_events").
		Columns("user", "wallet_id", "old_amount", "new_amount", "date").
		Values(userID, orgWallet.ID, orgWallet.Balance, updatedOrgBalance, date).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	//Session balance bonus assignment
	err = tx.SelectBySql("SELECT balance FROM sessions WHERE id = ? FOR UPDATE", dbSessionUUID).
		LoadOneContext(ctx, &sessionBalance)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	updatedSessionBalance := sessionBalance.Decimal.Add(amount)

	_, err = tx.Update("sessions").
		Where("id = ?", dbSessionUUID).
		Set("balance", updatedSessionBalance).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	_, err = tx.InsertInto("sessions_balance_events").
		Columns("session", "old_amount", "new_amount", "date").
		Values(dbSessionUUID, sessionBalance, updatedSessionBalance, date).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func (r *repo) DiscardBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error) {
	op := "failed to discard bonuses: %w"

	if amount.LessThan(decimal.Zero) {
		return dbmodels.ErrBadRequest
	}

	var (
		userID         *string
		userWallet     dbmodels.Wallet
		sessionBalance decimal.NullDecimal

		org dbmodels.Organization
	)
	dbSessionUUID := uuid.NullUUID{UUID: sessionID, Valid: true}

	date := time.Now()

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	// Receiving user from session
	err = tx.Select("\"user\"").
		From("sessions").
		Where("id = ?", sessionID).
		LoadOneContext(ctx, &userID)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	if userID == nil {
		return dbmodels.ErrNotFound
	}

	// Session balance bonus consumption
	err = tx.SelectBySql("SELECT balance FROM sessions WHERE id = ? FOR UPDATE", dbSessionUUID).
		LoadOneContext(ctx, &sessionBalance)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	updatedSessionBalance := sessionBalance.Decimal.Sub(amount)
	if updatedSessionBalance.LessThan(decimal.Zero) {
		return dbmodels.ErrNotEnoughMoney
	}

	_, err = tx.Update("sessions").
		Where("id = ?", dbSessionUUID).
		Set("balance", updatedSessionBalance).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	_, err = tx.InsertInto("sessions_balance_events").
		Columns("session", "old_amount", "new_amount", "date").
		Values(dbSessionUUID, sessionBalance, updatedSessionBalance, date).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	//User balance bonus assignment

	err = tx.Select(OrgColumns...).
		From(dbr.I("organizations").As("org")).
		Join(dbr.I("server_groups").As("gr"), "gr.organization_id = org.id").
		Join(dbr.I("wash_servers").As("ser"), "ser.group_id = gr.id").
		Join(dbr.I("sessions").As("s"), "s.wash_server = ser.id").
		Where("s.id = ?", sessionID).
		LoadOneContext(ctx, &org)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = tx.SelectBySql("SELECT * FROM users WHERE user_id = ? AND organization_id = ? FOR UPDATE", userID, org.ID).
		LoadOneContext(ctx, &userWallet)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	updatedUserBalance := userWallet.Balance.Add(amount)
	_, err = tx.Update("wallets").
		Where("id = ?", userWallet.ID).
		Set("balance", updatedUserBalance).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	_, err = tx.InsertInto("balance_events").
		Columns("user", "wallet_id", "old_amount", "new_amount", "date").
		Values(userID, userWallet.ID, userWallet.Balance, updatedUserBalance, date).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func (r *repo) ConfirmBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error) {
	op := "failed to confirm bonuses: %w"

	if amount.LessThan(decimal.Zero) {
		return dbmodels.ErrBadRequest
	}

	var (
		userID         string
		sessionBalance decimal.NullDecimal
	)
	dbSessionUUID := uuid.NullUUID{UUID: sessionID, Valid: true}

	date := time.Now()

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	err = tx.Select("user").
		From("sessions").
		Where("id = ?").
		LoadOneContext(ctx, &userID)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	//Session balance bonus consumption
	err = tx.SelectBySql("SELECT balance FROM sessions WHERE id = ? FOR UPDATE", dbSessionUUID).
		LoadOneContext(ctx, &sessionBalance)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	updatedSessionBalance := sessionBalance.Decimal.Sub(amount)
	if updatedSessionBalance.LessThan(decimal.Zero) {
		return dbmodels.ErrNotEnoughMoney
	}

	_, err = tx.Update("sessions").
		Where("id = ?", dbSessionUUID).
		Set("balance", updatedSessionBalance).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	_, err = tx.InsertInto("sessions_balance_events").
		Columns("session", "old_amount", "new_amount", "date").
		Values(dbSessionUUID, sessionBalance, updatedSessionBalance, date).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func (r *repo) LogRewardBonuses(ctx context.Context, sessionID uuid.UUID, payload []byte, messageUuid uuid.UUID) (err error) {
	op := "failed to log bonuses rewarding: %w"

	_, err = r.db.NewSession(nil).
		InsertInto("bonus_reward_log").
		Columns("session_id", "payload", "uuid").
		Values(uuid.NullUUID{UUID: sessionID, Valid: true}, payload, uuid.NullUUID{UUID: messageUuid, Valid: true}).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	return nil
}

func (r *repo) DeleteUnusedSessions(ctx context.Context, SessionRetentionDays int64) (int64, error) {
	op := "failed to delete unused sessions: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	subquery := tx.Select("ses.id").
		From("sessions AS ses").
		Where("ses.user IS NULL").
		Where("ses.id NOT IN (SELECT se.session_id FROM session_money_report AS se)").
		Where("ses.created_at < now() - interval '? days'", SessionRetentionDays)

	_, err = tx.DeleteFrom("session_money_report").
		Where("session_id IN ?", subquery).
		ExecContext(ctx)
	if err != nil {
		return 0, fmt.Errorf(op, err)
	}

	_, err = tx.DeleteFrom("sessions_balance_events").
		Where("session IN ?", subquery).
		ExecContext(ctx)
	if err != nil {
		return 0, fmt.Errorf(op, err)
	}

	_, err = tx.DeleteFrom("bonus_reward_log").
		Where("session_id IN ?", subquery).
		ExecContext(ctx)
	if err != nil {
		return 0, fmt.Errorf(op, err)
	}

	res, err := tx.DeleteFrom("sessions").
		Where("id IN ?", subquery).
		ExecContext(ctx)
	if err != nil {
		return 0, fmt.Errorf(op, err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf(op, err)
	}

	return count, tx.Commit()
}
