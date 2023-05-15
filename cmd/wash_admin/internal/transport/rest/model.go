package rest

import (
	"wash_admin/internal/app"
	"wash_admin/openapi/models"

	uuid "github.com/satori/go.uuid"
)

func DeleteWashServerFromRest(model models.WashServerDelete) (uuid.UUID, error) {
	id, err := uuid.FromString(*model.ID)

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func PaginationFromRest(model models.Pagination) app.Pagination {
	return app.Pagination{
		Limit:  model.Limit,
		Offset: model.Offset,
	}
}

func WashServerListToRest(washServerEntity []app.WashServer) []*models.WashServer {
	res := make([]*models.WashServer, len(washServerEntity))

	for i, value := range washServerEntity {
		rest := WashServerToRest(value)
		res[i] = rest
	}

	return res
}

func UpdateWashServerFromRest(model models.WashServerUpdate) (app.UpdateWashServer, error) {
	id, err := uuid.FromString(model.ID)

	if err != nil {
		return app.UpdateWashServer{}, err
	}

	return app.UpdateWashServer{
		ID:          id,
		Title:       &model.Name,
		Description: &model.Description,
	}, nil
}

func WashServerToRest(washServer app.WashServer) *models.WashServer {
	return &models.WashServer{
		ID:          washServer.ID.String(),
		Name:        washServer.Title,
		Description: washServer.Description,
		ServiceKey:  washServer.ServiceKey,
	}
}

func RegisterWashServerFromRest(rest models.WashServerAdd) app.RegisterWashServer {
	return app.RegisterWashServer{
		Title:       *rest.Name,
		Description: rest.Description,
	}
}
