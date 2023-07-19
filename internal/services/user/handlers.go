package user

import (
	"context"
	"errors"
	"washBonus/internal/entity"

	"github.com/shopspring/decimal"
)

func (s *userService) Get(ctx context.Context, userID string) (user entity.User, err error) {
	return s.userRepo.GetByID(ctx, userID)
}

func (s *userService) GetOrCreate(ctx context.Context, userID string) (user entity.User, err error) {
	user, err = s.userRepo.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return s.userRepo.Create(ctx, userID)
		}
	}

	return
}

func (s *userService) AddBonuses(ctx context.Context, amount decimal.Decimal, userID string) (err error) {
	return s.userRepo.AddBonuses(ctx, amount, userID)
}

func (s *userService) Create(ctx context.Context, userID string) (user entity.User, err error) {
	return s.userRepo.Create(ctx, userID)
}

func (s *userService) UpdateUserRole(ctx context.Context, authorizedUserID string, userUpdate entity.UpdateUser) error {
	user, err := s.userRepo.GetByID(ctx, authorizedUserID)
	if err != nil {
		return err
	}

	switch user.Role {
	case entity.AdminRole:
		_, err = s.userRepo.GetByID(ctx, userUpdate.ID)
		if err != nil {
			return err
		}

		err = s.userRepo.UpdateUserRole(ctx, userUpdate)
		if err != nil {
			return err
		}

		return nil
	default:
		return entity.ErrAccessDenied
	}
}
