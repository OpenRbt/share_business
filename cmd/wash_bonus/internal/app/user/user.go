package user

import (
	"context"
	"errors"
	"wash_bonus/internal/entity"

	"github.com/shopspring/decimal"
)

func (s *service) Get(ctx context.Context, userID string) (user entity.User, err error) {
	user, err = s.userRepo.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			user, err = s.userRepo.Create(ctx, userID)
		}
	}

	return
}

func (s *service) AddBonuses(ctx context.Context, amount decimal.Decimal, userID string) (err error) {
	return s.userRepo.AddBonuses(ctx, amount, userID)
}

func (s *service) Create(ctx context.Context, userID string) (user entity.User, err error) {
	return s.userRepo.Create(ctx, userID)
}
