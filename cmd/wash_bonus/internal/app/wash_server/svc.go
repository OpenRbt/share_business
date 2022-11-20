package wash_server

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type Service interface {
	GetWashServer(ctx context.Context, id uuid.UUID)
}
type Repository interface {
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
