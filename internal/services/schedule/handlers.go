package schedule

import (
	"context"
	"time"
)

func (s *scheduleService) Run(reportsDelayMinutes int, sessionsDelayMinutes int, SessionRetentionDays int64) {
	reportsDelay := time.Duration(reportsDelayMinutes) * time.Minute
	sessionsDelay := time.Duration(sessionsDelayMinutes) * time.Minute

	go s.ProcessMoneyReports(reportsDelay)
	go s.ProcessUnusedSessions(sessionsDelay, SessionRetentionDays)
}

func (s *scheduleService) ProcessMoneyReports(delay time.Duration) {
	l := s.l.Named("ProcessMoneyReports")

	for {
		ctx := context.TODO()

		err := s.sessionSvc.ProcessMoneyReports(ctx)
		if err != nil {
			l.Error(err)
		}

		time.Sleep(delay)
	}
}

func (s *scheduleService) ProcessUnusedSessions(delay time.Duration, SessionRetentionDays int64) {
	l := s.l.Named("ProcessUnusedSessions")

	for {
		ctx := context.TODO()

		count, err := s.sessionSvc.DeleteUnusedSessions(ctx, SessionRetentionDays)
		if err != nil {
			l.Error(err)
		}

		for count > 0 {
			count, err = s.sessionSvc.DeleteUnusedSessions(ctx, SessionRetentionDays)
			if err != nil {
				l.Error(err)
			}
		}

		time.Sleep(delay)
	}
}
