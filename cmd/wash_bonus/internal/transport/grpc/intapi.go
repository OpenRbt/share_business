package grpc

import (
	"go.uber.org/zap"
	"wash_bonus/intapi"
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

	connectionsCache ConnectionsCache
	sessionsCache    SessionsCache

	intapi.UnsafeWashBonusServer
}

func New(l *zap.SugaredLogger, userSvc user.Service, washServerSvc wash_server.Service) *Service {
	svc := Service{
		l:             l,
		userSvc:       userSvc,
		washServerSvc: washServerSvc,
	}

	return &svc
}
