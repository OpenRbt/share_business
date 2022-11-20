package conversions

import (
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity"
	"wash_admin/openapi/models"
)

func WashServerFromDB(dbWashServer dbmodels.WashServer) entity.WashServer {
	return entity.WashServer{
		ID:          dbWashServer.ID.UUID,
		Name:        dbWashServer.Name,
		Description: dbWashServer.Description,
		APIKey:      dbWashServer.APIKey,
	}
}

func WashServerToRest(wash_server entity.WashServer) models.WashServer {
	return models.WashServer{
		ID:          wash_server.ID.String(),
		Name:        wash_server.Name,
		Description: wash_server.Description,
		APIKey:      wash_server.APIKey,
	}
}
