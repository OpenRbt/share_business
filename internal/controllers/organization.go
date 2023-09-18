package controllers

import (
	"context"
	"washbonus/internal/app"
	"washbonus/internal/entities"

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

func (ctrl *organizationController) Get(ctx context.Context, auth app.AdminAuth, filter entities.OrganizationFilter) ([]entities.Organization, error) {
	if app.IsSystemManager(auth.User) {
		return ctrl.orgSvc.Get(ctx, auth.User.ID, filter)
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
	if app.IsSystemManager(auth.User) {
		return ctrl.orgSvc.Create(ctx, ent)
	}

	return entities.Organization{}, entities.ErrForbidden
}

func (ctrl *organizationController) Update(ctx context.Context, auth app.AdminAuth, id uuid.UUID, ent entities.OrganizationUpdate) (entities.Organization, error) {
	if app.IsSystemManager(auth.User) || app.IsAdminManageOrganization(auth.User, id) {
		return ctrl.orgSvc.Update(ctx, id, ent)
	}

	return entities.Organization{}, entities.ErrForbidden
}

func (ctrl *organizationController) Delete(ctx context.Context, auth app.AdminAuth, id uuid.UUID) error {
	if app.IsSystemManager(auth.User) {
		return ctrl.orgSvc.Delete(ctx, id)
	}

	return entities.ErrForbidden
}

func (ctrl *organizationController) AssignManager(ctx context.Context, auth app.AdminAuth, organizationID uuid.UUID, userID string) error {
	if app.IsSystemManager(auth.User) {
		return ctrl.orgSvc.AssignManager(ctx, organizationID, userID)
	}

	return entities.ErrForbidden
}

func (ctrl *organizationController) RemoveManager(ctx context.Context, auth app.AdminAuth, organizationID uuid.UUID, userID string) error {
	if app.IsSystemManager(auth.User) {
		return ctrl.orgSvc.RemoveManager(ctx, organizationID, userID)
	}

	return entities.ErrForbidden
}
