package sessions

import (
	"context"
	"errors"
	"time"
	"washBonus/internal/conversions"
	"washBonus/internal/dal"
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/internal/entity/vo"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (r *repo) CreateSession(ctx context.Context, serverID uuid.UUID, postID int64) (entity.Session, error) {
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
		return entity.Session{}, err
	}

	return conversions.SessionFromDB(session), err
}

func (r *repo) GetSession(ctx context.Context, sessionID uuid.UUID) (entity.Session, error) {
	var session dbmodels.Session
	err := r.db.NewSession(nil).
		Select("*").
		From("sessions").
		Where("id = ?", uuid.NullUUID{
			UUID:  sessionID,
			Valid: true,
		}).
		LoadOneContext(ctx, &session)

	switch {
	case err == nil:
		return conversions.SessionFromDB(session), err
	case errors.Is(err, dbr.ErrNotFound):
		return entity.Session{}, entity.ErrNotFound
	default:
		return entity.Session{}, err
	}
}

func (r *repo) UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state vo.SessionState) (err error) {
	updateStmt := r.db.NewSession(nil).
		Update("sessions").
		Where("id = ?", uuid.NullUUID{
			UUID:  sessionID,
			Valid: true,
		})

	switch state {
	case vo.SessionStateStart:
		updateStmt.Set("started", true)
	case vo.SessionStateFinish:
		updateStmt.Set("finished", true)
	}

	_, err = updateStmt.ExecContext(ctx)
	return
}

func (r *repo) SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error) {
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
	var tx *dbr.Tx

	defer func() {
		dal.LogOptionalError(r.l, "session", err)
		if err != nil && tx != nil {
			err = tx.Rollback()
			dal.LogOptionalError(r.l, "session", err)
		}
	}()

	tx, err = r.db.NewSession(nil).BeginTx(ctx, nil)

	dbUUID := uuid.NullUUID{
		UUID:  sessionID,
		Valid: true,
	}
	date := time.Now().UTC()

	var sessionBalance decimal.NullDecimal

	err = tx.SelectBySql("SELECT balance FROM users WHERE id = ? FOR UPDATE", dbUUID).LoadOneContext(ctx, &sessionBalance)

	if amount.LessThan(decimal.Zero) && amount.Add(sessionBalance.Decimal).LessThan(decimal.Zero) {
		err = entity.ErrNotEnoughMoney
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

	err = tx.Commit()

	return
}

func (r *repo) SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error) {
	defer func() {
		dal.LogOptionalError(r.l, "session", err)
	}()

	dbReport := conversions.MoneyReportToDB(report)

	_, err = r.db.NewSession(nil).
		InsertInto("session_money_report").
		Columns("station_id", "banknotes", "cars_total", "coins", "electronical", "service", "session_id", "bonuses", "processed", "uuid").
		Record(dbReport).
		ExecContext(ctx)

	return
}

func (r *repo) UpdateMoneyReport(ctx context.Context, id int64, processed bool) (err error) {
	defer func() {
		dal.LogOptionalError(r.l, "session", err)
	}()

	_, err = r.db.NewSession(nil).
		Update("session_money_report").
		Set("processed", processed).
		Where("id = ?", id).
		ExecContext(ctx)

	return
}

