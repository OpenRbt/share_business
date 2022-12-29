package user

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"wash_bonus/internal/app"
	"wash_bonus/internal/entity"
)

func (s *service) Get(ctx context.Context, auth *app.Auth) (user entity.User, err error) {
	return s.userRepo.Get(ctx, auth.UID)
}
func (s *service) GetByID(ctx context.Context, auth *app.Auth, ID uuid.UUID) (user entity.User, err error) {
	return s.userRepo.GetByID(ctx, ID)
}

func (s *service) UpdateBalance(ctx context.Context, user uuid.UUID, amount decimal.Decimal) (newBalance decimal.Decimal, err error) {
	err = s.userRepo.UpdateBalance(ctx, user, amount)
	newBalance, err = s.userRepo.GetBalance(ctx, user)
	return
}
