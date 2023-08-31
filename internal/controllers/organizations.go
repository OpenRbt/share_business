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

func (ctrl *organizationController) Get(ctx context.Context, auth app.Auth, filter entity.OrganizationFilter) ([]entity.Organization, error) {
	if !app.IsEngineer(auth.User) {
		filter.IsManagedByMe = false
	}

	return ctrl.orgSvc.Get(ctx, auth.User.ID, filter)
}

func (ctrl *organizationController) GetById(ctx context.Context, auth app.Auth, id uuid.UUID) (entity.Organization, error) {
	return ctrl.orgSvc.GetById(ctx, id)
}

func (ctrl *organizationController) Create(ctx context.Context, auth app.Auth, ent entity.OrganizationCreation) (entity.Organization, error) {
	if !app.IsAdmin(auth.User) {
		return entity.Organization{}, entity.ErrAccessDenied
	}

	return ctrl.orgSvc.Create(ctx, ent)
}

func (ctrl *organizationController) Update(ctx context.Context, auth app.Auth, id uuid.UUID, ent entity.OrganizationUpdate) (entity.Organization, error) {
	if !app.IsAdmin(auth.User) {
		return entity.Organization{}, entity.ErrAccessDenied
	}

	return ctrl.orgSvc.Update(ctx, id, ent)
}

func (ctrl *organizationController) Delete(ctx context.Context, auth app.Auth, id uuid.UUID) error {
	if !app.IsAdmin(auth.User) {
		return entity.ErrAccessDenied
	}

	return ctrl.orgSvc.Delete(ctx, id)
}

func (ctrl *organizationController) AssignManager(ctx context.Context, auth app.Auth, organizationID uuid.UUID, userID string) error {
	if !app.IsAdmin(auth.User) {
		return entity.ErrAccessDenied
	}

	return ctrl.orgSvc.AssignManager(ctx, organizationID, userID)
}

func (ctrl *organizationController) RemoveManager(ctx context.Context, auth app.Auth, organizationID uuid.UUID, userID string) error {
	if !app.IsAdmin(auth.User) {
		return entity.ErrAccessDenied
	}

	return ctrl.orgSvc.RemoveManager(ctx, organizationID, userID)
}
