package user

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"wash_bonus/internal/app"
	"wash_bonus/internal/entity"
)

type Service interface {
	Get(ctx context.Context, auth *app.Auth) (user entity.User, err error)
	GetByID(ctx context.Context, auth *app.Auth, ID uuid.UUID) (user entity.User, err error)

	UpdateBalance(ctx context.Context, user uuid.UUID, amount decimal.Decimal) (newBalance decimal.Decimal, err error)
}

type Repo interface {
	Get(ctx context.Context, identity string) (user entity.User, err error)
	GetByID(ctx context.Context, ID uuid.UUID) (user entity.User, err error)
	UpdateBalance(ctx context.Context, user uuid.UUID, amount decimal.Decimal) (err error)
	GetBalance(ctx context.Context, user uuid.UUID) (balance decimal.Decimal, err error)
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
