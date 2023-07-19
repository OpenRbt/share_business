package conversions

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/openapi/models"
)

func UserFromDb(dbUser dbmodels.User) entity.User {
	return entity.User{
		Balance: dbUser.Balance.Decimal,
		Role:    entity.Role(dbUser.Role),
		Deleted: dbUser.Deleted,
		ID:      dbUser.ID,
	}
}

func UserToRest(user entity.User) models.User {
	return models.User{
		Balance: user.Balance.IntPart(),
		Role:    string(user.Role),
		ID:      user.ID,
	}
}

func WashUserFromDB(dbWashUser dbmodels.User) entity.User {
	role := RoleSelectionApp(dbWashUser.Role)
	return entity.User{
		ID:   dbWashUser.ID,
		Role: role,
	}
}

func WashUserToDB(user entity.UpdateUser) dbmodels.UpdateUser {
	role := RoleSelectionDB(user.Role)
	return dbmodels.UpdateUser{
		ID:   user.ID,
		Role: role,
	}
}

func RoleSelectionApp(dbRole string) entity.Role {
	role := entity.AdminRole
	switch dbRole {
	case dbmodels.AdminRole:
		role = entity.AdminRole
	case dbmodels.EngineerRole:
		role = entity.EngineerRole
	case dbmodels.UserRole:
		role = entity.UserRole
	default:
		panic("Unknown role from db, dbRole - " + dbRole)
	}
	return role
}

func RoleSelectionDB(appRole entity.Role) string {
	role := dbmodels.AdminRole
	switch appRole {
	case entity.AdminRole:
		role = dbmodels.AdminRole
	case entity.EngineerRole:
		role = dbmodels.EngineerRole
	case entity.UserRole:
		role = dbmodels.UserRole
	default:
		panic("Unknown role from app, appRole - " + appRole)
	}
	return role
}
