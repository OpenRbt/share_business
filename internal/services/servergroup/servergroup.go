package servergroup

import (
	"washbonus/internal/app"

	"go.uber.org/zap"
)

type serverGroupService struct {
	logger    *zap.SugaredLogger
	groupRepo app.ServerGroupRepo
	orgRepo   app.OrganizationRepo
}

func New(l *zap.SugaredLogger, groupRepo app.ServerGroupRepo, orgRepo app.OrganizationRepo) app.ServerGroupService {
	return &serverGroupService{
		logger:    l,
		groupRepo: groupRepo,
		orgRepo:   orgRepo,
	}
}
