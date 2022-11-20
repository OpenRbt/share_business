package conversions

import (
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity"
)

func WashAdminFromDB(dbWashAdmin dbmodels.WashAdmin) entity.WashAdmin {
	return entity.WashAdmin{
		ID: dbWashAdmin.ID.UUID,
		Identity: dbWashAdmin.Identity,
	}
}