package conversions

import (
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
)

func WashServerFromDB(dbWashServer dbmodels.WashServer) entity.WashServer {
	return entity.WashServer{
		Id:          dbWashServer.ID,
		Name:        dbWashServer.Name,
		Description: dbWashServer.Description,
		WashKey:     dbWashServer.WashKey,
	}
}
