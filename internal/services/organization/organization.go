package organization

import (
	"washbonus/internal/app"

	"go.uber.org/zap"
)

type organizationService struct {
	logger           *zap.SugaredLogger
	organizationRepo app.OrganizationRepo
	adminRepo        app.AdminRepo
}

func New(l *zap.SugaredLogger, orgRepo app.OrganizationRepo, adminRepo app.AdminRepo) app.OrganizationService {
	return &organizationService{
		logger:           l,
		organizationRepo: orgRepo,
		adminRepo:        adminRepo,
	}
}
