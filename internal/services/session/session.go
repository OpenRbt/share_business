package session

import (
	"washbonus/internal/app"

	"go.uber.org/zap"
)

type sessionService struct {
	logger         *zap.SugaredLogger
	sessionRepo    app.SessionRepo
	userRepo       app.UserRepo
	washServerRepo app.WashServerRepo
	walletRepo     app.WalletRepo
}

func New(l *zap.SugaredLogger, userRepo app.UserRepo, sessionRepo app.SessionRepo, washServerRepo app.WashServerRepo, walletRepo app.WalletRepo) app.SessionService {
	return &sessionService{
		logger:         l,
		sessionRepo:    sessionRepo,
		userRepo:       userRepo,
		washServerRepo: washServerRepo,
		walletRepo:     walletRepo,
	}
}
