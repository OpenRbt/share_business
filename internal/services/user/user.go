package user

import (
	"washBonus/internal/app"

	"go.uber.org/zap"
)

type userService struct {
	logger   *zap.SugaredLogger
	userRepo app.UserRepo
	orgRepo  app.OrganizationRepo
}

func New(l *zap.SugaredLogger, userRepo app.UserRepo, orgRepo app.OrganizationRepo) app.UserService {
	return &userService{
		logger:   l,
		userRepo: userRepo,
		orgRepo:  orgRepo,
	}
}
