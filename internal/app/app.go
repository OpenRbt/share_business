package app

import (
	"washBonus/internal/config"
)

type (
	Auth struct {
		UID          string
		Disabled     bool
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
		Run(delayMinutes int)
	}
)

type App struct {
	cfg             config.Config
	washService     WashServerService
	userService     UserService
	sessionService  SessionService
	scheduleService ScheduleService
}
