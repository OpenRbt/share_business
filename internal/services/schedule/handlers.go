package schedule

import (
	"context"
	"time"
)

func (s *scheduleService) Run(delayMinutes int) {
	go s.ProcessMoneyReports(time.Duration(delayMinutes) * time.Minute)
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
