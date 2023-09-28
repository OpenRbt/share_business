package rabbit

import (
	"washbonus/internal/app"

	"go.uber.org/zap"
)

type rabbitService struct {
	logger     *zap.SugaredLogger
	sessionSvc app.SessionService
	userSvc    app.UserService
	washSvc    app.WashServerService
	walletSvc  app.WalletService
}

func New(l *zap.SugaredLogger, session app.SessionService, user app.UserService, wash app.WashServerService, walletSvc app.WalletService) app.RabbitService {
	return &rabbitService{
		logger:     l,
		sessionSvc: session,
		userSvc:    user,
		washSvc:    wash,
		walletSvc:  walletSvc,
	}
}
