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

type serverGroupController struct {
	logger         *zap.SugaredLogger
	serverGroupSvc app.ServerGroupService
	orgSvc         app.OrganizationService
	rabbitSvc      rabbit.RabbitService
}

func NewServerGroupController(l *zap.SugaredLogger, groupSvc app.ServerGroupService, orgSvc app.OrganizationService, rabbitSvc rabbit.RabbitService) app.ServerGroupController {
	return &serverGroupController{
		logger:         l,
		serverGroupSvc: groupSvc,
		orgSvc:         orgSvc,
		rabbitSvc:      rabbitSvc,
	}
}

func (ctrl *serverGroupController) Get(ctx context.Context, auth app.AdminAuth, filter entities.ServerGroupFilter) ([]entities.ServerGroup, error) {
	if app.IsSystemManager(auth.User) {
		return ctrl.serverGroupSvc.Get(ctx, filter)
	}

	if app.IsAdmin(auth.User) && auth.User.Organization != nil {
		if filter.OrganizationID != nil && *filter.OrganizationID != auth.User.Organization.ID {
			return nil, entities.ErrForbidden
		}

		filter.OrganizationID = &auth.User.Organization.ID
		return ctrl.serverGroupSvc.Get(ctx, filter)
	}

	return nil, entities.ErrForbidden
}

func (ctrl *serverGroupController) GetById(ctx context.Context, auth app.AdminAuth, id uuid.UUID) (entities.ServerGroup, error) {
	group, err := ctrl.serverGroupSvc.GetById(ctx, id)
	if err != nil {
		return group, err
	}

	if app.IsSystemManager(auth.User) || app.IsAdminManageOrganization(auth.User, group.OrganizationID) {
		return group, nil
	}

	return entities.ServerGroup{}, nil
}

func (ctrl *serverGroupController) Create(ctx context.Context, auth app.AdminAuth, ent entities.ServerGroupCreation) (entities.ServerGroup, error) {
	if !app.IsSystemManager(auth.User) && !app.IsAdminManageOrganization(auth.User, ent.OrganizationID) {
		return entities.ServerGroup{}, entities.ErrForbidden
	}

	createdServer, err := ctrl.serverGroupSvc.Create(ctx, ent)
	if err != nil {
		return entities.ServerGroup{}, err
	}

	return createdServer, ctrl.sendServerGroupToServices(ctx, createdServer.ID)
}

func (ctrl *serverGroupController) Update(ctx context.Context, auth app.AdminAuth, id uuid.UUID, ent entities.ServerGroupUpdate) (entities.ServerGroup, error) {
	server, err := ctrl.serverGroupSvc.GetById(ctx, id)
	if err != nil {
		return entities.ServerGroup{}, err
	}

	if !app.IsSystemManager(auth.User) && !app.IsAdminManageOrganization(auth.User, server.OrganizationID) {
		return entities.ServerGroup{}, entities.ErrForbidden
	}

	updatedServer, err := ctrl.serverGroupSvc.Update(ctx, id, ent)
	if err != nil {
		return entities.ServerGroup{}, err
	}

	return updatedServer, ctrl.sendServerGroupToServices(ctx, id)
}

func (ctrl *serverGroupController) Delete(ctx context.Context, auth app.AdminAuth, id uuid.UUID) error {
	server, err := ctrl.serverGroupSvc.GetById(ctx, id)
	if err != nil {
		return err
	}

	if !app.IsSystemManager(auth.User) && !app.IsAdminManageOrganization(auth.User, server.OrganizationID) {
		return entities.ErrForbidden
	}

	err = ctrl.serverGroupSvc.Delete(ctx, id)
	if err != nil {
		return err
	}

	return ctrl.sendServerGroupToServices(ctx, id)
}

func (ctrl *serverGroupController) sendServerGroupToServices(ctx context.Context, id uuid.UUID) error {
	server, err := ctrl.serverGroupSvc.GetAnyById(ctx, id)
	if err != nil {
		return err
	}

	rabbitGroup := conversions.ServerGroupToRabbit(server)
	return ctrl.rabbitSvc.SendMessage(rabbitGroup, rabbitEntities.AdminsExchange, rabbitEntities.WashBonusRoutingKey, rabbitEntities.ServerGroupMessageType)
}
