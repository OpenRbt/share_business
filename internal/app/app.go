package app

import (
	"context"
	"washBonus/internal/config"
	"washBonus/internal/entity"
)

type (
	Ctx = context.Context
)

type (
	Auth struct {
		UID          string
		Disabled     bool
		User         entity.User
		UserMetadata *AuthUserMeta
	}

	AuthUserMeta struct {
		CreationTimestamp  int64
		LastLogInTimestamp int64

		LastRefreshTimestamp int64
	}
)

type (
	ScheduleService interface {
		Run(reportsDelayMinutes int, sessionsDelayMinutes int, SessionRetentionDays int64)
	}
)

type App struct {
	cfg             config.Config
	washService     WashServerService
	userService     UserService
	sessionService  SessionService
	scheduleService ScheduleService
}
