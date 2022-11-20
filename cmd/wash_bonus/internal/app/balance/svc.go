package balance

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"wash_bonus/internal/entity"
)

func (s *service) Get(ctx context.Context, auth Auth, userID uuid.UUID) (decimal.Decimal, error) {
	user, err := s.repo.GetProfileOrCreateIfNotExists(ctx, auth.UID)
	if err != nil {
		return decimal.Decimal{}, err
	}

	if !user.Active {
		return decimal.Decimal{}, entity.ErrProfileInactive
	}

	return user.Balance, nil
}

func (s *service) ChangeBalance(ctx context.Context, userID uuid.UUID, amount decimal.Decimal) (err error) {
	defer func() {
		if err != nil {
		}
	}()

	return
}

func (s *service) Add(ctx context.Context, auth Auth, userID uuid.UUID, amount uuid.UUID) (decimal.Decimal, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Remove(ctx context.Context, auth Auth, userID uuid.UUID, amount uuid.UUID) (decimal.Decimal, error) {
	//TODO implement me
	panic("implement me")
}
