package controllers

import (
	"context"
	"washBonus/internal/app"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type serverGroupController struct {
	logger         *zap.SugaredLogger
	serverGroupSvc app.ServerGroupService
	orgSvc         app.OrganizationService
}

func NewServerGroupController(l *zap.SugaredLogger, groupSvc app.ServerGroupService, orgSvc app.OrganizationService) app.ServerGroupController {
	return &serverGroupController{
		logger:         l,
		serverGroupSvc: groupSvc,
		orgSvc:         orgSvc,
	}
}

func (ctrl *serverGroupController) Get(ctx context.Context, auth app.Auth, filter entity.ServerGroupFilter) ([]entity.ServerGroup, error) {
	if !app.IsEngineer(auth.User) {
		filter.IsManagedByMe = false
	}

	return ctrl.serverGroupSvc.Get(ctx, auth.User.ID, filter)
}

func (ctrl *serverGroupController) GetById(ctx context.Context, auth app.Auth, id uuid.UUID) (entity.ServerGroup, error) {
	return ctrl.serverGroupSvc.GetById(ctx, id)
}

func (ctrl *serverGroupController) Create(ctx context.Context, auth app.Auth, ent entity.ServerGroupCreation) (entity.ServerGroup, error) {
	if app.IsUser(auth.User) {
		return entity.ServerGroup{}, entity.ErrAccessDenied
	}

	isUserManager, err := ctrl.orgSvc.IsUserManager(ctx, ent.OrganizationID, auth.User.ID)
	if err != nil {
		return entity.ServerGroup{}, err
	}

	if app.IsEngineer(auth.User) && !isUserManager {
		return entity.ServerGroup{}, entity.ErrAccessDenied
	}

	return ctrl.serverGroupSvc.Create(ctx, ent)
}

func (ctrl *serverGroupController) Update(ctx context.Context, auth app.Auth, id uuid.UUID, ent entity.ServerGroupUpdate) (entity.ServerGroup, error) {
	if app.IsUser(auth.User) {
		return entity.ServerGroup{}, entity.ErrAccessDenied
	}

	server, err := ctrl.serverGroupSvc.GetById(ctx, id)
	if err != nil {
		return entity.ServerGroup{}, err
	}

	isUserManager, err := ctrl.orgSvc.IsUserManager(ctx, server.OrganizationID, auth.User.ID)
	if err != nil {
		return entity.ServerGroup{}, err
	}

	if app.IsEngineer(auth.User) && !isUserManager {
		return entity.ServerGroup{}, entity.ErrAccessDenied
	}

	return ctrl.serverGroupSvc.Update(ctx, id, ent)
}

func (ctrl *serverGroupController) Delete(ctx context.Context, auth app.Auth, id uuid.UUID) error {
	if app.IsUser(auth.User) {
		return entity.ErrAccessDenied
	}

	server, err := ctrl.serverGroupSvc.GetById(ctx, id)
	if err != nil {
		return err
	}

	isUserManager, err := ctrl.orgSvc.IsUserManager(ctx, server.OrganizationID, auth.User.ID)
	if err != nil {
		return err
	}

	if app.IsEngineer(auth.User) && !isUserManager {
		return entity.ErrAccessDenied
	}

	return ctrl.serverGroupSvc.Delete(ctx, id)
}
