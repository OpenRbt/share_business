package app

import (
	dalEntity "washbonus/internal/entities"
	rabbitEntities "washbonus/internal/infrastructure/rabbit/entities"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type (
	RabbitService interface {
		CreatePool(ctx Ctx, serverID uuid.UUID, postId int64, amount int64) (sessions rabbitEntities.PostSessions, err error)
		UpdateState(ctx Ctx, sessionID uuid.UUID, state rabbitEntities.SessionState) error

		ConfirmBonuses(ctx Ctx, sessionID uuid.UUID, amount decimal.Decimal) error
		DiscardBonuses(ctx Ctx, sessionID uuid.UUID, amount decimal.Decimal) error
		RewardBonuses(ctx Ctx, payload []byte, sessionID uuid.UUID, amount decimal.Decimal, messageUUID uuid.UUID) error

		SaveMoneyReport(ctx Ctx, report dalEntity.MoneyReport) error
	}
)
