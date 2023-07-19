package controllers

import (
	"context"
	"washBonus/internal/app"
	"washBonus/internal/entity"

	"go.uber.org/zap"
)

type userController struct {
	logger  *zap.SugaredLogger
	userSvc app.UserService
}

func NewUserController(l *zap.SugaredLogger, userSvc app.UserService) app.UserController {
	return &userController{
		logger:  l,
		userSvc: userSvc,
	}
}

func (ctrl *userController) Get(ctx context.Context, authorizedUserID string, userID string) (entity.User, error) {
	if authorizedUserID == userID {
		return ctrl.userSvc.GetOrCreate(ctx, authorizedUserID)
	}

	user, err := ctrl.userSvc.GetOrCreate(ctx, authorizedUserID)
	if err != nil {
		return entity.User{}, err
	}

	switch user.Role {
	case entity.AdminRole:
		return ctrl.userSvc.Get(ctx, userID)
	default:
		return entity.User{}, entity.ErrAccessDenied
	}
}

func (ctrl *userController) UpdateUserRole(ctx context.Context, authorizedUserID string, userUpdate entity.UpdateUser) error {
	user, err := ctrl.userSvc.GetOrCreate(ctx, authorizedUserID)
	if err != nil {
		return err
	}

	switch user.Role {
	case entity.AdminRole:
		return ctrl.userSvc.UpdateUserRole(ctx, authorizedUserID, userUpdate)
	default:
		return entity.ErrAccessDenied
	}
}
