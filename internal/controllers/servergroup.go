package controllers

import (
	"context"
	"washbonus/internal/app"
	"washbonus/internal/entities"

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

func (ctrl *serverGroupController) Get(ctx context.Context, auth app.AdminAuth, filter entities.ServerGroupFilter) ([]entities.ServerGroup, error) {
	if app.IsSystemManager(auth.User) {
		return ctrl.serverGroupSvc.Get(ctx, filter)
	}

	if app.IsAdmin(auth.User) && auth.User.OrganizationID != nil {
		if filter.OrganizationID != nil && *filter.OrganizationID != *auth.User.OrganizationID {
			return nil, entities.ErrForbidden
		}

		filter.OrganizationID = auth.User.OrganizationID
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
	if app.IsSystemManager(auth.User) || app.IsAdminManageOrganization(auth.User, ent.OrganizationID) {
		return ctrl.serverGroupSvc.Create(ctx, ent)
	}

	return entities.ServerGroup{}, entities.ErrForbidden
}

func (ctrl *serverGroupController) Update(ctx context.Context, auth app.AdminAuth, id uuid.UUID, ent entities.ServerGroupUpdate) (entities.ServerGroup, error) {
	server, err := ctrl.serverGroupSvc.GetById(ctx, id)
	if err != nil {
		return entities.ServerGroup{}, err
	}

	if app.IsSystemManager(auth.User) || app.IsAdminManageOrganization(auth.User, server.OrganizationID) {
		return ctrl.serverGroupSvc.Update(ctx, id, ent)
	}

	return entities.ServerGroup{}, entities.ErrForbidden
}

func (ctrl *serverGroupController) Delete(ctx context.Context, auth app.AdminAuth, id uuid.UUID) error {
	server, err := ctrl.serverGroupSvc.GetById(ctx, id)
	if err != nil {
		return err
	}

	if app.IsSystemManager(auth.User) || app.IsAdminManageOrganization(auth.User, server.OrganizationID) {
		return ctrl.serverGroupSvc.Delete(ctx, id)
	}

	return entities.ErrForbidden
}
