package user

import (
	"context"
	"wash_bonus/internal/entity"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type Service interface {
	Create(ctx context.Context, userID string) (user entity.User, err error)
	Get(ctx context.Context, userID string) (user entity.User, err error)
	AddBonuses(ctx context.Context, amount decimal.Decimal, userID string) (err error)
}

type Repo interface {
	GetByID(ctx context.Context, userID string) (user entity.User, err error)
	Create(ctx context.Context, userID string) (user entity.User, err error)
	AddBonuses(ctx context.Context, amount decimal.Decimal, userID string) (err error)
	GetBalance(ctx context.Context, userID string) (balance decimal.Decimal, err error)
}

type service struct {
	l        *zap.SugaredLogger
	userRepo Repo
}

func New(l *zap.SugaredLogger, userRepo Repo) *service {
	return &service{
		l:        l,
		userRepo: userRepo,
	}
}
