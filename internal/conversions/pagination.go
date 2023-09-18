package conversions

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
)

func PaginationToDB(e entities.Pagination) dbmodels.Pagination {
	return dbmodels.Pagination{
		Limit:  e.Limit,
		Offset: e.Offset,
	}
}

func PaginationFromRest(limit, offset int64) entities.Pagination {
	return entities.Pagination{
		Limit:  limit,
		Offset: offset,
	}
}
