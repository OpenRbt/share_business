package conversions

import (
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity"
	"wash_admin/internal/entity/vo"
	models2 "wash_admin/internal/infrastructure/rabbit/models"
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

func WashServerToRabbit(e entity.WashServer) models2.ServerRegistered {
	return models2.ServerRegistered{
		ID:          e.ID.String(),
		Title:       e.Title,
		Description: e.Description,
	}
}

func WashServerUpdateToRabbit(e vo.UpdateWashServer, deleted bool) models2.ServerUpdate {
	var delete *bool
	if deleted {
		t := true
		delete = &t
	}

	return models2.ServerUpdate{
		ID:          e.ID.String(),
		Title:       e.Title,
		Description: e.Description,
		Deleted:     delete,
	}
}
