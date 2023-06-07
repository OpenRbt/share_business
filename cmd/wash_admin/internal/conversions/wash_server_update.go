package conversions

import (
	"wash_admin/internal/app"
	"wash_admin/internal/dal/dbmodels"

	uuid "github.com/satori/go.uuid"
)

func UpdateWashServerToDb(entity app.UpdateWashServer) dbmodels.UpdateWashServer {
	return dbmodels.UpdateWashServer{
		ID:          uuid.NullUUID{UUID: entity.ID, Valid: true},
		Name:        entity.Title,
		Description: entity.Description,
	}
}
