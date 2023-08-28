package sessions

import (
	"context"
	"errors"
	"fmt"
	"time"
	"washBonus/internal/dal"
	"washBonus/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (r *repo) CreateSession(ctx context.Context, serverID uuid.UUID, postID int64) (dbmodels.Session, error) {
	var err error
	defer dal.LogOptionalError(r.l, "session", err)

	var session dbmodels.Session

	err = r.db.NewSession(nil).
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
		return dbmodels.Session{}, err
	}

	return session, err
}

func (r *repo) GetSession(ctx context.Context, sessionID uuid.UUID) (dbmodels.Session, error) {
	var err error
	defer dal.LogOptionalError(r.l, "session", err)

	var session dbmodels.Session
	err = r.db.NewSession(nil).
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

		return session, err
	}

	return session, err
}

func (r *repo) UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state dbmodels.SessionState) (err error) {
	defer dal.LogOptionalError(r.l, "session", err)

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

	_, err = updateStmt.ExecContext(ctx)
	return
}

func (r *repo) SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error) {
	defer dal.LogOptionalError(r.l, "session", err)

	updateStmt := r.db.NewSession(nil).
		Update("sessions").
		Where("id = ?", uuid.NullUUID{
			UUID:  sessionID,
			Valid: true,
		})

	updateStmt.Set("user", userID)

	_, err = updateStmt.ExecContext(ctx)
	return
}

func (r *repo) UpdateSessionBalance(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error) {
	defer dal.LogOptionalError(r.l, "session", err)

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	dbUUID := uuid.NullUUID{
		UUID:  sessionID,
		Valid: true,
	}
	date := time.Now().UTC()

	var sessionBalance decimal.NullDecimal

	err = tx.SelectBySql("SELECT balance FROM users WHERE id = ? FOR UPDATE", dbUUID).LoadOneContext(ctx, &sessionBalance)

	if amount.LessThan(decimal.Zero) && amount.Add(sessionBalance.Decimal).LessThan(decimal.Zero) {
		err = dbmodels.ErrNotEnoughMoney
		return
	}

	newDbAmount := decimal.NullDecimal{
		Decimal: sessionBalance.Decimal.Add(amount),
		Valid:   true,
	}

	_, err = tx.Update("sessions").
		Where("id = ?", dbUUID).
		Set("balance", newDbAmount).
		ExecContext(ctx)

	_, err = tx.InsertInto("sessions_balance_events").
		Columns("session", "old_amount", "new_amount", "date").
		Values(dbUUID, sessionBalance, newDbAmount, date).
		ExecContext(ctx)
	if err != nil {
		return
	}

	return tx.Commit()
}

func (r *repo) SaveMoneyReport(ctx context.Context, report dbmodels.MoneyReport) (err error) {
	defer dal.LogOptionalError(r.l, "session", err)

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	var org dbmodels.Organization
	err = tx.Select("org.*").
		From(dbr.I("organizations").As("org")).
		Join(dbr.I("server_groups").As("gr"), "gr.organization_id = org.id").
		Join(dbr.I("wash_servers").As("ser"), "ser.group_id = gr.id").
		Join(dbr.I("sessions").As("s"), "s.wash_server = ser.id").
		Where("s.id = ?", report.SessionID).
		LoadOneContext(ctx, &org)
	if err != nil {
		return err
	}

	report.OrganizationID = org.ID

	_, err = tx.
		InsertInto("session_money_report").
		Columns("station_id", "banknotes", "cars_total", "coins", "electronical", "service", "session_id", "organization_id", "bonuses", "processed", "uuid").
		Record(report).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repo) UpdateMoneyReport(ctx context.Context, id int64, processed bool) (err error) {
	defer dal.LogOptionalError(r.l, "session", err)

	_, err = r.db.NewSession(nil).
		Update("session_money_report").
		Set("processed", processed).
		Where("id = ?", id).
		ExecContext(ctx)

	return
}

func (r *repo) GetUnprocessedMoneyReports(ctx context.Context, lastId int64, olderThenNMinutes int64) (reports []dbmodels.UserMoneyReport, err error) {
	defer dal.LogOptionalError(r.l, "session", err)

	var dbReports []dbmodels.UserMoneyReport

	_, err = r.db.NewSession(nil).
		SelectBySql(`
			select "reports".id, "reports".station_id, "reports".banknotes, "reports".cars_total, "reports".coins, "reports".electronical, "reports".service, "reports".bonuses, "reports".session_id, "reports".organization_id, "reports".processed, "reports".uuid,"s".user
			from session_money_report "reports"
			left join sessions "s" on "reports".session_id = "s".id
			where "reports".processed = false 
				and "reports".session_id is not null 
				and  "s".user is not null  
				and "reports".id > ? 
				and "reports".ctime < now() - interval '? minutes'
			order by "reports".id
			limit 100
		`, lastId, olderThenNMinutes).
		LoadContext(ctx, &dbReports)

	if err != nil {
		return
	}

	return dbReports, nil
}

