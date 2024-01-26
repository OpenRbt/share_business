package conversions

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
	rabbitEntities "washbonus/internal/infrastructure/rabbit/entities"
	"washbonus/openapi/admin/models"

	"github.com/go-openapi/strfmt"
)

func AdminUserFromDb(dbUser dbmodels.AdminUser) entities.AdminUser {
	var org *entities.AdminOrganization
	if dbUser.OrganizationID != nil {
		org = &entities.AdminOrganization{
			ID:          *dbUser.OrganizationID,
			Name:        *dbUser.OrganizationName,
			DisplayName: *dbUser.OrganizationDisplayName,
			Description: *dbUser.OrganizationDescription,
			Deleted:     *dbUser.OrganizationDeleted,
		}
	}

	return entities.AdminUser{
		ID:           dbUser.ID,
		Name:         dbUser.Name,
		Email:        dbUser.Email,
		Role:         RoleSelectionApp(dbUser.Role),
		Organization: org,
		Version:      dbUser.Version,
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
		ID:    user.ID,
		Role:  models.AdminUserRole(user.Role),
		Name:  user.Name,
		Email: strfmt.Email(user.Email),
	}

	if user.Organization != nil {
		orgId := user.Organization.ID.String()
		mod.Organization = &models.AdminUserOrganization{
			ID:          strfmt.UUID(orgId),
			Name:        user.Organization.Name,
			DisplayName: user.Organization.DisplayName,
			Description: user.Organization.Description,
			Deleted:     user.Organization.Deleted,
		}
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

func AdminUserRoleUpdateFromRest(id string, role models.AdminUserRole) (entities.AdminUserRoleUpdate, error) {
	r, err := RoleSelectionRest(role)
	if err != nil {
		return entities.AdminUserRoleUpdate{}, err
	}

	return entities.AdminUserRoleUpdate{
		ID:   id,
		Role: r,
	}, nil
}

func AdminUserFilterFromRest(limit, offset int64, role *string, isBlocked *bool) (entities.AdminUserFilter, error) {
	filter := entities.AdminUserFilter{
		Pagination: PaginationFromRest(limit, offset),
		IsBlocked:  isBlocked,
	}

	if role == nil {
		return filter, nil
	}

	r, err := RoleSelectionRest(models.AdminUserRole(*role))
	if err != nil {
		return entities.AdminUserFilter{}, err
	}

	filter.Role = &r

	return filter, nil
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

func AdminUserToRabbit(user entities.AdminUser) rabbitEntities.AdminUser {
	ent := rabbitEntities.AdminUser{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Role:    string(RoleSelectionDB(user.Role)),
		Version: user.Version,
	}

	if user.Organization != nil {
		orgID := user.Organization.ID.String()
		ent.OrganizationID = &orgID
	}

	return ent
}

func AdminUsersToRabbit(users []entities.AdminUser) []rabbitEntities.AdminUser {
	res := make([]rabbitEntities.AdminUser, len(users))

	for idx, user := range users {
		res[idx] = AdminUserToRabbit(user)
	}

	return res
}

func RoleSelectionRest(role models.AdminUserRole) (entities.Role, error) {
	var entRole entities.Role
	var err error

	switch role {
	case models.AdminUserRoleAdmin:
		entRole = entities.AdminRole
	case models.AdminUserRoleSystemManager:
		entRole = entities.SystemManagerRole
	case models.AdminUserRoleNoAccess:
		entRole = entities.NoAccessRole
	default:
		err = entities.ErrInvalidRole
	}

	return entRole, err
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
