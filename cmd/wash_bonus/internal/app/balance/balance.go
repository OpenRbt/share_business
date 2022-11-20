package balance

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"wash_bonus/internal/app"
	"wash_bonus/internal/entity"
)

type Auth app.Auth

type Service interface {
	Get(ctx context.Context, auth Auth, userID uuid.UUID) (decimal.Decimal, error)
	Add(ctx context.Context, auth Auth, userID uuid.UUID, amount uuid.UUID) (decimal.Decimal, error)
	Remove(ctx context.Context, auth Auth, userID uuid.UUID, amount uuid.UUID) (decimal.Decimal, error)
}

type Repository interface {
	GetProfileOrCreateIfNotExists(ctx context.Context, identity string) (entity.User, error)
	Get(ctx context.Context, userID uuid.UUID) (decimal.Decimal, error)
	Add(ctx context.Context, userID uuid.UUID, amount uuid.UUID) (decimal.Decimal, error)
	Remove(ctx context.Context, userID uuid.UUID, amount uuid.UUID) (decimal.Decimal, error)
	LogAction(ctx context.Context, event entity.BalanceEvent) error
}

type service struct {
	l    *zap.SugaredLogger
	repo Repository
}

func New(l *zap.SugaredLogger, repo Repository) *service {
	return &service{
		l:    l,
		repo: repo,
	}
}
