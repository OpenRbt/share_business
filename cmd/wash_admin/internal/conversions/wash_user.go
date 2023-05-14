package conversions

import (
	"wash_admin/internal/app/role"
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity"
)

func WashUserFromDB(dbWashUser dbmodels.WashUser) entity.WashUser {
	return entity.WashUser{
		ID:       dbWashUser.ID.UUID,
		Identity: dbWashUser.Identity,
		Role:     role.Role(dbWashUser.Role),
	}
}
