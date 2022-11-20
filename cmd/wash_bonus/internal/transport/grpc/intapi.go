package grpc

import (
	"go.uber.org/zap"
	"sync"
	"wash_bonus/intapi"
	"wash_bonus/internal/app/balance"
	"wash_bonus/internal/app/wash_server"
	"wash_bonus/internal/entity"
)

type Service interface {
}

type service struct {
	l *zap.SugaredLogger

	balanceSvc    balance.Service
	washServerSvc wash_server.Service

	connectionMutex sync.RWMutex
	connections     map[string]*entity.WashServerConnection
	intapi.UnimplementedServerServiceServer
	intapi.UnimplementedSessionServiceServer
}

func New(l *zap.SugaredLogger, balanceSvc balance.Service, washServerSvc wash_server.Service) Service {
	svc := service{
		l:             l,
		balanceSvc:    balanceSvc,
		washServerSvc: washServerSvc,
		connections:   make(map[string]*entity.WashServerConnection),
	}
	return &svc
}
