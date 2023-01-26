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
	AssignRabbit(func(msg interface{}, service string, target string, messageType int) error)
	CreateSession(ctx context.Context, serverID uuid.UUID, postID int64) (session entity.Session, err error)
	GetSession(ctx context.Context, sessionID uuid.UUID) (entity.Session, error)
	GetUserSession(ctx context.Context, auth *app.Auth, sessionID uuid.UUID) (session entity.Session, err error)

	AssignSessionUser(ctx context.Context, serverID uuid.UUID, sessionID uuid.UUID, user entity.User) (err error)

	ChargeBonuses(ctx context.Context, auth *app.Auth, sessionID uuid.UUID, amount decimal.Decimal) (err error)
	ConfirmBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error)
	DiscardBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error)
}

type Repo interface {
	GetSession(ctx context.Context, sessionID uuid.UUID) (entity.Session, error)
	CreateSession(ctx context.Context, serverID uuid.UUID) (entity.Session, error)
	SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID uuid.UUID) (err error)
	UpdateSessionBalance(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error)
}

type WashRepo interface {
	GetWashServer(ctx context.Context, id uuid.UUID) (entity.WashServer, error)
}

type UserRepo interface {
	Get(ctx context.Context, identity string) (user entity.User, err error)
	GetByID(ctx context.Context, id uuid.UUID) (user entity.User, err error)
	UpdateBalance(ctx context.Context, user uuid.UUID, amount decimal.Decimal) (err error)
}

type service struct {
	l           *zap.SugaredLogger
	sessionRepo Repo
	washRepo    WashRepo
	userRepo    UserRepo

	rabbitPublisherFunc func(msg interface{}, service string, target string, messageType int) error
}

func New(l *zap.SugaredLogger, washRepo WashRepo, userRepo UserRepo, sessionRepo Repo) Service {
	return &service{
		l:           l,
		sessionRepo: sessionRepo,
		washRepo:    washRepo,
		userRepo:    userRepo,
	}
}
