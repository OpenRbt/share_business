package app

import (
	"context"
	dalEntity "washBonus/internal/entity"
	"washBonus/internal/infrastructure/rabbit/entity/session"
	"washBonus/internal/infrastructure/rabbit/entity/vo"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type (
	RabbitService interface {
		CreatePool(ctx context.Context, serverID uuid.UUID, postId int64, amount int64) (sessions session.PostSessions, err error)
		UpdateState(ctx context.Context, sessionID uuid.UUID, state vo.SessionState) error

		ConfirmBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) error
		DiscardBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) error
		RewardBonuses(ctx context.Context, payload []byte, sessionID uuid.UUID, amount decimal.Decimal, messageUUID uuid.UUID) error

		SaveMoneyReport(ctx context.Context, report dalEntity.MoneyReport) error
	}
)
