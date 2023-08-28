package organizations

import (
	"washBonus/internal/app"

	"go.uber.org/zap"
)

type organizationService struct {
	logger           *zap.SugaredLogger
	organizationRepo app.OrganizationRepo
	userRepo         app.UserRepo
}

func New(l *zap.SugaredLogger, orgRepo app.OrganizationRepo, userRepo app.UserRepo) app.OrganizationService {
	return &organizationService{
		logger:           l,
		organizationRepo: orgRepo,
		userRepo:         userRepo,
	}
}
