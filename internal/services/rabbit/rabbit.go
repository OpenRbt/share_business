package rabbit

import (
	"washBonus/internal/app"

	"go.uber.org/zap"
)

type rabbitService struct {
	logger     *zap.SugaredLogger
	sessionSvc app.SessionService
	userSvc    app.UserService
	washSvc    app.WashServerService
}

func New(l *zap.SugaredLogger, session app.SessionService, user app.UserService, wash app.WashServerService) app.RabbitService {
	return &rabbitService{
		logger:     l,
		sessionSvc: session,
		userSvc:    user,
		washSvc:    wash,
	}
}
