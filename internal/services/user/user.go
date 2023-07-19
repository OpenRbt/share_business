package user

import (
	"washBonus/internal/app"

	"go.uber.org/zap"
)

type userService struct {
	logger   *zap.SugaredLogger
	userRepo app.UserRepo
}

func New(l *zap.SugaredLogger, userRepo app.UserRepo) app.UserService {
	return &userService{
		logger:   l,
		userRepo: userRepo,
	}
}
