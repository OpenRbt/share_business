package wash_server

import (
	"context"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/entity/vo"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type Service interface {
	GetWashServer(ctx context.Context, id uuid.UUID) (entity.WashServer, error)
	GetWashServers(ctx context.Context) ([]entity.WashServer, error)
	CreateWashServer(ctx context.Context, server entity.WashServer) (entity.WashServer, error)
	UpdateWashServer(ctx context.Context, update vo.WashServerUpdate) error
}

type Repository interface {
	GetWashServer(ctx context.Context, id uuid.UUID) (entity.WashServer, error)
	GetWashServers(ctx context.Context) ([]entity.WashServer, error)
	CreateWashServer(ctx context.Context, server entity.WashServer) (entity.WashServer, error)
	UpdateWashServer(ctx context.Context, update vo.WashServerUpdate) error
}

type service struct {
	l    *zap.SugaredLogger
	repo Repository
}

func New(l *zap.SugaredLogger, repo Repository) Service {
	return &service{
		l:    l,
		repo: repo,
	}
}
