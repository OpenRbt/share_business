package conversions

import (
	"wash_admin/internal/app"
	"wash_admin/internal/dal/dbmodels"
)

func WashUserFromDB(dbWashUser dbmodels.User) app.User {
	role := RoleSelectionApp(dbWashUser.Role)
	return app.User{
		ID:   dbWashUser.ID,
		Role: role,
	}
}

func WashUserToDB(user app.UpdateUser) dbmodels.UpdateUser {
	role := RoleSelectionDB(user.Role)
	return dbmodels.UpdateUser{
		ID:   user.ID,
		Role: role,
	}
}

func RoleSelectionApp(dbRole string) app.Role {
	role := app.AdminRole
	switch dbRole {
	case dbmodels.AdminRole:
		role = app.AdminRole
	case dbmodels.EngineerRole:
		role = app.EngineerRole
	case dbmodels.UserRole:
		role = app.UserRole
	default:
		panic("Unknown role from db, dbRole - " + dbRole)
	}
	return role
}

func RoleSelectionDB(appRole app.Role) string {
	role := dbmodels.AdminRole
	switch appRole {
	case app.AdminRole:
		role = dbmodels.AdminRole
	case app.EngineerRole:
		role = dbmodels.EngineerRole
	case app.UserRole:
		role = dbmodels.UserRole
	default:
		panic("Unknown role from app, appRole - " + appRole)
	}
	return role
}
