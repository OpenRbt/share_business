package conversions

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/openapi/models"
)

func WashServerFromDB(dbWashServer dbmodels.WashServer) entity.WashServer {
	return entity.WashServer{
		ID:          dbWashServer.ID.UUID,
		Title:       dbWashServer.Title,
		Description: dbWashServer.Description,
		CreatedBy:   dbWashServer.CreatedBy,
		ServiceKey:  dbWashServer.ServiceKey,
	}
}

func WashServerToRest(e entity.WashServer) *models.WashServer {
	return &models.WashServer{
		ID:          e.ID.String(),
		Description: e.Description,
		Name:        e.Title,
	}
}

func WashServerToAdminRest(e entity.WashServer) *models.WashServer {
	return &models.WashServer{
		ID:          e.ID.String(),
		Name:        e.Title,
		Description: e.Description,
		CreatedBy:   e.CreatedBy,
		ServiceKey:  e.ServiceKey,
	}
}

func PaginationToDB(entity entity.Pagination) dbmodels.Pagination {
	return dbmodels.Pagination{
		Limit:  entity.Limit,
		Offset: entity.Offset,
	}
}

func WashServerListFromDB(washServerList []dbmodels.WashServer) []entity.WashServer {
	res := make([]entity.WashServer, len(washServerList))

	for i, value := range washServerList {
		res[i] = WashServerFromDB(value)
	}

	return res
}

func UpdateWashServerToDb(entity entity.UpdateWashServer) dbmodels.UpdateWashServer {
	return dbmodels.UpdateWashServer{
		Name:        entity.Title,
		Description: entity.Description,
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

func WashServerListToRest(washServerEntity []entity.WashServer) []*models.WashServer {
	res := make([]*models.WashServer, len(washServerEntity))

	for i, value := range washServerEntity {
		rest := WashServerToRest(value)
		res[i] = rest
	}

	return res
}

func UpdateWashServerFromRest(model models.WashServerUpdate) entity.UpdateWashServer {
	return entity.UpdateWashServer{
		Title:       &model.Name,
		Description: &model.Description,
	}
}

func CreateWashServerFromRest(rest models.WashServerCreation) entity.CreateWashServer {
	return entity.CreateWashServer{
		Title:       *rest.Name,
		Description: rest.Description,
	}
}
