package conversions

import (
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity/vo"
	"wash_admin/openapi/models"
)

func AddWashServerToDB(entity vo.AddWashServer) dbmodels.AddWashServer {
	return dbmodels.AddWashServer{
		Name:        entity.Name,
		Description: entity.Description,
	}
}

func AddWashServerFromRest(model models.WashServerAdd) vo.AddWashServer {
	return vo.AddWashServer{
		Name:        *model.Name,
		Description: model.Description,
	}
}
