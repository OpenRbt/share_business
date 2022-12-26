package grpc

import (
	"go.uber.org/zap"
	"wash_bonus/intapi"
	"wash_bonus/internal/app/session"
	"wash_bonus/internal/app/user"
	"wash_bonus/internal/app/wash_server"
	"wash_bonus/internal/entity"
)

type SessionsCache interface {
}

type ConnectionsCache interface {
	Get(apiKey string) *entity.WashServerConnection
	Set(apiKey string, connection entity.WashServerConnection)
	Refresh(apiKey string, connection entity.WashServerConnection) error
}

type Service struct {
	l             *zap.SugaredLogger
	userSvc       user.Service
	washServerSvc wash_server.Service

	sessionsSvc session.Service

	connectionsCache ConnectionsCache

	intapi.UnsafeWashBonusServer

	basePath string
}

func New(l *zap.SugaredLogger, userSvc user.Service, washServerSvc wash_server.Service, sessionSvc session.Service, basePath string) *Service {
	svc := Service{
		l:             l,
		userSvc:       userSvc,
		washServerSvc: washServerSvc,
		sessionsSvc:   sessionSvc,
		basePath:      basePath,
	}

	return &svc
}
