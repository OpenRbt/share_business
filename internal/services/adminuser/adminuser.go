package adminuser

import (
	"washbonus/internal/app"

	"go.uber.org/zap"
)

type adminService struct {
	logger    *zap.SugaredLogger
	adminRepo app.AdminRepo
	orgRepo   app.OrganizationRepo
}

func New(l *zap.SugaredLogger, adminRepo app.AdminRepo, orgRepo app.OrganizationRepo) app.AdminService {
	return &adminService{
		logger:    l,
		adminRepo: adminRepo,
		orgRepo:   orgRepo,
	}
}
