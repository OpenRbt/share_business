package wallets

import (
	"context"
	"errors"
	"washBonus/internal/conversions"
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (s *walletService) Get(ctx context.Context, userID string, pagination entity.Pagination) ([]entity.Wallet, error) {
	wallets, err := s.walletRepo.Get(ctx, userID, conversions.PaginationToDB(pagination))
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entity.ErrNotFound
		}
		return []entity.Wallet{}, err
	}

	return conversions.WalletsFromDB(wallets), nil
}

func (s *walletService) GetOrCreate(ctx context.Context, userID string, organizationID uuid.UUID) (entity.Wallet, error) {
	_, err := s.orgRepo.GetById(ctx, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entity.ErrNotFound
		}

		return entity.Wallet{}, err
	}

	wallet, err := s.walletRepo.GetByOrganizationId(ctx, userID, organizationID)
	if errors.Is(err, dbmodels.ErrNotFound) {
		wallet, err = s.walletRepo.Create(ctx, userID, organizationID)
		if err != nil {
			return entity.Wallet{}, err
		}
	}

	if err != nil {
		return entity.Wallet{}, err
	}

	return conversions.WalletFromDB(wallet), nil
}

func (s *walletService) GetById(ctx context.Context, walletID uuid.UUID) (entity.Wallet, error) {
	wallet, err := s.walletRepo.GetById(ctx, walletID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entity.ErrNotFound
		}
		return entity.Wallet{}, err
	}

	return conversions.WalletFromDB(wallet), nil
}

func (s *walletService) ChargeBonusesByUserAndOrganization(ctx context.Context, amount decimal.Decimal, userID string, organizationID uuid.UUID) error {
	err := s.walletRepo.ChargeBonusesByUserAndOrganization(ctx, amount, userID, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrBadValue) {
			return entity.ErrAccessDenied
		}

		return err
	}

	return err
}
