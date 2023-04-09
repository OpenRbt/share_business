package rabbit

import (
	"context"
	rabbit_session "github.com/OpenRbt/share_business/wash_rabbit/entity/session"
	rabbit_vo "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	session_svc "wash_bonus/internal/app/session"
	user_svc "wash_bonus/internal/app/user"
	"wash_bonus/internal/app/wash_server"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/entity/vo"
)

type UseCase interface {
	CreatePool(ctx context.Context, serverID uuid.UUID, postId int64, amount int64) (sessions rabbit_session.PostSessions, err error)
	UpdateState(ctx context.Context, sessionID uuid.UUID, state rabbit_vo.SessionState) error

	ConfirmBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error)
	DiscardBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error)
	RewardBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error)

	SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error)

	CreateWashServer(ctx context.Context, server entity.WashServer) (entity.WashServer, error)
	UpdateWashServer(ctx context.Context, update vo.WashServerUpdate) error
}

type useCase struct {
	l          *zap.SugaredLogger
	SessionSvc session_svc.Service
	UserSvc    user_svc.Service
	WashSvc    wash_server.Service
}

func New(l *zap.SugaredLogger, session session_svc.Service, user user_svc.Service, wash wash_server.Service) UseCase {
	return &useCase{
		l:          l,
		SessionSvc: session,
		UserSvc:    user,
		WashSvc:    wash,
	}
}
