package session

import (
	"context"
	"wash_bonus/internal/app"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/infrastructure/rabbit/models"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type Service interface {
	AssignRabbit(func(msg interface{}, service string, target string, messageType int) error)
	CreateSession(ctx context.Context, serverID uuid.UUID, postID int64) (session entity.Session, err error)
	CreateSessionPool(ctx context.Context, serverID uuid.UUID, postID int64, sessionsAmount int64) (postSessions models.SessionCreation, err error)
	UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state models.SessionState) error
	GetSession(ctx context.Context, sessionID uuid.UUID) (entity.Session, error)
	GetUserSession(ctx context.Context, auth *app.Auth, sessionID uuid.UUID) (session entity.Session, err error)

	AssignSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error)

	ChargeBonuses(ctx context.Context, sessionID uuid.UUID, userID string, amount decimal.Decimal) (err error)
	ConfirmBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error)
	DiscardBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error)
}

type Repo interface {
	GetSession(ctx context.Context, sessionID uuid.UUID) (entity.Session, error)
	CreateSession(ctx context.Context, serverID uuid.UUID, postID int64) (entity.Session, error)
	UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state models.SessionState) error
	SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error)
	UpdateSessionBalance(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error)
}

type WashRepo interface {
	GetWashServer(ctx context.Context, id uuid.UUID) (entity.WashServer, error)
}

type UserRepo interface {
	GetByID(ctx context.Context, userID string) (user entity.User, err error)
	Create(ctx context.Context, userID string) (user entity.User, err error)
	UpdateBalance(ctx context.Context, userID string, amount decimal.Decimal) (err error)
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
