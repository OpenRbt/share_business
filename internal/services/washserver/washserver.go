package washserver

import (
	"washbonus/internal/app"

	"go.uber.org/zap"
)

type washService struct {
	logger          *zap.SugaredLogger
	washServerRepo  app.WashServerRepo
	serverGroupRepo app.ServerGroupRepo
}

func New(l *zap.SugaredLogger, washServerRepo app.WashServerRepo, serverGroupRepo app.ServerGroupRepo) app.WashServerService {
	return &washService{
		logger:          l,
		washServerRepo:  washServerRepo,
		serverGroupRepo: serverGroupRepo,
	}
}
