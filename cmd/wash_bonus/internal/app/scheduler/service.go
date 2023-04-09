package scheduler

import (
	"context"
	"go.uber.org/zap"
	"time"
	"wash_bonus/internal/app/session"
)

type Service interface {
	Run(delayMinutes int)
}

type service struct {
	l          *zap.SugaredLogger
	sessionSvc session.Service
}

func New(l *zap.SugaredLogger, sessionSvc session.Service) Service {
	return &service{
		l:          l,
		sessionSvc: sessionSvc,
	}
}

func (s *service) Run(delayMinutes int) {
	go s.ProcessMoneyReports(time.Duration(delayMinutes) * time.Minute)
}

func (s *service) ProcessMoneyReports(delay time.Duration) {
	l := s.l.Named("ProcessMoneyReports")

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for {
		<-ticker.C

		ctx := context.TODO()

		err := s.sessionSvc.ProcessMoneyReports(ctx)
		if err != nil {
			l.Error(err)
		}
	}
}
