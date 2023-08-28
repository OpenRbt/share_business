package serverGroups

import (
	"washBonus/internal/app"

	"go.uber.org/zap"
)

type serverGroupService struct {
	logger           *zap.SugaredLogger
	serverGroupRepo  app.ServerGroupRepo
	organizationRepo app.OrganizationRepo
}

func New(l *zap.SugaredLogger, groupRepo app.ServerGroupRepo, orgRepo app.OrganizationRepo) app.ServerGroupService {
	return &serverGroupService{
		logger:           l,
		serverGroupRepo:  groupRepo,
		organizationRepo: orgRepo,
	}
}
