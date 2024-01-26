package app

import (
	"context"
	"washbonus/internal/config"
	"washbonus/internal/entities"
)

type (
	Ctx = context.Context
)

type (
	Auth struct {
		User entities.User
	}

	AdminAuth struct {
		User         entities.AdminUser
		UserMetadata *AuthUserMeta
		Disabled     bool
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

type Repositories struct {
	Admin AdminRepo
	Org   OrganizationRepo
	Group ServerGroupRepo
	Wash  WashServerRepo

	User    UserRepo
	Session SessionRepo
	Wallet  WalletRepo
}

type Services struct {
	Admin AdminService
	Org   OrganizationService
	Group ServerGroupService
	Wash  WashServerService

	User    UserService
	Session SessionService
	Wallet  WalletService
}

type Controllers struct {
	Admin AdminController
	Org   OrganizationController
	Group ServerGroupController
	Wash  WashServerController

	User    UserController
	Session SessionController
	Wallet  WalletController
}

type FirebaseService interface {
	BonusAuth(token string) (*Auth, error)
	AdminAuth(token string) (*AdminAuth, error)
}
