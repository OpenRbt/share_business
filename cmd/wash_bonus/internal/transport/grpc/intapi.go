package grpc

import (
	"go.uber.org/zap"
	"sync"
	"wash_bonus/intapi"
	"wash_bonus/internal/app/balance"
	"wash_bonus/internal/app/wash_server"
	"wash_bonus/internal/entity"
)

type Service struct {
	l *zap.SugaredLogger

	balanceSvc    balance.Service
	washServerSvc wash_server.Service

	connectionsMutex sync.RWMutex
	connections      map[string]*entity.WashServerConnection
	intapi.UnsafeWashBonusServer
}

func New(l *zap.SugaredLogger, balanceSvc balance.Service, washServerSvc wash_server.Service) *Service {
	svc := Service{
		l:             l,
		balanceSvc:    balanceSvc,
		washServerSvc: washServerSvc,
		connections:   make(map[string]*entity.WashServerConnection),
	}

	return &svc
}
