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
		Where("id = ?", userID).
		LoadOneContext(ctx, &dbUser)

	switch {
	case err == nil:
		return conversions.UserFromDb(dbUser), err
	case errors.Is(err, dbr.ErrNotFound):
		err = entity.ErrNotFound
		return
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

func (r *repo) AddBonuses(ctx context.Context, amount decimal.Decimal, userID string) (err error) {
	var tx *dbr.Tx

	defer func() {
		dal.LogOptionalError(r.l, "user", err)
		if err != nil && tx != nil {
			err = tx.Rollback()
			dal.LogOptionalError(r.l, "user", err)
		}
	}()

	if amount.LessThan(decimal.Zero) {
		return entity.ErrBadValue
	}

	tx, err = r.db.NewSession(nil).BeginTx(ctx, nil)

	date := time.Now()

	var (
		userBalance        decimal.NullDecimal
		updatedUserBalance decimal.NullDecimal
	)

	err = tx.SelectBySql("SELECT balance FROM users WHERE  id = ? FOR UPDATE", userID).LoadOneContext(ctx, &userBalance)
	if err != nil {
		return
	}

	updatedUserBalance.Decimal = userBalance.Decimal.Add(amount)
	updatedUserBalance.Valid = true

	_, err = tx.Update("users").
		Where("id = ?", userID).
		Set("balance", updatedUserBalance).
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
