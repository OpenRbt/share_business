package controllers

import (
	"context"
	"washBonus/internal/app"
	"washBonus/internal/entity"
	"washBonus/internal/infrastructure/rabbit"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type washServerController struct {
	logger        *zap.SugaredLogger
	washServerSvc app.WashServerService
	userSvc       app.UserService
	groupSvc      app.ServerGroupService
	orgSvc        app.OrganizationService
	rabbitSvc     rabbit.RabbitService
}

func NewWashServerController(l *zap.SugaredLogger, washServerSvc app.WashServerService, userSvc app.UserService, groupSvc app.ServerGroupService, orgSvc app.OrganizationService, rabbitSvc rabbit.RabbitService) app.WashServerController {
	return &washServerController{
		logger:        l,
		washServerSvc: washServerSvc,
		userSvc:       userSvc,
		orgSvc:        orgSvc,
		groupSvc:      groupSvc,
		rabbitSvc:     rabbitSvc,
	}
}

func (ctrl *washServerController) CreateWashServer(ctx context.Context, authUser entity.User, newServer entity.WashServerCreation) (entity.WashServer, error) {
	if app.IsUser(authUser) {
		return entity.WashServer{}, entity.ErrAccessDenied
	}

	groupID := newServer.GroupID
	if groupID.Valid {
		group, err := ctrl.groupSvc.GetById(ctx, groupID.UUID)
		if err != nil {
			return entity.WashServer{}, err
		}

		isUserManager, err := ctrl.orgSvc.IsUserManager(ctx, group.OrganizationID, authUser.ID)
		if err != nil {
			return entity.WashServer{}, err
		}

		if app.IsEngineer(authUser) && !isUserManager {
			return entity.WashServer{}, entity.ErrAccessDenied
		}
	}

	registered, err := ctrl.washServerSvc.CreateWashServer(ctx, authUser.ID, newServer)
	if err != nil {
		return entity.WashServer{}, err
	}

	err = ctrl.rabbitSvc.CreateRabbitUser(registered.ID.String(), registered.ServiceKey)
	if err != nil {
		return entity.WashServer{}, err
	}

	return registered, nil
}

func (ctrl *washServerController) GetWashServerById(ctx context.Context, authUser entity.User, id uuid.UUID) (entity.WashServer, error) {
	if app.IsUser(authUser) {
		return entity.WashServer{}, entity.ErrAccessDenied
	}

	return ctrl.washServerSvc.GetWashServerById(ctx, id)
}

func (ctrl *washServerController) UpdateWashServer(ctx context.Context, authUser entity.User, id uuid.UUID, updateWashServer entity.WashServerUpdate) (entity.WashServer, error) {
	if app.IsUser(authUser) {
		return entity.WashServer{}, entity.ErrAccessDenied
	}

	server, err := ctrl.washServerSvc.GetWashServerById(ctx, id)
	if err != nil {
		return entity.WashServer{}, err
	}

	isUserManager, err := ctrl.orgSvc.IsUserManager(ctx, server.OrganizationID, authUser.ID)
	if err != nil {
		return entity.WashServer{}, err
	}

	if app.IsEngineer(authUser) && !isUserManager {
		return entity.WashServer{}, entity.ErrAccessDenied
	}

	updatedServer, err := ctrl.washServerSvc.UpdateWashServer(ctx, server.ID, updateWashServer)
	if err != nil {
		return entity.WashServer{}, err
	}

	return updatedServer, nil
}

func (ctrl *washServerController) DeleteWashServer(ctx context.Context, authUser entity.User, id uuid.UUID) error {
	if app.IsUser(authUser) {
		return entity.ErrAccessDenied
	}

	server, err := ctrl.washServerSvc.GetWashServerById(ctx, id)
	if err != nil {
		return err
	}

	isUserManager, err := ctrl.orgSvc.IsUserManager(ctx, server.OrganizationID, authUser.ID)
	if err != nil {
		return err
	}

	if app.IsEngineer(authUser) && !isUserManager {
		return entity.ErrAccessDenied
	}

	err = ctrl.washServerSvc.DeleteWashServer(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *washServerController) GetWashServers(ctx context.Context, authUser entity.User, filter entity.WashServerFilter) ([]entity.WashServer, error) {
	if app.IsAdmin(authUser) {
		return ctrl.washServerSvc.GetWashServers(ctx, filter)
	} else if app.IsEngineer(authUser) {
		return ctrl.washServerSvc.GetForManager(ctx, authUser.ID, filter)
	}

	return ctrl.washServerSvc.GetWashServers(ctx, filter)
}

func (ctrl *washServerController) AssignToServerGroup(ctx context.Context, authUser entity.User, serverID uuid.UUID, groupID uuid.UUID) error {
	if app.IsUser(authUser) {
		return entity.ErrAccessDenied
	}

	server, err := ctrl.washServerSvc.GetWashServerById(ctx, serverID)
	if err != nil {
		return err
	}

	group, err := ctrl.groupSvc.GetById(ctx, groupID)
	if err != nil {
		return err
	}

	isUserServerManager, err := ctrl.orgSvc.IsUserManager(ctx, server.OrganizationID, authUser.ID)
	if err != nil {
		return err
	}

	isUserGroupManager, err := ctrl.orgSvc.IsUserManager(ctx, group.OrganizationID, authUser.ID)
	if err != nil {
		return err
	}

	if app.IsEngineer(authUser) && (!isUserServerManager || !isUserGroupManager) {
		return entity.ErrAccessDenied
	}

	return ctrl.washServerSvc.AssignToServerGroup(ctx, serverID, groupID)
}