func (r *repo) GetUnporcessedReportsByUserAndOrganization(ctx context.Context, userID string, organizationID uuid.UUID) ([]dbmodels.UserMoneyReport, error) {
	var err error
	defer dal.LogOptionalError(r.l, "session", err)

	session := r.db.NewSession(nil)
	var dbReports []dbmodels.UserMoneyReport

	_, err = session.SelectBySql(`
			select "reports".id, "reports".station_id, "reports".banknotes, "reports".cars_total, "reports".coins, "reports".electronical, "reports".service, "reports".bonuses, "reports".session_id, "reports".organization_id, "reports".processed, "reports".uuid,"s".user
			from session_money_report "reports"
			left join sessions "s" on "reports".session_id = "s".id
			where "reports".processed = false 
				and "reports".session_id is not null 
				and "s".user = ?
				and "reports".organization_id = ? 
		`, userID, organizationID).
		LoadContext(ctx, &dbReports)

	if err != nil {
		return []dbmodels.UserMoneyReport{}, nil
	}

	return dbReports, nil
}

func (r *repo) ChargeBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID, userID string) error {
	var err error
	defer dal.LogOptionalError(r.l, "session", err)

	if amount.LessThan(decimal.Zero) {
		return dbmodels.ErrBadValue
	}

	if amount.LessThan(decimal.Zero) {
		return dbmodels.ErrBadValue
	}

	var (
		sessionBalance decimal.NullDecimal

		defaultWallet dbmodels.Wallet
		orgWallet     dbmodels.Wallet
	)
	dbSessionUUID := uuid.NullUUID{UUID: sessionID, Valid: true}

	date := time.Now()

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	var org dbmodels.Organization
	err = tx.Select("org.*").
		From(dbr.I("organizations").As("org")).
		Join(dbr.I("server_groups").As("gr"), "gr.organization_id = org.id").
		Join(dbr.I("wash_servers").As("ser"), "ser.group_id = gr.id").
		Join(dbr.I("sessions").As("s"), "s.wash_server = ser.id").
		Where("s.id = ?", sessionID).
		LoadOneContext(ctx, &org)
	if err != nil {
		return err
	}

	//User balance bonus consumption

	err = tx.SelectBySql("SELECT * from wallets WHERE user_id = ? AND is_default FOR UPDATE", userID).
		LoadOneContext(ctx, &defaultWallet)
	if err != nil {
		return err
	}

	err = tx.SelectBySql("SELECT * from wallets WHERE user_id = ? AND organization_id = ? FOR UPDATE", userID, org.ID).
		LoadOneContext(ctx, &orgWallet)
	if err != nil {
		return err
	}

	updatedDefaultBalance := defaultWallet.Balance.Sub(amount)
	updatedOrgBalance := orgWallet.Balance
	if updatedDefaultBalance.LessThan(decimal.Zero) {
		if orgWallet.IsDefault {
			return dbmodels.ErrNotEnoughMoney
		}

		updatedDefaultBalance = decimal.Zero
		updatedOrgBalance = updatedOrgBalance.Add(updatedDefaultBalance)
		if updatedOrgBalance.LessThan(decimal.Zero) {
			return dbmodels.ErrNotEnoughMoney
		}
	}

	fmt.Println(updatedDefaultBalance, updatedOrgBalance)

	_, err = tx.Update("wallets").
		Where("id = ?", defaultWallet.ID).
		Set("balance", updatedDefaultBalance).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	_, err = tx.InsertInto("balance_events").
		Columns("user", "wallet_id", "old_amount", "new_amount", "date").
		Values(userID, defaultWallet.ID, defaultWallet.Balance, updatedDefaultBalance, date).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	fmt.Println(orgWallet, defaultWallet)

	if !orgWallet.IsDefault {
		fmt.Println("Suuka")
		_, err = tx.Update("wallets").
			Where("id = ?", orgWallet.ID).
			Set("balance", updatedOrgBalance).
			ExecContext(ctx)
		if err != nil {
			return err
		}

		_, err = tx.InsertInto("balance_events").
			Columns("user", "wallet_id", "old_amount", "new_amount", "date").
			Values(userID, orgWallet.ID, orgWallet.Balance, updatedOrgBalance, date).
			ExecContext(ctx)
		if err != nil {
			return err
		}
	}

	//Session balance bonus assignment
	err = tx.SelectBySql("SELECT balance FROM sessions WHERE id = ? FOR UPDATE", dbSessionUUID).
		LoadOneContext(ctx, &sessionBalance)
	if err != nil {
		return err
	}

	updatedSessionBalance := sessionBalance.Decimal.Add(amount)

	_, err = tx.Update("sessions").
		Where("id = ?", dbSessionUUID).
		Set("balance", updatedSessionBalance).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	_, err = tx.InsertInto("sessions_balance_events").
		Columns("session", "old_amount", "new_amount", "date").
		Values(dbSessionUUID, sessionBalance, updatedSessionBalance, date).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repo) DiscardBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error) {
	defer dal.LogOptionalError(r.l, "session", err)

	if amount.LessThan(decimal.Zero) {
		return dbmodels.ErrBadValue
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
		return
	}
	defer tx.RollbackUnlessCommitted()

	// Receiving user from session
	err = tx.Select("\"user\"").
		From("sessions").
		Where("id = ?", sessionID).
		LoadOneContext(ctx, &userID)

	if err != nil {
		return
	}

	if userID == nil {
		return dbmodels.ErrNotFound
	}

	// Session balance bonus consumption
	err = tx.SelectBySql("SELECT balance FROM sessions WHERE id = ? FOR UPDATE", dbSessionUUID).
		LoadOneContext(ctx, &sessionBalance)
	if err != nil {
		return
	}

	updatedSessionBalance := sessionBalance.Decimal.Sub(amount)
	if updatedSessionBalance.LessThan(decimal.Zero) {
		err = dbmodels.ErrNotEnoughMoney
		return
	}

	_, err = tx.Update("sessions").
		Where("id = ?", dbSessionUUID).
		Set("balance", updatedSessionBalance).
		ExecContext(ctx)
	if err != nil {
		return
	}

	_, err = tx.InsertInto("sessions_balance_events").
		Columns("session", "old_amount", "new_amount", "date").
		Values(dbSessionUUID, sessionBalance, updatedSessionBalance, date).
		ExecContext(ctx)
	if err != nil {
		return
	}

	//User balance bonus assignment

	err = tx.Select("org.*").
		From(dbr.I("organizations").As("org")).
		Join(dbr.I("server_groups").As("gr"), "gr.organization_id = org.id").
		Join(dbr.I("wash_servers").As("ser"), "ser.group_id = gr.id").
		Join(dbr.I("sessions").As("s"), "s.wash_server = ser.id").
		Where("s.id = ?", sessionID).
		LoadOneContext(ctx, &org)
	if err != nil {
		return err
	}

	err = tx.SelectBySql("SELECT * FROM users WHERE user_id = ? AND organization_id = ? FOR UPDATE", userID, org.ID).
		LoadOneContext(ctx, &userWallet)
	if err != nil {
		return
	}

	updatedUserBalance := userWallet.Balance.Add(amount)

	_, err = tx.Update("wallets").
		Where("id = ?", userWallet.ID).
		Set("balance", updatedUserBalance).
		ExecContext(ctx)
	if err != nil {
		return
	}

	_, err = tx.InsertInto("balance_events").
		Columns("user", "wallet_id", "old_amount", "new_amount", "date").
		Values(userID, userWallet.ID, userWallet.Balance, updatedUserBalance, date).
		ExecContext(ctx)
	if err != nil {
		return
	}

	return tx.Commit()
}

