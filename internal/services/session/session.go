package session

import (
	"washBonus/internal/app"

	"go.uber.org/zap"
)

type sessionService struct {
	logger         *zap.SugaredLogger
	sessionRepo    app.SessionRepo
	userRepo       app.UserRepo
	washServerRepo app.WashServerRepo

	reportsProcessingDelayInMinutes  int64
	moneyReportsRewardPercentDefault int64
}

func New(l *zap.SugaredLogger, userRepo app.UserRepo, sessionRepo app.SessionRepo, washServerRepo app.WashServerRepo, reportsProcessingDelayInMinutes int64, moneyReportsRewardPercent int64) app.SessionService {
	return &sessionService{
		logger:         l,
		sessionRepo:    sessionRepo,
		userRepo:       userRepo,
		washServerRepo: washServerRepo,

		reportsProcessingDelayInMinutes:  reportsProcessingDelayInMinutes,
		moneyReportsRewardPercentDefault: moneyReportsRewardPercent,
	}
}
