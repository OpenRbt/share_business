package user

import (
	"context"
	"errors"
	"time"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/dal"
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"

	"github.com/gocraft/dbr/v2"
	"github.com/shopspring/decimal"
)

func (r *repo) GetByID(ctx context.Context, userID string) (user entity.User, err error) {
	defer dal.LogOptionalError(r.l, "user", err)

	var dbUser dbmodels.User

	err = r.db.NewSession(nil).
		Select("*").
		From("users").
		Where("id", userID).
		LoadOneContext(ctx, &dbUser)

	switch {
	case err == nil:
		return conversions.UserFromDb(dbUser), err
	default:
		return
	}
}

func (r *repo) Create(ctx context.Context, userID string) (user entity.User, err error) {
	defer dal.LogOptionalError(r.l, "user", err)

	var dbUser dbmodels.User

	err = r.db.NewSession(nil).
		InsertInto("users").
		Columns("id").
		Values(userID).
		Returning("id", "balance", "deleted").
		LoadContext(ctx, &dbUser)

	if err != nil {
		return entity.User{}, err
	}

	return conversions.UserFromDb(dbUser), nil
}

func (r *repo) UpdateBalance(ctx context.Context, userID string, amount decimal.Decimal) (err error) {
	var tx *dbr.Tx

	defer func() {
		dal.LogOptionalError(r.l, "user", err)
		if err != nil && tx != nil {
			err = tx.Rollback()
			dal.LogOptionalError(r.l, "user", err)
		}
	}()

	tx, err = r.db.NewSession(nil).BeginTx(ctx, nil)

	dbAmount := decimal.NullDecimal{
		Decimal: amount,
		Valid:   true,
	}

	date := time.Now()
	res := tx.QueryRowContext(ctx, `
	DO
$do$
    DECLARE userBalance numeric(10,2);

    BEGIN
        SELECT balance FROM users WHERE id = ? FOR UPDATE INTO userBalance;
        if ? < 0 THEN
            IF userBalance < abs(?) THEN
                RAISE EXCEPTION ?;
            END IF;
        end if;

        UPDATE users SET balance = userBalance + ? WHERE id = ?;
        
        INSERT INTO balance_events(user,old_amount,new_amount,date) values (?, userBalance, userBalance + ?, ?);
    END;
$do$
`,
		// main args
		userID,
		dbAmount,
		dbAmount,
		entity.ErrNotEnoughMoney.Error(),
		dbAmount,
		userID,
		//logging args
		userID, dbAmount, date,
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

func (r *repo) GetBalance(ctx context.Context, userID string) (balance decimal.Decimal, err error) {
	defer dal.LogOptionalError(r.l, "user", err)

	var dbBalance decimal.NullDecimal

	err = r.db.NewSession(nil).
		Select("balance").
		From("users").
		Where("id = ?", userID).
		LoadOneContext(ctx, &dbBalance)
	if err == nil {
		balance = dbBalance.Decimal
	}

	return
}
