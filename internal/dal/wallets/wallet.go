package wallets

import (
	"context"
	"errors"
	"time"
	"washBonus/internal/dal"
	"washBonus/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (r *walletRepo) Get(ctx context.Context, userID string, pagination dbmodels.Pagination) ([]dbmodels.Wallet, error) {
	var err error
	defer dal.LogOptionalError(r.l, "wallet", err)

	var wallets []dbmodels.Wallet

	_, err = r.db.NewSession(nil).
		Select("*").
		From("wallets").
		Where("NOT deleted AND user_id = ?", userID).
		Limit(uint64(pagination.Limit)).
		Offset(uint64(pagination.Offset)).
		LoadContext(ctx, &wallets)

	return wallets, err
}

func (r *walletRepo) GetById(ctx context.Context, walletID uuid.UUID) (dbmodels.Wallet, error) {
	var err error
	defer dal.LogOptionalError(r.l, "wallet", err)

	var wallet dbmodels.Wallet

	err = r.db.NewSession(nil).
		Select("*").
		From("wallets").
		Where("NOT deleted").
		Where("id = ?", walletID).
		LoadOneContext(ctx, &wallet)

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.Wallet{}, dbmodels.ErrNotFound
	}

	return wallet, err
}

func (r *walletRepo) GetUserDefaultWallet(ctx context.Context, userID string) (dbmodels.Wallet, error) {
	var err error
	defer dal.LogOptionalError(r.l, "wallet", err)

	var wallet dbmodels.Wallet

	err = r.db.NewSession(nil).
		Select("*").
		From("wallets").
		Where("NOT deleted").
		Where("is_default AND user_id = ?", userID).
		LoadOneContext(ctx, &wallet)

	return wallet, err
}

func (r *walletRepo) GetByOrganizationId(ctx context.Context, userID string, organizationID uuid.UUID) (dbmodels.Wallet, error) {
	var err error
	defer dal.LogOptionalError(r.l, "wallet", err)

	var wallet dbmodels.Wallet

	err = r.db.NewSession(nil).
		Select("*").
		From("wallets").
		Where("NOT deleted").
		Where("user_id = ? AND organization_id = ?", userID, organizationID).
		LoadOneContext(ctx, &wallet)

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.Wallet{}, dbmodels.ErrNotFound
	}

	return wallet, err
}

func (r *walletRepo) Create(ctx context.Context, userID string, organizationID uuid.UUID) (dbmodels.Wallet, error) {
	var err error
	defer dal.LogOptionalError(r.l, "wallet", err)

	var wallet dbmodels.Wallet

	err = r.db.NewSession(nil).
		InsertInto("wallets").
		Columns("user_id", "organization_id").
		Values(userID, organizationID).
		Returning("id", "user_id", "organization_id", "is_default", "balance").
		LoadContext(ctx, &wallet)

	return wallet, err
}

func (r *walletRepo) ChargeBonusesByUserAndOrganization(ctx context.Context, amount decimal.Decimal, userID string, organizationID uuid.UUID) error {
	var err error
	defer dal.LogOptionalError(r.l, "wallet", err)

	if amount.LessThan(decimal.Zero) {
		return dbmodels.ErrBadValue
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
		return err
	}

	updatedUserBalance.Decimal = wallet.Balance.Add(amount)
	updatedUserBalance.Valid = true

	_, err = tx.Update("wallets").
		Where("id = ?", wallet.ID).
		Set("balance", updatedUserBalance).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	_, err = tx.InsertInto("balance_events").
		Columns("user", "wallet_id", "old_amount", "new_amount", "date").
		Values(userID, wallet.ID, wallet.Balance, updatedUserBalance, time.Now().UTC()).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
