package washServer

import (
	"washBonus/internal/app"

	"go.uber.org/zap"
)

type washService struct {
	logger         *zap.SugaredLogger
	washServerRepo app.WashServerRepo
}

func New(l *zap.SugaredLogger, repo app.WashServerRepo) app.WashServerService {
	return &washService{
		logger:         l,
		washServerRepo: repo,
	}
}
