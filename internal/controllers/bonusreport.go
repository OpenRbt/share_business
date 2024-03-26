package controllers

import (
	"context"
	"washbonus/internal/app"
	"washbonus/internal/entities"

	"go.uber.org/zap"
)

type bonusReportController struct {
	logger          *zap.SugaredLogger
	bonusReportSvc  app.BonusReportService
	organizationSvc app.OrganizationService
}

func NewBonusReportController(l *zap.SugaredLogger, bonusReportSvc app.BonusReportService, organizationSvc app.OrganizationService) app.BonusReportController {
	return &bonusReportController{
		logger:          l,
		bonusReportSvc:  bonusReportSvc,
		organizationSvc: organizationSvc,
	}
}

func (ctrl *bonusReportController) List(ctx context.Context, auth app.AdminAuth, filter entities.BonusReportFilter) (entities.Page[entities.BonusReport], error) {
	if app.IsAdminHasNoAccess(auth.User) {
		return entities.Page[entities.BonusReport]{}, entities.ErrForbidden
	}

	if filter.OrganizationID != nil {
		org, err := ctrl.organizationSvc.GetById(ctx, *filter.OrganizationID)
		if err != nil {
			return entities.Page[entities.BonusReport]{}, err
		}
		if !app.IsSystemManager(auth.User) && !app.IsAdminManageOrganization(auth.User, org.ID) {
			return entities.Page[entities.BonusReport]{}, entities.ErrForbidden
		}
	}

	return ctrl.bonusReportSvc.List(ctx, filter)
}
