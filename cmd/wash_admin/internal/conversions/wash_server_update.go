package conversions

import (
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity/vo"
	"wash_admin/openapi/models"

	uuid "github.com/satori/go.uuid"
)

func UpdateWashServerToDb(entity vo.UpdateWashServer) dbmodels.UpdateWashServer {
	return dbmodels.UpdateWashServer{
		ID:          uuid.NullUUID{UUID: entity.ID, Valid: true},
		Name:        entity.Title,
		Description: entity.Description,
	}
}

func UpdateWashServerFromRest(model models.WashServerUpdate) (vo.UpdateWashServer, error) {
	id, err := uuid.FromString(model.ID)

	if err != nil {
		return vo.UpdateWashServer{}, err
	}

	return vo.UpdateWashServer{
		ID:          id,
		Title:       &model.Name,
		Description: &model.Description,
	}, nil
}
