package wallets

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

var walletColumns = []string{"w.id", "w.user_id", "w.organization_id", "o.display_name as organization_name", "w.balance", "w.is_default"}

func (r *walletRepo) Get(ctx context.Context, userID string, pagination dbmodels.Pagination) ([]dbmodels.Wallet, error) {
	op := "failed to get wallet: %w"

	var wallets []dbmodels.Wallet
	_, err := r.db.NewSession(nil).
		Select(walletColumns...).
		From(dbr.I("wallets").As("w")).
		Join(dbr.I("organizations").As("o"), "o.id = w.organization_id").
		Where("NOT w.deleted AND w.user_id = ?", userID).
		Limit(uint64(pagination.Limit)).
		Offset(uint64(pagination.Offset)).
		LoadContext(ctx, &wallets)

	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return wallets, err
}

func (r *walletRepo) GetByOrganizationId(ctx context.Context, userID string, organizationID uuid.UUID) (dbmodels.Wallet, error) {
	op := "failed to get wallet by org ID: %w"

	var wallet dbmodels.Wallet
	err := r.db.NewSession(nil).
		Select(walletColumns...).
		From(dbr.I("wallets").As("w")).
		Join(dbr.I("organizations").As("o"), "o.id = w.organization_id").
		Where("NOT w.deleted").
		Where("w.user_id = ? AND w.organization_id = ?", userID, organizationID).
		LoadOneContext(ctx, &wallet)

	if err == nil {
		return wallet, nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.Wallet{}, dbmodels.ErrNotFound
	}

	return wallet, fmt.Errorf(op, err)
}

func (r *walletRepo) Create(ctx context.Context, userID string, organizationID uuid.UUID) (dbmodels.Wallet, error) {
	op := "failed to create wallet: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.Wallet{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	var id uuid.UUID
	err = tx.InsertInto("wallets").
		Columns("user_id", "organization_id").
		Values(userID, organizationID).
		Returning("id").
		LoadContext(ctx, &id)

	if err != nil {
		return dbmodels.Wallet{}, fmt.Errorf(op, err)
	}

	var wallet dbmodels.Wallet
	err = tx.Select(walletColumns...).
		From(dbr.I("wallets").As("w")).
		Join(dbr.I("organizations").As("o"), "o.id = w.organization_id").
		Where("NOT w.deleted").
		Where("w.id = ?", id).
		LoadOneContext(ctx, &wallet)

	if err != nil {
		return dbmodels.Wallet{}, fmt.Errorf(op, err)
	}

	return wallet, tx.Commit()
}

func (r *walletRepo) ChargeBonusesByUserAndOrganization(ctx context.Context, amount decimal.Decimal, userID string, organizationID uuid.UUID) error {
	op := "failed to charge bonuses on wallets: %w"

	if amount.LessThan(decimal.Zero) {
		return dbmodels.ErrBadRequest
	}

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	var (
		wallet             dbmodels.Wallet
		updatedUserBalance decimal.NullDecimal
	)

	err = tx.SelectBySql("SELECT * FROM wallets WHERE user_id = ? AND organization_id = ? FOR UPDATE", userID, organizationID).
		LoadOneContext(ctx, &wallet)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	updatedUserBalance.Decimal = wallet.Balance.Add(amount)
	updatedUserBalance.Valid = true

	_, err = tx.Update("wallets").
		Where("id = ?", wallet.ID).
		Set("balance", updatedUserBalance).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	_, err = tx.InsertInto("balance_events").
		Columns("user", "wallet_id", "old_amount", "new_amount", "date").
		Values(userID, wallet.ID, wallet.Balance, updatedUserBalance, time.Now().UTC()).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}
