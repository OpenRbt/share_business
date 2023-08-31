package conversions

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
)

func PaginationToDB(entity entity.Pagination) dbmodels.Pagination {
	return dbmodels.Pagination{
		Limit:  entity.Limit,
		Offset: entity.Offset,
	}
}

func PaginationFromRest(limit, offset int64) entity.Pagination {
	return entity.Pagination{
		Limit:  limit,
		Offset: offset,
	}
}