func (r *repo) GetUnprocessedMoneyReports(ctx context.Context, lastId int64, olderThenNMinutes int64) (reports []entity.UserMoneyReport, err error) {
	defer func() {
		dal.LogOptionalError(r.l, "session", err)
	}()

	var dbReports []dbmodels.UserMoneyReport

	_, err = r.db.NewSession(nil).
		SelectBySql(`
			select "reports".id, "reports".station_id, "reports".banknotes, "reports".cars_total, "reports".coins, "reports".electronical, "reports".service, "reports".bonuses, "reports".session_id, "reports".processed, "reports".uuid,"s".user
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
		if errors.Is(err, dbr.ErrNotFound) {
			err = nil
			return
		}

		return
	}

	reports = conversions.UserMoneyReportsFromDB(dbReports)

	return
}

func (r *repo) ChargeBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID, userID string) (err error) {
	var tx *dbr.Tx

	defer func() {
		dal.LogOptionalError(r.l, "session", err)
		if err != nil && tx != nil {
			err = tx.Rollback()
			dal.LogOptionalError(r.l, "session", err)
		}
	}()

	if amount.LessThan(decimal.Zero) {
		return entity.ErrBadValue
	}

	if amount.LessThan(decimal.Zero) {
		return entity.ErrBadValue
	}

	var (
		userBalance    decimal.NullDecimal
		sessionBalance decimal.NullDecimal
	)
	dbSessionUUID := uuid.NullUUID{UUID: sessionID, Valid: true}

	date := time.Now()

	tx, err = r.db.NewSession(nil).BeginTx(ctx, nil)

	//User balance bonus consumption

	err = tx.SelectBySql("SELECT balance from users WHERE id = ? FOR UPDATE", userID).
		LoadOneContext(ctx, &userBalance)
	if err != nil {
		return
	}

	updatedUserBalance := userBalance.Decimal.Sub(amount)
	if updatedUserBalance.LessThan(decimal.Zero) {
		err = entity.ErrNotEnoughMoney
		return
	}

	_, err = tx.Update("users").
		Where("id = ?", userID).
		Set("balance", decimal.NullDecimal{Decimal: updatedUserBalance, Valid: true}).
		ExecContext(ctx)
	if err != nil {
		return
	}

	_, err = tx.InsertInto("balance_events").
		Columns("user", "old_amount", "new_amount", "date").
		Values(userID, userBalance, updatedUserBalance, date).
		ExecContext(ctx)
	if err != nil {
		return
	}

	//Session balance bonus assignment
	err = tx.SelectBySql("SELECT balance FROM sessions WHERE id = ? FOR UPDATE", dbSessionUUID).
		LoadOneContext(ctx, &sessionBalance)
	if err != nil {
		return
	}

	updatedSessionBalance := sessionBalance.Decimal.Add(amount)

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

	err = tx.Commit()

	return
}

func (r *repo) DiscardBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error) {
	var tx *dbr.Tx

	defer func() {
		dal.LogOptionalError(r.l, "session", err)
		if err != nil && tx != nil {
			err = tx.Rollback()
			dal.LogOptionalError(r.l, "session", err)
		}
	}()

	if amount.LessThan(decimal.Zero) {
		return entity.ErrBadValue
	}

	var (
		userID         *string
		userBalance    decimal.NullDecimal
		sessionBalance decimal.NullDecimal
	)
	dbSessionUUID := uuid.NullUUID{UUID: sessionID, Valid: true}

	date := time.Now()

	tx, err = r.db.NewSession(nil).BeginTx(ctx, nil)

	// Receiving user from session
	err = tx.Select("\"user\"").
		From("sessions").
		Where("id = ?", sessionID).
		LoadOneContext(ctx, &userID)

	if err != nil {
		return
	}

	if userID == nil {
		return entity.ErrNotFound
	}

	// Session balance bonus consumption
	err = tx.SelectBySql("SELECT balance FROM sessions WHERE id = ? FOR UPDATE", dbSessionUUID).
		LoadOneContext(ctx, &sessionBalance)
	if err != nil {
		return
	}

	updatedSessionBalance := sessionBalance.Decimal.Sub(amount)
	if updatedSessionBalance.LessThan(decimal.Zero) {
		err = entity.ErrNotEnoughMoney
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

	err = tx.SelectBySql("SELECT balance from users WHERE id = ? FOR UPDATE", userID).
		LoadOneContext(ctx, &userBalance)
	if err != nil {
		return
	}

	updatedUserBalance := userBalance.Decimal.Add(amount)

	_, err = tx.Update("users").
		Where("id = ?", userID).
		Set("balance", decimal.NullDecimal{Decimal: updatedUserBalance, Valid: true}).
		ExecContext(ctx)
	if err != nil {
		return
	}

	_, err = tx.InsertInto("balance_events").
		Columns("user", "old_amount", "new_amount", "date").
		Values(userID, userBalance, updatedUserBalance, date).
		ExecContext(ctx)
	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func (r *repo) ConfirmBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error) {
	var tx *dbr.Tx

	defer func() {
		dal.LogOptionalError(r.l, "session", err)
		if err != nil && tx != nil {
			err = tx.Rollback()
			dal.LogOptionalError(r.l, "session", err)
		}
	}()

	if amount.LessThan(decimal.Zero) {
		return entity.ErrBadValue
	}

	var (
		userID         string
		sessionBalance decimal.NullDecimal
	)
	dbSessionUUID := uuid.NullUUID{UUID: sessionID, Valid: true}

	date := time.Now()

	tx, err = r.db.NewSession(nil).BeginTx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

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
		err = entity.ErrNotEnoughMoney
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

	err = tx.Commit()

	return
}
func (r *repo) LogRewardBonuses(ctx context.Context, sessionID uuid.UUID, payload []byte, messageUuid uuid.UUID) (err error) {
	defer func() {
		dal.LogOptionalError(r.l, "session", err)
	}()
	_, err = r.db.NewSession(nil).
		InsertInto("bonus_reward_log").
		Columns("session_id", "payload", "uuid").
		Values(uuid.NullUUID{UUID: sessionID, Valid: true}, payload, uuid.NullUUID{UUID: messageUuid, Valid: true}).
		ExecContext(ctx)
	return
}
