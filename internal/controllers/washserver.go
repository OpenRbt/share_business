package controllers

import (
	"context"
	"washbonus/internal/app"
	"washbonus/internal/entities"
	"washbonus/internal/infrastructure/rabbit"
	rabbitEntities "washbonus/internal/infrastructure/rabbit/entities"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type washServerController struct {
	logger        *zap.SugaredLogger
	washServerSvc app.WashServerService
	groupSvc      app.ServerGroupService
	orgSvc        app.OrganizationService
	rabbitSvc     rabbit.RabbitService
}

func NewWashServerController(l *zap.SugaredLogger, washServerSvc app.WashServerService, groupSvc app.ServerGroupService, orgSvc app.OrganizationService, rabbitSvc rabbit.RabbitService) app.WashServerController {
	return &washServerController{
		logger:        l,
		washServerSvc: washServerSvc,
		orgSvc:        orgSvc,
		groupSvc:      groupSvc,
		rabbitSvc:     rabbitSvc,
	}
}

func (ctrl *washServerController) CreateWashServer(ctx context.Context, auth app.AdminAuth, newServer entities.WashServerCreation) (entities.WashServer, error) {
	var isAdminManager bool

	if newServer.GroupID != nil {
		group, err := ctrl.groupSvc.GetById(ctx, *newServer.GroupID)
		if err != nil {
			return entities.WashServer{}, err
		}

		isAdminManager = app.IsAdminManageOrganization(auth.User, group.OrganizationID)
	}

	if !app.IsSystemManager(auth.User) && !isAdminManager {
		return entities.WashServer{}, entities.ErrForbidden
	}

	registered, err := ctrl.washServerSvc.CreateWashServer(ctx, auth.User.ID, newServer)
	if err != nil {
		return entities.WashServer{}, err
	}

	return registered, ctrl.rabbitSvc.CreateRabbitUser(registered.ID.String(), *registered.ServiceKey)
}

func (ctrl *washServerController) GetWashServerById(ctx context.Context, auth app.AdminAuth, id uuid.UUID) (entities.WashServer, error) {
	server, err := ctrl.washServerSvc.GetWashServerById(ctx, id)
	if err != nil {
		return entities.WashServer{}, err
	}

	if app.IsSystemManager(auth.User) || app.IsAdminManageOrganization(auth.User, server.OrganizationID) {
		return server, nil
	}

	return entities.WashServer{}, entities.ErrForbidden
}

func (ctrl *washServerController) UpdateWashServer(ctx context.Context, auth app.AdminAuth, id uuid.UUID, updateWashServer entities.WashServerUpdate) (entities.WashServer, error) {
	server, err := ctrl.washServerSvc.GetWashServerById(ctx, id)
	if err != nil {
		return entities.WashServer{}, err
	}

	if app.IsSystemManager(auth.User) || app.IsAdminManageOrganization(auth.User, server.OrganizationID) {
		return ctrl.washServerSvc.UpdateWashServer(ctx, server.ID, updateWashServer)
	}

	return entities.WashServer{}, nil
}

func (ctrl *washServerController) DeleteWashServer(ctx context.Context, auth app.AdminAuth, id uuid.UUID) error {
	server, err := ctrl.washServerSvc.GetWashServerById(ctx, id)
	if err != nil {
		return err
	}

	if !app.IsSystemManager(auth.User) && !app.IsAdminManageOrganization(auth.User, server.OrganizationID) {
		return entities.ErrForbidden
	}

	err = ctrl.washServerSvc.DeleteWashServer(ctx, id)
	if err != nil {
		return err
	}

	err = ctrl.rabbitSvc.SendMessage(nil, rabbitEntities.WashBonusService, rabbitEntities.RoutingKey(id.String()), rabbitEntities.WashServerDeletionMessageType)
	if err != nil {
		return err
	}

	return ctrl.rabbitSvc.DeleteRabbitUser(ctx, id.String())
}

func (ctrl *washServerController) GetWashServers(ctx context.Context, auth app.AdminAuth, filter entities.WashServerFilter) ([]entities.WashServer, error) {
	if app.IsSystemManager(auth.User) {
		return ctrl.washServerSvc.GetWashServers(ctx, filter)
	}

	if app.IsAdmin(auth.User) && auth.User.Organization != nil {
		filter.OrganizationID = &auth.User.Organization.ID
		return ctrl.washServerSvc.GetWashServers(ctx, filter)
	}

	return nil, entities.ErrForbidden
}

func (ctrl *washServerController) AssignToServerGroup(ctx context.Context, auth app.AdminAuth, serverID uuid.UUID, groupID uuid.UUID) error {
	assigningServer, err := ctrl.washServerSvc.GetWashServerById(ctx, serverID)
	if err != nil {
		return err
	}

	targetGroup, err := ctrl.groupSvc.GetById(ctx, groupID)
	if err != nil {
		return err
	}

	isUserServerManager := app.IsAdminManageOrganization(auth.User, assigningServer.OrganizationID)
	if err != nil {
		return err
	}

	isUserGroupManager := app.IsAdminManageOrganization(auth.User, targetGroup.OrganizationID)
	if err != nil {
		return err
	}

	if app.IsSystemManager(auth.User) || (isUserServerManager && isUserGroupManager) {
		return ctrl.washServerSvc.AssignToServerGroup(ctx, serverID, groupID)
	}

	return entities.ErrForbidden
}
