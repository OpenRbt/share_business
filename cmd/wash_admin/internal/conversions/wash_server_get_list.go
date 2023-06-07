package conversions

import (
	"wash_admin/internal/app"
	"wash_admin/internal/dal/dbmodels"
)

func PaginationToDB(entity app.Pagination) dbmodels.Pagination {
	return dbmodels.Pagination{
		Limit:  entity.Limit,
		Offset: entity.Offset,
	}
}

func WashServerListFromDB(washServerList []dbmodels.WashServer) []app.WashServer {
	res := make([]app.WashServer, len(washServerList))

	for i, value := range washServerList {
		res[i] = WashServerFromDB(value)
	}

	return res
}
