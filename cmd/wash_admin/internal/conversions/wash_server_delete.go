package conversions

import (
	"wash_admin/internal/dal/dbmodels"

	uuid "github.com/satori/go.uuid"
)

func DeleteWashServerToDB(id uuid.UUID) dbmodels.DeleteWashServer {
	return dbmodels.DeleteWashServer{
		ID: uuid.NullUUID{UUID: id, Valid: true},
	}
}
