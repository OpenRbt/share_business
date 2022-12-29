package session

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"wash_bonus/internal/app"
	"wash_bonus/internal/entity"
)

type Service interface {
	CreateSession(ctx context.Context, connectionID uuid.UUID, postID int64, washKey string) (session entity.Session, err error)
	GetSession(ctx context.Context, sessionID uuid.UUID) (session entity.Session, err error)
	RefreshSession(ctx context.Context, sessionID uuid.UUID, PostBalance decimal.Decimal) (session entity.Session, err error)
	EndSession(ctx context.Context, sessionID uuid.UUID) (err error)

	ConsumeMoney(ctx context.Context, sessionID uuid.UUID) (err error)
	AssignUser(ctx context.Context, auth app.Auth, sessionID uuid.UUID) (err error)
}

type WashRepo interface {
	GetWashServerByKey(ctx context.Context, key string) (entity.WashServer, error)
}

type UserRepo interface {
	Get(ctx context.Context, identity string) (user entity.User, err error)
	GetByID(ctx context.Context, id uuid.UUID) (user entity.User, err error)
	UpdateBalance(ctx context.Context, user uuid.UUID, amount decimal.Decimal) (err error)
}

type Cache interface {
	GetSession(sessionID uuid.UUID) *entity.Session
	SetSession(session entity.Session)
	RefreshSession(session entity.Session) error
}

type service struct {
	l        *zap.SugaredLogger
	washRepo WashRepo
	userRepo UserRepo
	cache    Cache
}

func New(l *zap.SugaredLogger, washRepo WashRepo, userRepo UserRepo, cache Cache) *service {
	return &service{
		l:        l,
		washRepo: washRepo,
		userRepo: userRepo,
		cache:    cache,
	}
}
