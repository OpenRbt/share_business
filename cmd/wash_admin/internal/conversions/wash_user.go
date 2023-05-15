package conversions

import (
	"wash_admin/internal/app"
	"wash_admin/internal/dal/dbmodels"
)

func WashUserFromDB(dbWashUser dbmodels.WashUser) app.WashUser {
	r := app.AdminRole
	switch dbWashUser.Role {
	case string(app.AdminRole):
		r = app.AdminRole
	case string(app.EngineerRole):
		r = app.EngineerRole
	case string(app.UserRole):
		r = app.UserRole
	default:
		panic(1)
	}
	return app.WashUser{
		ID:       dbWashUser.ID.UUID,
		Identity: dbWashUser.Identity,
		Role:     r,
	}
}
