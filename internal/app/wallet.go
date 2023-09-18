package app

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type (
	WalletController interface {
		Get(ctx Ctx, auth Auth, pagination entities.Pagination) ([]entities.Wallet, error)
		GetByOrganizationId(ctx Ctx, auth Auth, organizationID uuid.UUID) (entities.Wallet, error)
	}

	WalletService interface {
		Get(ctx Ctx, userID string, pagination entities.Pagination) ([]entities.Wallet, error)
		GetOrCreate(ctx Ctx, userID string, organizationID uuid.UUID) (entities.Wallet, error)

		ChargeBonusesByUserAndOrganization(ctx Ctx, amount decimal.Decimal, userID string, organizationID uuid.UUID) error
	}

	WalletRepo interface {
		Get(ctx Ctx, userID string, pagination dbmodels.Pagination) ([]dbmodels.Wallet, error)
		GetByOrganizationId(ctx Ctx, userID string, organizationID uuid.UUID) (dbmodels.Wallet, error)
		Create(ctx Ctx, userID string, organizationID uuid.UUID) (dbmodels.Wallet, error)

		ChargeBonusesByUserAndOrganization(ctx Ctx, amount decimal.Decimal, userID string, organizationID uuid.UUID) error
	}
)
