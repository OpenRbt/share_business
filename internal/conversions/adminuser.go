package conversions

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
	"washbonus/openapi/admin/models"

	"github.com/go-openapi/strfmt"
)

func AdminUserFromDb(dbUser dbmodels.AdminUser) entities.AdminUser {
	return entities.AdminUser{
		ID:             dbUser.ID,
		Name:           dbUser.Name,
		Email:          dbUser.Email,
		Role:           RoleSelectionApp(dbUser.Role),
		OrganizationID: dbUser.OrganizationID,
	}
}

func AdminUsersFromDb(dbUsers []dbmodels.AdminUser) []entities.AdminUser {
	users := make([]entities.AdminUser, len(dbUsers))

	for idx, user := range dbUsers {
		users[idx] = AdminUserFromDb(user)
	}

	return users
}

func AdminUserToRest(user entities.AdminUser) models.AdminUser {
	mod := models.AdminUser{
		ID:   user.ID,
		Role: models.AdminUserRole(user.Role),
	}

	if user.Email != nil {
		mod.Email = strfmt.Email(*user.Email)
	}

	if user.Name != nil {
		mod.Name = *user.Name
	}

	if user.OrganizationID != nil {
		orgId := user.OrganizationID.String()
		mod.OrganizationID = (*strfmt.UUID)(&orgId)
	}

	return mod
}

func AdminUsersToRest(users []entities.AdminUser) []*models.AdminUser {
	models := make([]*models.AdminUser, len(users))

	for idx, user := range users {
		rest := AdminUserToRest(user)
		models[idx] = &rest
	}

	return models
}

func AdminUserUpdateToDB(user entities.AdminUserUpdate) dbmodels.AdminUserUpdate {
	return dbmodels.AdminUserUpdate{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
}

func AdminUserCreationToDB(e entities.AdminUserCreation) dbmodels.AdminUserCreation {
	return dbmodels.AdminUserCreation{
		ID:             e.ID,
		Email:          e.Email,
		Name:           e.Name,
		OrganizationId: e.OrganizationID,
	}
}

func AdminUserRoleUpdateToDB(e entities.AdminUserRoleUpdate) dbmodels.AdminUserRoleUpdate {
	return dbmodels.AdminUserRoleUpdate{
		ID:   e.ID,
		Role: RoleSelectionDB(e.Role),
	}
}

func AdminUserRoleUpdateFromRest(id string, role models.AdminUserRole) entities.AdminUserRoleUpdate {
	return entities.AdminUserRoleUpdate{
		ID:   id,
		Role: RoleSelectionRest(role),
	}
}

func AdminUserFilterFromRest(limit, offset int64, role *string, isBlocked *bool) entities.AdminUserFilter {
	filter := entities.AdminUserFilter{
		Pagination: PaginationFromRest(limit, offset),
		IsBlocked:  isBlocked,
	}

	if role == nil {
		return filter
	}

	r := RoleSelectionRest(models.AdminUserRole(*role))
	filter.Role = &r

	return filter
}

func AdminUserFilterToDB(filter entities.AdminUserFilter) dbmodels.AdminUserFilter {
	dbFilter := dbmodels.AdminUserFilter{
		Pagination: PaginationToDB(filter.Pagination),
		IsBlocked:  filter.IsBlocked,
	}

	if filter.Role == nil {
		return dbFilter
	}

	r := RoleSelectionDB(*filter.Role)
	dbFilter.Role = &r

	return dbFilter
}

func RoleSelectionRest(role models.AdminUserRole) entities.Role {
	switch role {
	case models.AdminUserRoleAdmin:
		return entities.AdminRole
	case models.AdminUserRoleSystemManager:
		return entities.SystemManagerRole
	case models.AdminUserRoleNoAccess:
		return entities.NoAccessRole
	default:
		panic("Unknown rest role, restRole - " + role)
	}
}

func RoleSelectionApp(role dbmodels.Role) entities.Role {
	switch role {
	case dbmodels.AdminRole:
		return entities.AdminRole
	case dbmodels.SystemManagerRole:
		return entities.SystemManagerRole
	case dbmodels.NoAccessRole:
		return entities.NoAccessRole
	default:
		panic("Unknown db role, dbRole - " + role)
	}
}

func RoleSelectionDB(appRole entities.Role) dbmodels.Role {
	switch appRole {
	case entities.AdminRole:
		return dbmodels.AdminRole
	case entities.SystemManagerRole:
		return dbmodels.SystemManagerRole
	case entities.NoAccessRole:
		return dbmodels.NoAccessRole
	default:
		panic("Unknown app role, appRole - " + appRole)
	}
}
