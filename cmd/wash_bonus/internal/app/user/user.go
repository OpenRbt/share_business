package user

import (
	"context"
	"wash_bonus/internal/entity"

	"github.com/shopspring/decimal"
)

func (s *service) GetByID(ctx context.Context, userID string) (user entity.User, err error) {
	return s.userRepo.GetByID(ctx, userID)
}

func (s *service) UpdateBalance(ctx context.Context, userID string, amount decimal.Decimal) (newBalance decimal.Decimal, err error) {
	err = s.userRepo.UpdateBalance(ctx, userID, amount)
	if err != nil {
		return
	}

	newBalance, err = s.userRepo.GetBalance(ctx, userID)
	return
}
