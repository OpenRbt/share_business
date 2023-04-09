package user

import (
	"context"
	"go.uber.org/zap"
	user_svc "wash_bonus/internal/app/user"
	"wash_bonus/internal/entity"
)

type UseCase interface {
	Get(ctx context.Context, userID string) (entity.User, error)
}

type useCase struct {
	l       *zap.SugaredLogger
	UserSvc user_svc.Service
}

func New(l *zap.SugaredLogger, user user_svc.Service) UseCase {
	return &useCase{
		l:       l,
		UserSvc: user,
	}
}
