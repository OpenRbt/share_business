package controllers

import (
	"context"
	"washBonus/internal/app"
	"washBonus/internal/entity"

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

func (ctrl *userController) Get(ctx context.Context, auth app.Auth, pagination entity.Pagination) ([]entity.User, error) {
	if app.IsAdmin(auth.User) {
		return ctrl.userSvc.Get(ctx, pagination)
	}

	return nil, entity.ErrAccessDenied
}

func (ctrl *userController) GetById(ctx context.Context, auth app.Auth, userID string) (entity.User, error) {
	if auth.User.ID == userID || app.IsAdmin(auth.User) {
		return ctrl.userSvc.GetById(ctx, userID)
	}

	return entity.User{}, entity.ErrAccessDenied
}

func (ctrl *userController) UpdateUserRole(ctx context.Context, auth app.Auth, userUpdate entity.UserUpdate) error {
	if app.IsAdmin(auth.User) {
		return ctrl.userSvc.UpdateUserRole(ctx, userUpdate)
	}

	return entity.ErrAccessDenied
}