func (r *repo) ConfirmBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error) {
	defer dal.LogOptionalError(r.l, "session", err)

	if amount.LessThan(decimal.Zero) {
		return dbmodels.ErrBadValue
	}

	var (
		userID         string
		sessionBalance decimal.NullDecimal
	)
	dbSessionUUID := uuid.NullUUID{UUID: sessionID, Valid: true}

	date := time.Now()

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.RollbackUnlessCommitted()

	err = tx.Select("user").
		From("sessions").
		Where("id = ?").
		LoadOneContext(ctx, &userID)

	//Session balance bonus consumption
	err = tx.SelectBySql("SELECT balance FROM sessions WHERE id = ? FOR UPDATE", dbSessionUUID).
		LoadOneContext(ctx, &sessionBalance)
	if err != nil {
		return
	}

	updatedSessionBalance := sessionBalance.Decimal.Sub(amount)
	if updatedSessionBalance.LessThan(decimal.Zero) {
		err = dbmodels.ErrNotEnoughMoney
		return
	}

	_, err = tx.Update("sessions").
		Where("id = ?", dbSessionUUID).
		Set("balance", updatedSessionBalance).
		ExecContext(ctx)
	if err != nil {
		return
	}

	_, err = tx.InsertInto("sessions_balance_events").
		Columns("session", "old_amount", "new_amount", "date").
		Values(dbSessionUUID, sessionBalance, updatedSessionBalance, date).
		ExecContext(ctx)
	if err != nil {
		return
	}

	return tx.Commit()
}

func (r *repo) LogRewardBonuses(ctx context.Context, sessionID uuid.UUID, payload []byte, messageUuid uuid.UUID) (err error) {
	defer dal.LogOptionalError(r.l, "session", err)

	_, err = r.db.NewSession(nil).
		InsertInto("bonus_reward_log").
		Columns("session_id", "payload", "uuid").
		Values(uuid.NullUUID{UUID: sessionID, Valid: true}, payload, uuid.NullUUID{UUID: messageUuid, Valid: true}).
		ExecContext(ctx)
	return
}

func (r *repo) DeleteUnusedSessions(ctx context.Context, SessionRetentionDays int64) (int64, error) {
	var err error
	defer func() {
		dal.LogOptionalError(r.l, "session", err)
	}()

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return 0, err
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
		return 0, err
	}

	_, err = tx.DeleteFrom("sessions_balance_events").
		Where("session IN ?", subquery).
		ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	_, err = tx.DeleteFrom("bonus_reward_log").
		Where("session_id IN ?", subquery).
		ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	res, err := tx.DeleteFrom("sessions").
		Where("id IN ?", subquery).
		ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, tx.Commit()
}
