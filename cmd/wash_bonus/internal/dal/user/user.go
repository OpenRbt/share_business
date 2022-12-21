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

func (r *repo) UpdateBalance(ctx context.Context, user uuid.UUID, amount decimal.Decimal) (newBalance decimal.Decimal, err error) {
	var tx *dbr.Tx

	defer func() {
		dal.LogOptionalError(r.l, "user", err)
		if err != nil && tx != nil {
			err = tx.Rollback()
			dal.LogOptionalError(r.l, "user", err)
		}
	}()

	tx, err = r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return
	}

	var oldBalance decimal.NullDecimal
	err = tx.Select("balance").
		From("users").
		Where("id = ?", uuid.NullUUID{UUID: user, Valid: true}).
		LoadOneContext(ctx, &oldBalance)
	if err != nil {
		return
	}

	newBalanceDb := decimal.NullDecimal{Decimal: oldBalance.Decimal.Add(amount), Valid: true}
	err = r.LogBalanceAction(ctx, tx, user, oldBalance, newBalanceDb)
	if err != nil {
		return
	}

	_, err = tx.Update("user").
		Where("id = ?", user.String()).
		Set("balance", newBalanceDb).
		ExecContext(ctx)
	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func (r *repo) LogBalanceAction(ctx context.Context, tx *dbr.Tx, user uuid.UUID, oldAmount decimal.NullDecimal, newAmount decimal.NullDecimal) (err error) {
	_, err = tx.InsertInto("balance_event").
		Columns("user", "old_amount", "new_amount", "date").
		Values(uuid.NullUUID{UUID: user, Valid: true}, oldAmount, newAmount, time.Now().UTC()).
		ExecContext(ctx)

	return
}
