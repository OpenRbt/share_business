package app

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/internal/entity/vo"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type (
	SessionController interface {
		GetSession(ctx Ctx, auth Auth, sessionID uuid.UUID) (entity.Session, error)
		ChargeBonuses(ctx Ctx, auth Auth, amount decimal.Decimal, sessionID uuid.UUID) error
		AssignUserToSession(ctx Ctx, auth Auth, sessionID uuid.UUID) error
	}

	SessionService interface {
		Create(ctx Ctx, serverID uuid.UUID, postID int64) (session entity.Session, err error)
		Get(ctx Ctx, sessionID uuid.UUID, userID *string) (entity.Session, error)

		UpdateSessionState(ctx Ctx, sessionID uuid.UUID, state vo.SessionState) error
		SetSessionUser(ctx Ctx, sessionID uuid.UUID, userID string) (err error)

		ChargeBonuses(ctx Ctx, amount decimal.Decimal, sessionID uuid.UUID, userID string) (err error)
		DiscardBonuses(ctx Ctx, amount decimal.Decimal, sessionID uuid.UUID) (err error)
		ConfirmBonuses(ctx Ctx, amount decimal.Decimal, sessionID uuid.UUID) (err error)
		LogRewardBonuses(ctx Ctx, sessionID uuid.UUID, payload []byte, messageUuid uuid.UUID) (err error)

		SaveMoneyReport(ctx Ctx, report entity.MoneyReport) (err error)
		DeleteUnusedSessions(ctx Ctx, SessionRetentionDays int64) (int64, error)
		ProcessMoneyReports(ctx Ctx) (err error)
		GetUserOrganizationPendingBalance(ctx Ctx, userID string, organizationID uuid.UUID) (decimal.Decimal, error)
	}

	SessionRepo interface {
		GetSession(ctx Ctx, sessionID uuid.UUID) (dbmodels.Session, error)
		CreateSession(ctx Ctx, serverID uuid.UUID, postID int64) (dbmodels.Session, error)

		UpdateSessionState(ctx Ctx, sessionID uuid.UUID, state dbmodels.SessionState) error
		SetSessionUser(ctx Ctx, sessionID uuid.UUID, userID string) (err error)

		ChargeBonuses(ctx Ctx, amount decimal.Decimal, sessionID uuid.UUID, userID string) (err error)
		DiscardBonuses(ctx Ctx, amount decimal.Decimal, sessionID uuid.UUID) (err error)
		ConfirmBonuses(ctx Ctx, amount decimal.Decimal, sessionID uuid.UUID) (err error)

		LogRewardBonuses(ctx Ctx, sessionID uuid.UUID, payload []byte, messageUuid uuid.UUID) (err error)

		SaveMoneyReport(ctx Ctx, report dbmodels.MoneyReport) (err error)
		DeleteUnusedSessions(ctx Ctx, SessionRetentionDays int64) (int64, error)
		GetUnporcessedReportsByUserAndOrganization(ctx Ctx, userID string, organizationID uuid.UUID) ([]dbmodels.UserMoneyReport, error)
		GetUnprocessedMoneyReports(ctx Ctx, lastId int64, olderThenNMinutes int64) (reports []dbmodels.UserMoneyReport, err error)
		UpdateMoneyReport(ctx Ctx, id int64, processed bool) (err error)
	}
)
