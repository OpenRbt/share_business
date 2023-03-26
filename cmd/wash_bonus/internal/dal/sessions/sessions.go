package sessions

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/dal"
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/infrastructure/rabbit/models"

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

func (r *repo) UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state models.SessionState) (err error) {
	updateStmt := r.db.NewSession(nil).
		Update("sessions").
		Where("id = ?", uuid.NullUUID{
			UUID:  sessionID,
			Valid: true,
		})

	if state == models.SessionStateStart {
		updateStmt.Set("started", true)
	} else if state == models.SessionStateFinish {
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
		dal.LogOptionalError(r.l, "user", err)
		if err != nil && tx != nil {
			err = tx.Rollback()
			dal.LogOptionalError(r.l, "user", err)
		}
	}()

	tx, err = r.db.NewSession(nil).BeginTx(ctx, nil)

	dbUUID := uuid.NullUUID{
		UUID:  sessionID,
		Valid: true,
	}
	dbAmount := decimal.NullDecimal{
		Decimal: amount,
		Valid:   true,
	}
	date := time.Now()
	res := tx.QueryRowContext(ctx, `
	DO
$do$
    DECLARE sessionBalance numeric(10,2);

    BEGIN
        SELECT balance FROM users WHERE id = ? FOR UPDATE INTO sessionBalance;
        if ? < 0 THEN
            IF sessionBalance < abs(?) THEN
                RAISE EXCEPTION ?;
            END IF;
        end if;

        UPDATE sessions SET balance = sessionBalance + ? WHERE id = ?;
        
        INSERT INTO sessions_balance_events(session,old_amount,new_amount,date) values (?, sessionBalance, sessionBalance + ?, ?);
    END;
$do$
`,
		// main args
		dbUUID,
		dbAmount,
		dbAmount,
		entity.ErrNotEnoughMoney.Error(),
		dbAmount,
		dbUUID,
		//logging args
		dbUUID, dbAmount, date,
	)

	if res.Err() != nil {
		switch {
		case errors.Is(err, entity.ErrNotEnoughMoney):
			err = entity.ErrNotEnoughMoney
			tx.Rollback()
			return
		default:
			tx.Rollback()
			return
		}
	}

	err = tx.Commit()

	return
}
