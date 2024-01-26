package controllers

import (
	"context"
	"washbonus/internal/app"
	"washbonus/internal/conversions"
	"washbonus/internal/entities"
	"washbonus/internal/infrastructure/rabbit"
	rabbitEntities "washbonus/internal/infrastructure/rabbit/entities"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type adminController struct {
	logger    *zap.SugaredLogger
	adminSvc  app.AdminService
	rabbitSvc rabbit.RabbitService
}

func NewAdminUserController(l *zap.SugaredLogger, adminSvc app.AdminService, rabbitSvc rabbit.RabbitService) app.AdminController {
	return &adminController{
		logger:    l,
		adminSvc:  adminSvc,
		rabbitSvc: rabbitSvc,
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
	if !app.IsSystemManager(auth.User) {
		return entities.ErrForbidden
	}

	err := ctrl.adminSvc.UpdateRole(ctx, userUpdate)
	if err != nil {
		return err
	}

	return ctrl.sendAdminUserToServices(ctx, userUpdate.ID)
}

func (ctrl *adminController) Block(ctx context.Context, auth app.AdminAuth, id string) error {
	if !app.IsSystemManager(auth.User) {
		return entities.ErrForbidden
	}

	err := ctrl.adminSvc.Block(ctx, id)
	if err != nil {
		return err
	}

	return ctrl.sendAdminUserToServices(ctx, id)
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
	if !app.IsSystemManager(auth.User) {
		return entities.ErrForbidden
	}

	appl, err := ctrl.adminSvc.GetApplicationByID(ctx, id)
	if err != nil {
		return err
	}

	err = ctrl.adminSvc.ReviewApplication(ctx, id, ent)
	if err != nil {
		return err
	}

	if ent.Status == entities.Accepted {
		return ctrl.sendAdminUserToServices(ctx, appl.User.ID)
	}

	return nil
}

func (ctrl *adminController) GetApplicationByID(ctx context.Context, auth app.AdminAuth, id uuid.UUID) (entities.AdminApplication, error) {
	if app.IsSystemManager(auth.User) {
		return ctrl.adminSvc.GetApplicationByID(ctx, id)
	}

	return entities.AdminApplication{}, entities.ErrForbidden
}

func (ctrl *adminController) sendAdminUserToServices(ctx context.Context, id string) error {
	user, err := ctrl.adminSvc.GetById(ctx, id)
	if err != nil {
		return err
	}

	rabbitUser := conversions.AdminUserToRabbit(user)
	return ctrl.rabbitSvc.SendMessage(rabbitUser, rabbitEntities.AdminsExchange, rabbitEntities.WashBonusRoutingKey, rabbitEntities.AdminUserMessageType)
}
