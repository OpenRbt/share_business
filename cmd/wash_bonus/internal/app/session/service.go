package session

import (
	"context"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/entity/vo"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type Service interface {
	Create(ctx context.Context, serverID uuid.UUID, postID int64) (session entity.Session, err error)
	Get(ctx context.Context, sessionID uuid.UUID) (entity.Session, error)

	UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state vo.SessionState) error
	SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error)
	DeleteUnusedSessions(ctx context.Context, SessionRetentionDays int64) (int64, error)

	ChargeBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID, userID string) (err error)
	DiscardBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error)
	ConfirmBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error)
	LogRewardBonuses(ctx context.Context, sessionID uuid.UUID, payload []byte, messageUuid uuid.UUID) (err error)

	SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error)
	ProcessMoneyReports(ctx context.Context) (err error)
}

type Repo interface {
	GetSession(ctx context.Context, sessionID uuid.UUID) (entity.Session, error)
	CreateSession(ctx context.Context, serverID uuid.UUID, postID int64) (entity.Session, error)

	UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state vo.SessionState) error
	SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error)
	DeleteUnusedSessions(ctx context.Context, SessionRetentionDays int64) (int64, error)

	ChargeBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID, userID string) (err error)
	DiscardBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error)
	ConfirmBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error)
	LogRewardBonuses(ctx context.Context, sessionID uuid.UUID, payload []byte, messageUuid uuid.UUID) (err error)

	SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error)
	GetUnprocessedMoneyReports(ctx context.Context, lastId int64, olderThenNMinutes int64) (reports []entity.UserMoneyReport, err error)
	UpdateMoneyReport(ctx context.Context, id int64, processed bool) (err error)
}

type UserRepo interface {
	AddBonuses(ctx context.Context, amount decimal.Decimal, userID string) (err error)
}

type service struct {
	l           *zap.SugaredLogger
	sessionRepo Repo
	userRepo    UserRepo

	reportsProcessingDelayInMinutes  int64
	moneyReportsRewardPercentDefault int64
}

func New(l *zap.SugaredLogger, userRepo UserRepo, sessionRepo Repo, reportsProcessingDelayInMinutes int64, moneyReportsRewardPercent int64) Service {

	return &service{
		l:                                l,
		sessionRepo:                      sessionRepo,
		userRepo:                         userRepo,
		reportsProcessingDelayInMinutes:  reportsProcessingDelayInMinutes,
		moneyReportsRewardPercentDefault: moneyReportsRewardPercent,
	}
}
