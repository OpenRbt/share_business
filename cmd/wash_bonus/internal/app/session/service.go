package session

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/entity/vo"
)

type Service interface {
	Create(ctx context.Context, serverID uuid.UUID, postID int64) (session entity.Session, err error)
	Get(ctx context.Context, sessionID uuid.UUID) (entity.Session, error)

	UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state vo.SessionState) error
	SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error)
	UpdateSessionBalance(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error)

	SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error)
	ProcessMoneyReports(ctx context.Context) (err error)
}

type Repo interface {
	GetSession(ctx context.Context, sessionID uuid.UUID) (entity.Session, error)
	CreateSession(ctx context.Context, serverID uuid.UUID, postID int64) (entity.Session, error)

	UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state vo.SessionState) error
	SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error)
	UpdateSessionBalance(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error)

	SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error)
	GetUnprocessedMoneyReports(ctx context.Context) (reports []entity.UserMoneyReport, err error)
	UpdateMoneyReport(ctx context.Context, id int64, processed bool) (err error)
}

type UserRepo interface {
	UpdateBalance(ctx context.Context, userID string, amount decimal.Decimal) (err error)
}

type service struct {
	l           *zap.SugaredLogger
	sessionRepo Repo
	userRepo    UserRepo
}

func New(l *zap.SugaredLogger, userRepo UserRepo, sessionRepo Repo) Service {
	return &service{
		l:           l,
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}
