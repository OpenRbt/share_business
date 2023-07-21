package app

import (
	"context"
	"washBonus/internal/entity"
	"washBonus/internal/entity/vo"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type (
	SessionController interface {
		GetSession(ctx context.Context, sessionID uuid.UUID, userID string) (entity.Session, error)
		ChargeBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID, userID string) error
		AssignUserToSession(ctx context.Context, sessionID uuid.UUID, userID string) error
	}

	SessionService interface {
		Create(ctx context.Context, serverID uuid.UUID, postID int64) (session entity.Session, err error)
		Get(ctx context.Context, sessionID uuid.UUID, userID *string) (entity.Session, error)

		UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state vo.SessionState) error
		SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error)

		ChargeBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID, userID string) (err error)
		DiscardBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error)
		ConfirmBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error)
		LogRewardBonuses(ctx context.Context, sessionID uuid.UUID, payload []byte, messageUuid uuid.UUID) (err error)

		SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error)
		ProcessMoneyReports(ctx context.Context) (err error)
		GetUserPendingBalance(ctx context.Context, userID string) (decimal.Decimal, error)
	}

	SessionRepo interface {
		GetSession(ctx context.Context, sessionID uuid.UUID) (entity.Session, error)
		CreateSession(ctx context.Context, serverID uuid.UUID, postID int64) (entity.Session, error)

		UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state vo.SessionState) error
		SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error)

		ChargeBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID, userID string) (err error)
		DiscardBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error)
		ConfirmBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error)

		LogRewardBonuses(ctx context.Context, sessionID uuid.UUID, payload []byte, messageUuid uuid.UUID) (err error)

		SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error)
		GetUnporcessedReportsByUser(ctx context.Context, userID string) ([]entity.UserMoneyReport, error)
		GetUnprocessedMoneyReports(ctx context.Context, lastId int64, olderThenNMinutes int64) (reports []entity.UserMoneyReport, err error)
		UpdateMoneyReport(ctx context.Context, id int64, processed bool) (err error)
	}
)
