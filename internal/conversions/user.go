package conversions

import (
	"encoding/json"
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/openapi/models"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func UserFromDb(dbUser dbmodels.User) entity.User {
	ent := entity.User{
		ID:      dbUser.ID,
		Email:   dbUser.Email,
		Role:    RoleSelectionApp(dbUser.Role),
		Deleted: dbUser.Deleted,
	}

	var jsonOrgIDs []string
	json.Unmarshal([]byte(dbUser.OrganizationIDsRaw), &jsonOrgIDs)

	orgIDs := make([]uuid.UUID, 0, len(jsonOrgIDs))

	for _, id := range jsonOrgIDs {
		parsedID, err := uuid.FromString(id)
		if err != nil {
			continue
		}
		orgIDs = append(orgIDs, parsedID)
	}
	ent.OrganizationIDs = orgIDs

	return ent
}

func UsersFromDb(dbUsers []dbmodels.User) []entity.User {
	users := make([]entity.User, len(dbUsers))

	for idx, user := range dbUsers {
		users[idx] = UserFromDb(user)
	}

	return users
}

func UserToRest(user entity.User) models.User {
	mod := models.User{
		ID:   user.ID,
		Role: string(user.Role),
	}

	if user.Email != nil {
		mod.Email = strfmt.Email(*user.Email)
	}

	if len(user.OrganizationIDs) <= 0 {
		return mod
	}

	orgIDs := make([]strfmt.UUID, len(user.OrganizationIDs))

	for idx, id := range user.OrganizationIDs {
		orgIDs[idx] = strfmt.UUID(id.String())
	}
	mod.OrganizationIds = orgIDs

	return mod
}

func UsersToRest(users []entity.User) []*models.User {
	models := make([]*models.User, len(users))

	for idx, user := range users {
		rest := UserToRest(user)
		models[idx] = &rest
	}

	return models
}

func UserUpdateRoleToDB(user entity.UserUpdateRole) dbmodels.UserUpdateRole {
	role := RoleSelectionDB(user.Role)
	return dbmodels.UserUpdateRole{
		ID:   user.ID,
		Role: role,
	}
}

func UserUpdateToDB(user entity.UserUpdate) dbmodels.UserUpdate {
	return dbmodels.UserUpdate{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
}

func UserCreationToDB(e entity.UserCreation) dbmodels.UserCreation {
	return dbmodels.UserCreation{
		ID:    e.ID,
		Email: e.Email,
		Name:  e.Name,
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
