package controllers

import (
	"context"
	"washbonus/internal/app"
	"washbonus/internal/entities"

	"go.uber.org/zap"
)

type userController struct {
	logger     *zap.SugaredLogger
	userSvc    app.UserService
	sessionSvc app.SessionService
}

func NewUserController(l *zap.SugaredLogger, userSvc app.UserService, sessionSvc app.SessionService) app.UserController {
	return &userController{
		logger:     l,
		userSvc:    userSvc,
		sessionSvc: sessionSvc,
	}
}

func (ctrl *userController) GetById(ctx context.Context, auth app.Auth, userID string) (entities.User, error) {
	if auth.User.ID == userID {
		return ctrl.userSvc.GetById(ctx, userID)
	}

	return entities.User{}, entities.ErrForbidden
}
