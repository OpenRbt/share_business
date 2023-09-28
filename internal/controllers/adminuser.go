package controllers

import (
	"context"
	"washbonus/internal/app"
	"washbonus/internal/entities"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type adminController struct {
	logger   *zap.SugaredLogger
	adminSvc app.AdminService
}

func NewAdminUserController(l *zap.SugaredLogger, adminSvc app.AdminService) app.AdminController {
	return &adminController{
		logger:   l,
		adminSvc: adminSvc,
	}
}

func (ctrl *adminController) GetById(ctx context.Context, auth app.AdminAuth, userID string) (entities.AdminUser, error) {
	if app.IsSystemManager(auth.User) || auth.User.ID == userID {
		return ctrl.adminSvc.GetById(ctx, userID)
	}

	return entities.AdminUser{}, entities.ErrForbidden
}

func (ctrl *adminController) Get(ctx context.Context, auth app.AdminAuth, filter entities.AdminUserFilter) ([]entities.AdminUser, error) {
	if app.IsSystemManager(auth.User) {
		return ctrl.adminSvc.Get(ctx, filter)
	}

	return nil, entities.ErrForbidden
}

func (ctrl *adminController) UpdateRole(ctx context.Context, auth app.AdminAuth, userUpdate entities.AdminUserRoleUpdate) error {
	if app.IsSystemManager(auth.User) {
		return ctrl.adminSvc.UpdateRole(ctx, userUpdate)
	}

	return entities.ErrForbidden
}

func (ctrl *adminController) Block(ctx context.Context, auth app.AdminAuth, id string) error {
	if app.IsSystemManager(auth.User) {
		return ctrl.adminSvc.Block(ctx, id)
	}

	return entities.ErrForbidden
}

func (ctrl *adminController) GetApplications(ctx context.Context, auth app.AdminAuth, filter entities.AdminApplicationFilter) ([]entities.AdminApplication, error) {
	if app.IsSystemManager(auth.User) {
		return ctrl.adminSvc.GetApplications(ctx, filter)
	}

	return nil, entities.ErrForbidden
}

func (ctrl *adminController) CreateApplication(ctx context.Context, ent entities.AdminApplicationCreation) (entities.AdminApplication, error) {
	return ctrl.adminSvc.CreateApplication(ctx, ent)
}

func (ctrl *adminController) ReviewApplication(ctx context.Context, auth app.AdminAuth, id uuid.UUID, ent entities.AdminApplicationReview) error {
	if app.IsSystemManager(auth.User) {
		return ctrl.adminSvc.ReviewApplication(ctx, id, ent)
	}

	return entities.ErrForbidden
}

func (ctrl *adminController) GetApplicationByID(ctx context.Context, auth app.AdminAuth, id uuid.UUID) (entities.AdminApplication, error) {
	if app.IsSystemManager(auth.User) {
		return ctrl.adminSvc.GetApplicationByID(ctx, id)
	}

	return entities.AdminApplication{}, entities.ErrForbidden
}
