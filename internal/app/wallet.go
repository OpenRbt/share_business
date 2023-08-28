package app

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type (
	WalletController interface {
		Get(ctx Ctx, authUser entity.User, pagination entity.Pagination) ([]entity.Wallet, error)
		GetById(ctx Ctx, authUser entity.User, walletID uuid.UUID) (entity.Wallet, error)
		GetByOrganizationId(ctx Ctx, authUser entity.User, organizationID uuid.UUID) (entity.Wallet, error)
	}

	WalletService interface {
		Get(ctx Ctx, userID string, pagination entity.Pagination) ([]entity.Wallet, error)
		GetOrCreate(ctx Ctx, userID string, organizationID uuid.UUID) (entity.Wallet, error)
		GetById(ctx Ctx, walletID uuid.UUID) (entity.Wallet, error)

		ChargeBonusesByUserAndOrganization(ctx Ctx, amount decimal.Decimal, userID string, organizationID uuid.UUID) error
	}

	WalletRepo interface {
		Get(ctx Ctx, userID string, pagination dbmodels.Pagination) ([]dbmodels.Wallet, error)
		GetById(ctx Ctx, walletID uuid.UUID) (dbmodels.Wallet, error)
		GetUserDefaultWallet(ctx Ctx, userID string) (dbmodels.Wallet, error)
		GetByOrganizationId(ctx Ctx, userID string, organizationID uuid.UUID) (dbmodels.Wallet, error)
		Create(ctx Ctx, userID string, organizationID uuid.UUID) (dbmodels.Wallet, error)

		ChargeBonusesByUserAndOrganization(ctx Ctx, amount decimal.Decimal, userID string, organizationID uuid.UUID) error
	}
)
