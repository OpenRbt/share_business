package conversions

import (
	"github.com/OpenRbt/share_business/wash_rabbit/entity/admin"
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity"
	"wash_admin/internal/entity/vo"
	"wash_admin/openapi/models"
)

func WashServerFromDB(dbWashServer dbmodels.WashServer) entity.WashServer {
	return entity.WashServer{
		ID:          dbWashServer.ID.UUID,
		Title:       dbWashServer.Title,
		Description: dbWashServer.Description,
		Owner:       dbWashServer.Owner.UUID,
		ServiceKey:  dbWashServer.ServiceKey,
	}
}

func WashServerToRest(washServer entity.WashServer) *models.WashServer {
	return &models.WashServer{
		ID:          washServer.ID.String(),
		Name:        washServer.Title,
		Description: washServer.Description,
		ServiceKey:  washServer.ServiceKey,
	}
}

func RegisterWashServerFromRest(rest models.WashServerAdd) vo.RegisterWashServer {
	return vo.RegisterWashServer{
		Title:       *rest.Name,
		Description: rest.Description,
	}
}

func WashServerToRabbit(e entity.WashServer) admin.ServerRegistered {
	return admin.ServerRegistered{
		ID:          e.ID.String(),
		Title:       e.Title,
		Description: e.Description,
	}
}

func WashServerUpdateToRabbit(e vo.UpdateWashServer, deleted bool) admin.ServerUpdate {
	var del *bool
	if deleted {
		t := true
		del = &t
	}

	return admin.ServerUpdate{
		ID:          e.ID.String(),
		Title:       e.Title,
		Description: e.Description,
		Deleted:     del,
	}
}
