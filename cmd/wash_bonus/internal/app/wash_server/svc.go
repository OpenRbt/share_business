package wash_server

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"wash_bonus/internal/entity"
)

type Service interface {
	GetWashServer(ctx context.Context, id uuid.UUID) (entity.WashServer, error)
	GetWashServerByKey(ctx context.Context, key string) (entity.WashServer, error)
}
type Repository interface {
	GetWashServer(ctx context.Context, id uuid.UUID) (entity.WashServer, error)
	GetWashServerByKey(ctx context.Context, key string) (entity.WashServer, error)
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
