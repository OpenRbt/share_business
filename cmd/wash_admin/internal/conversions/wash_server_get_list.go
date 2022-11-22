package conversions

import (
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity"
	"wash_admin/internal/entity/vo"
	"wash_admin/openapi/models"
)

func PaginationToDB(entity vo.Pagination) dbmodels.Pagination {
	return dbmodels.Pagination{
		Limit: entity.Limit,
		Offset: entity.Offset,
	}
}

func PaginationFromRest (model models.Pagination) vo.Pagination {
	return vo.Pagination{
		Limit: model.Limit,
		Offset: model.Offset,
	}
}

func WashServerListFromDB (washServerList []dbmodels.WashServer) []entity.WashServer {
	res := make([]entity.WashServer, len(washServerList))

	for i, value := range washServerList {
		res[i] = WashServerFromDB(value)
	}

	return res
}

func WashServerListToRest (washServerEntity []entity.WashServer) []*models.WashServer {
	res := make([]*models.WashServer, len(washServerEntity))

	for i, value := range washServerEntity {
		rest := WashServerToRest(value)
		res[i] = &rest
	}

	return res
}