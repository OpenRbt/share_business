package controllers

import (
	"context"
	"washBonus/internal/app"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type organizationController struct {
	logger *zap.SugaredLogger
	orgSvc app.OrganizationService
}

func NewOrganizationController(l *zap.SugaredLogger, orgSvc app.OrganizationService) app.OrganizationController {
	return &organizationController{
		logger: l,
		orgSvc: orgSvc,
	}
}

func (ctrl *organizationController) Get(ctx context.Context, authUser entity.User, filter entity.OrganizationFilter) ([]entity.Organization, error) {
	if app.IsAdmin(authUser) {
		return ctrl.orgSvc.Get(ctx, filter)
	} else if app.IsEngineer(authUser) {
		return ctrl.orgSvc.GetForManager(ctx, authUser.ID, filter)
	}

	return nil, entity.ErrAccessDenied
}

func (ctrl *organizationController) GetById(ctx context.Context, authUser entity.User, id uuid.UUID) (entity.Organization, error) {
	if app.IsUser(authUser) {
		return entity.Organization{}, entity.ErrAccessDenied
	}

	if app.IsAdmin(authUser) {
		return ctrl.orgSvc.GetById(ctx, id)
	}

	isUserManager, err := ctrl.orgSvc.IsUserManager(ctx, id, authUser.ID)
	if err != nil {
		return entity.Organization{}, err
	}
	
	if isUserManager {
		return ctrl.orgSvc.GetById(ctx, id)
	}

	return entity.Organization{}, entity.ErrAccessDenied
}

func (ctrl *organizationController) Create(ctx context.Context, authUser entity.User, ent entity.OrganizationCreation) (entity.Organization, error) {
	if !app.IsAdmin(authUser) {
		return entity.Organization{}, entity.ErrAccessDenied
	}

	return ctrl.orgSvc.Create(ctx, ent)
}

func (ctrl *organizationController) Update(ctx context.Context, authUser entity.User, id uuid.UUID, ent entity.OrganizationUpdate) (entity.Organization, error) {
	if !app.IsAdmin(authUser) {
		return entity.Organization{}, entity.ErrAccessDenied
	}

	return ctrl.orgSvc.Update(ctx, id, ent)
}

func (ctrl *organizationController) Delete(ctx context.Context, authUser entity.User, id uuid.UUID) error {
	if !app.IsAdmin(authUser) {
		return entity.ErrAccessDenied
	}

	return ctrl.orgSvc.Delete(ctx, id)
}

func (ctrl *organizationController) AssignManager(ctx context.Context, authUser entity.User, organizationID uuid.UUID, userID string) error {
	if !app.IsAdmin(authUser) {
		return entity.ErrAccessDenied
	}

	return ctrl.orgSvc.AssignManager(ctx, organizationID, userID)
}

func (ctrl *organizationController) RemoveManager(ctx context.Context, authUser entity.User, organizationID uuid.UUID, userID string) error {
	if !app.IsAdmin(authUser) {
		return entity.ErrAccessDenied
	}

	return ctrl.orgSvc.RemoveManager(ctx, organizationID, userID)
}
