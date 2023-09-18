package wallet

import (
	"context"
	"errors"
	"washbonus/internal/conversions"
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (s *walletService) Get(ctx context.Context, userID string, pagination entities.Pagination) ([]entities.Wallet, error) {
	wallets, err := s.walletRepo.Get(ctx, userID, conversions.PaginationToDB(pagination))
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entities.ErrNotFound
		}
		return []entities.Wallet{}, err
	}

	return conversions.WalletsFromDB(wallets), nil
}

func (s *walletService) GetOrCreate(ctx context.Context, userID string, organizationID uuid.UUID) (entities.Wallet, error) {
	_, err := s.orgRepo.GetById(ctx, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entities.ErrNotFound
		}

		return entities.Wallet{}, err
	}

	wallet, err := s.walletRepo.GetByOrganizationId(ctx, userID, organizationID)
	if errors.Is(err, dbmodels.ErrNotFound) {
		wallet, err = s.walletRepo.Create(ctx, userID, organizationID)
		if err != nil {
			return entities.Wallet{}, err
		}
	}

	if err != nil {
		return entities.Wallet{}, err
	}

	return conversions.WalletFromDB(wallet), nil
}

func (s *walletService) ChargeBonusesByUserAndOrganization(ctx context.Context, amount decimal.Decimal, userID string, organizationID uuid.UUID) error {
	err := s.walletRepo.ChargeBonusesByUserAndOrganization(ctx, amount, userID, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrBadRequest) {
			return entities.ErrForbidden
		}

		return err
	}

	return err
}
