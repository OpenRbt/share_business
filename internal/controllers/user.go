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

func (ctrl *userController) Get(ctx context.Context, authorizedUserID string, userID string) (entity.User, error) {
	var user entity.User
	var err error

	if authorizedUserID == userID {
		user, err = ctrl.userSvc.GetOrCreate(ctx, userID)
	} else {
		authorizer, err := ctrl.userSvc.GetOrCreate(ctx, authorizedUserID)
		if err != nil {
			return entity.User{}, err
		}

		if authorizer.Role != entity.AdminRole {
			return entity.User{}, entity.ErrAccessDenied
		}

		user, err = ctrl.userSvc.Get(ctx, userID)
	}

	if err != nil {
		return entity.User{}, err
	}

	pendingBalance, err := ctrl.sessionSvc.GetUserPendingBalance(ctx, userID)
	if err != nil {
		return entity.User{}, err
	}
	user.PendingBalance = pendingBalance

	return user, nil
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
