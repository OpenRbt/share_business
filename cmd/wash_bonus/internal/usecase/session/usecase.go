package session

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	session_svc "wash_bonus/internal/app/session"
	user_svc "wash_bonus/internal/app/user"
	"wash_bonus/internal/app/wash_server"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/infrastructure/rabbit"
)

type UseCase interface {
	Get(ctx context.Context, sessionID uuid.UUID, user string) (entity.Session, error)
	AssignUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error)
	ChargeBonuses(ctx context.Context, sessionID uuid.UUID, userID string, amount decimal.Decimal) (err error)
}

type useCase struct {
	l          *zap.SugaredLogger
	SessionSvc session_svc.Service
	UserSvc    user_svc.Service
	RabbitSvc  rabbit.Rabbit
	WashSvc    wash_server.Service
}

func New(l *zap.SugaredLogger, session session_svc.Service, user user_svc.Service, wash wash_server.Service, rabbit rabbit.Rabbit) UseCase {
	return &useCase{
		l:          l,
		SessionSvc: session,
		UserSvc:    user,
		RabbitSvc:  rabbit,
		WashSvc:    wash,
	}
}
