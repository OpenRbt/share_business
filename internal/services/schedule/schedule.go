package schedule

import (
	"washbonus/internal/app"

	"go.uber.org/zap"
)

type scheduleService struct {
	l          *zap.SugaredLogger
	sessionSvc app.SessionService
}

func New(l *zap.SugaredLogger, sessionSvc app.SessionService) app.ScheduleService {
	return &scheduleService{
		l:          l,
		sessionSvc: sessionSvc,
	}
}
