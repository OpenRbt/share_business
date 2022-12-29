package user

import (
	"context"
	"errors"
	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"time"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/dal"
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
)

func (r *repo) Get(ctx context.Context, identity string) (user entity.User, err error) {
	defer dal.LogOptionalError(r.l, "user", err)

	var dbUser dbmodels.User

	err = r.db.NewSession(nil).
		Select("*").
		From("users").
		Where("identity = ?", identity).
		LoadOneContext(ctx, &dbUser)

	switch {
	case err == nil:
		return conversions.UserFromDb(dbUser), err
	case errors.Is(err, dbr.ErrNotFound):
		user, err = r.Create(ctx, identity)
		return
	default:
		return
	}
}

func (r *repo) GetByID(ctx context.Context, id uuid.UUID) (user entity.User, err error) {
	defer dal.LogOptionalError(r.l, "user", err)

	var dbUser dbmodels.User

	err = r.db.NewSession(nil).
		Select("*").
		From("users").
		Where("id = ?", uuid.NullUUID{UUID: id, Valid: true}).
		LoadOneContext(ctx, &dbUser)

	switch {
	case err == nil:
		return conversions.UserFromDb(dbUser), err
	default:
		return
	}
}

func (r *repo) Create(ctx context.Context, identity string) (user entity.User, err error) {
	defer dal.LogOptionalError(r.l, "user", err)

	var dbUser dbmodels.User

	err = r.db.NewSession(nil).
		InsertInto("users").
		Columns("identity").
		Values(identity).
		Returning("id", "identity", "balance", "active").
		LoadContext(ctx, &dbUser)

	if err != nil {
		return entity.User{}, err
	}

	return conversions.UserFromDb(dbUser), nil
}

func (r *repo) UpdateBalance(ctx context.Context, user uuid.UUID, amount decimal.Decimal) (err error) {
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
		UUID:  user,
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

func (r *repo) GetBalance(ctx context.Context, user uuid.UUID) (balance decimal.Decimal, err error) {
	defer dal.LogOptionalError(r.l, "user", err)

	dbUUID := uuid.NullUUID{
		UUID:  user,
		Valid: true,
	}
	var dbBalance decimal.NullDecimal

	err = r.db.NewSession(nil).
		Select("balance").
		From("users").
		Where("id = ?", dbUUID).
		LoadOneContext(ctx, &dbBalance)
	if err == nil {
		balance = dbBalance.Decimal
	}

	return
}
