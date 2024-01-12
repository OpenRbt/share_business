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

type organizationController struct {
	logger    *zap.SugaredLogger
	orgSvc    app.OrganizationService
	groupSvc  app.ServerGroupService
	adminSvc  app.AdminService
	rabbitSvc rabbit.RabbitService
}

func NewOrganizationController(l *zap.SugaredLogger, orgSvc app.OrganizationService, groupSvc app.ServerGroupService, adminSvc app.AdminService, rabbitSvc rabbit.RabbitService) app.OrganizationController {
	return &organizationController{
		logger:    l,
		orgSvc:    orgSvc,
		groupSvc:  groupSvc,
		adminSvc:  adminSvc,
		rabbitSvc: rabbitSvc,
	}
}

func (ctrl *organizationController) Get(ctx context.Context, auth app.AdminAuth, filter entities.OrganizationFilter) ([]entities.Organization, error) {
	if app.IsSystemManager(auth.User) {
		return ctrl.orgSvc.Get(ctx, filter)
	}

	return nil, entities.ErrForbidden
}

func (ctrl *organizationController) GetById(ctx context.Context, auth app.AdminAuth, id uuid.UUID) (entities.Organization, error) {
	if app.IsSystemManager(auth.User) || app.IsAdminManageOrganization(auth.User, id) {
		return ctrl.orgSvc.GetById(ctx, id)
	}

	return entities.Organization{}, entities.ErrForbidden
}

func (ctrl *organizationController) Create(ctx context.Context, auth app.AdminAuth, ent entities.OrganizationCreation) (entities.Organization, error) {
	if !app.IsSystemManager(auth.User) {
		return entities.Organization{}, entities.ErrForbidden
	}

	org, err := ctrl.orgSvc.Create(ctx, ent)
	if err != nil {
		return entities.Organization{}, err
	}

	err = ctrl.sendOrganizationToServices(ctx, org.ID)
	if err != nil {
		return entities.Organization{}, err
	}

	return org, ctrl.sendServerGroupToServices(ctx, org.ID)
}

func (ctrl *organizationController) Update(ctx context.Context, auth app.AdminAuth, id uuid.UUID, ent entities.OrganizationUpdate) (entities.Organization, error) {
	if !app.IsSystemManager(auth.User) && !app.IsAdminManageOrganization(auth.User, id) {
		return entities.Organization{}, entities.ErrForbidden
	}

	org, err := ctrl.orgSvc.Update(ctx, id, ent)
	if err != nil {
		return entities.Organization{}, err
	}

	err = ctrl.sendOrganizationToServices(ctx, id)
	if err != nil {
		return entities.Organization{}, err
	}

	return org, nil
}

func (ctrl *organizationController) Delete(ctx context.Context, auth app.AdminAuth, id uuid.UUID) error {
	if !app.IsSystemManager(auth.User) {
		return entities.ErrForbidden
	}

	err := ctrl.orgSvc.Delete(ctx, id)
	if err != nil {
		return err
	}

	err = ctrl.sendOrganizationToServices(ctx, id)
	if err != nil {
		return err
	}

	err = ctrl.sendServerGroupToServices(ctx, id)
	if err != nil {
		return err
	}

	return ctrl.sendAdminUsersToServices(ctx, id)
}

func (ctrl *organizationController) sendOrganizationToServices(ctx context.Context, orgID uuid.UUID) error {
	org, err := ctrl.orgSvc.GetAnyByID(ctx, orgID)
	if err != nil {
		return err
	}

	rabbitOrg := conversions.OrganizationToRabbit(org)
	return ctrl.rabbitSvc.SendMessage(rabbitOrg, rabbitEntities.AdminsExchange, rabbitEntities.WashBonusRoutingKey, rabbitEntities.OrganizationMessageType)
}

func (ctrl *organizationController) sendServerGroupToServices(ctx context.Context, orgID uuid.UUID) error {
	group, err := ctrl.orgSvc.GetDefaultGroupByOrganizationId(ctx, orgID)
	if err != nil {
		return err
	}

	rabbitGroup := conversions.ServerGroupToRabbit(group)
	return ctrl.rabbitSvc.SendMessage(rabbitGroup, rabbitEntities.AdminsExchange, rabbitEntities.WashBonusRoutingKey, rabbitEntities.ServerGroupMessageType)
}

func (ctrl *organizationController) sendAdminUsersToServices(ctx context.Context, orgID uuid.UUID) error {
	users, err := ctrl.orgSvc.GetAdminUsersByOrganizationID(ctx, orgID)
	if err != nil {
		return err
	}

	rabbitUsers := conversions.AdminUsersToRabbit(users)

	for _, user := range rabbitUsers {
		err := ctrl.rabbitSvc.SendMessage(user, rabbitEntities.AdminsExchange, rabbitEntities.WashBonusRoutingKey, rabbitEntities.AdminUserMessageType)
		if err != nil {
			ctrl.logger.Warnf("failed to send admin user: %w", err)
		}
	}

	return nil
}

func (ctrl *organizationController) AssignManager(ctx context.Context, auth app.AdminAuth, organizationID uuid.UUID, userID string) error {
	if !app.IsSystemManager(auth.User) {
		return entities.ErrForbidden
	}

	err := ctrl.orgSvc.AssignManager(ctx, organizationID, userID)
	if err != nil {
		return err
	}

	user, err := ctrl.adminSvc.GetById(ctx, userID)
	if err != nil {
		return err
	}

	rabbitUser := conversions.AdminUserToRabbit(user)
	return ctrl.rabbitSvc.SendMessage(rabbitUser, rabbitEntities.AdminsExchange, rabbitEntities.WashBonusRoutingKey, rabbitEntities.AdminUserMessageType)
}

func (ctrl *organizationController) RemoveManager(ctx context.Context, auth app.AdminAuth, organizationID uuid.UUID, userID string) error {
	if !app.IsSystemManager(auth.User) {
		return entities.ErrForbidden
	}

	err := ctrl.orgSvc.RemoveManager(ctx, organizationID, userID)
	if err != nil {
		return err
	}

	user, err := ctrl.adminSvc.GetById(ctx, userID)
	if err != nil {
		return err
	}

	rabbitUser := conversions.AdminUserToRabbit(user)
	return ctrl.rabbitSvc.SendMessage(rabbitUser, rabbitEntities.AdminsExchange, rabbitEntities.WashBonusRoutingKey, rabbitEntities.AdminUserMessageType)
}
