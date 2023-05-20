package sessions

import (
	"context"
	"errors"
	"fmt"
	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"time"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/dal"
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/entity/vo"
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
		fmt.Println(err)
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

	return conversions.SessionFromDB(session), err
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
	date := time.Now()

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
		Columns("station_id", "banknotes", "cars_total", "coins", "electronical", "service", "session_id", "bonuses", "processed").
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
select "reports".id, "reports".station_id, "reports".banknotes, "reports".cars_total, "reports".coins, "reports".electronical, "reports".service, "reports".bonuses, "reports".session_id, "reports".processed, "s".user
from session_money_report "reports"
left join sessions "s" on "reports".session_id = "s".id
where "reports".processed = false 
  	and "reports".session_id is not null 
  	and  "s".user is not null  
  	and "reports".id > ? 
    and date_part('minute', now()::timestamp- "reports".ctime) > ?
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
