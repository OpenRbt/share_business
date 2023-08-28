package conversions

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/openapi/models"
)

func PaginationToDB(entity entity.Pagination) dbmodels.Pagination {
	return dbmodels.Pagination{
		Limit:  entity.Limit,
		Offset: entity.Offset,
	}
}

func PaginationFromRest(model models.Pagination) entity.Pagination {
	limit := model.Limit
	if limit == 0 {
		limit = 100
	}

	return entity.Pagination{
		Limit:  limit,
		Offset: model.Offset,
	}
}
