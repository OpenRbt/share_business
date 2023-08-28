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

func (ctrl *serverGroupController) Get(ctx context.Context, authUser entity.User, filter entity.ServerGroupFilter) ([]entity.ServerGroup, error) {
	if app.IsAdmin(authUser) {
		return ctrl.serverGroupSvc.Get(ctx, filter)
	} else if app.IsEngineer(authUser) {
		return ctrl.serverGroupSvc.GetForManager(ctx, authUser.ID, filter)
	}

	return nil, entity.ErrAccessDenied
}

func (ctrl *serverGroupController) GetById(ctx context.Context, authUser entity.User, id uuid.UUID) (entity.ServerGroup, error) {
	if app.IsUser(authUser) {
		return entity.ServerGroup{}, entity.ErrAccessDenied
	}

	if app.IsAdmin(authUser) {
		return ctrl.serverGroupSvc.GetById(ctx, id)
	}

	isUserManager, err := ctrl.orgSvc.IsUserManager(ctx, id, authUser.ID)
	if err != nil {
		return entity.ServerGroup{}, err
	}

	if isUserManager {
		return ctrl.serverGroupSvc.GetById(ctx, id)
	}

	return entity.ServerGroup{}, entity.ErrAccessDenied
}

func (ctrl *serverGroupController) Create(ctx context.Context, authUser entity.User, ent entity.ServerGroupCreation) (entity.ServerGroup, error) {
	if !app.IsAdmin(authUser) {
		return entity.ServerGroup{}, entity.ErrAccessDenied
	}

	return ctrl.serverGroupSvc.Create(ctx, ent)
}

func (ctrl *serverGroupController) Update(ctx context.Context, authUser entity.User, id uuid.UUID, ent entity.ServerGroupUpdate) (entity.ServerGroup, error) {
	if !app.IsAdmin(authUser) {
		return entity.ServerGroup{}, entity.ErrAccessDenied
	}

	return ctrl.serverGroupSvc.Update(ctx, id, ent)
}

func (ctrl *serverGroupController) Delete(ctx context.Context, authUser entity.User, id uuid.UUID) error {
	if !app.IsAdmin(authUser) {
		return entity.ErrAccessDenied
	}

	return ctrl.serverGroupSvc.Delete(ctx, id)
}
