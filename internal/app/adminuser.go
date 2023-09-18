package app

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"

	uuid "github.com/satori/go.uuid"
)

type (
	AdminController interface {
		GetById(ctx Ctx, auth AdminAuth, userID string) (entities.AdminUser, error)
		Get(ctx Ctx, auth AdminAuth, pagination entities.Pagination) ([]entities.AdminUser, error)
		UpdateRole(ctx Ctx, auth AdminAuth, userUpdate entities.AdminUserRoleUpdate) error
		Delete(ctx Ctx, auth AdminAuth, id string) error

		GetApplications(ctx Ctx, auth AdminAuth, filter entities.AdminApplicationFilter) ([]entities.AdminApplication, error)
		CreateApplication(ctx Ctx, ent entities.AdminApplicationCreation) (entities.AdminApplication, error)
		ReviewApplication(ctx Ctx, auth AdminAuth, id uuid.UUID, ent entities.AdminApplicationReview) error
	}

	AdminService interface {
		GetById(ctx Ctx, userID string) (entities.AdminUser, error)
		Get(ctx Ctx, pagination entities.Pagination) ([]entities.AdminUser, error)
		Create(ctx Ctx, userCreation entities.AdminUserCreation) (entities.AdminUser, error)
		UpdateRole(ctx Ctx, userRole entities.AdminUserRoleUpdate) error
		Update(ctx Ctx, userModel entities.AdminUserUpdate) error
		Delete(ctx Ctx, id string) error

		GetApplications(ctx Ctx, filter entities.AdminApplicationFilter) ([]entities.AdminApplication, error)
		CreateApplication(ctx Ctx, ent entities.AdminApplicationCreation) (entities.AdminApplication, error)
		ReviewApplication(ctx Ctx, id uuid.UUID, ent entities.AdminApplicationReview) error
		GetApplicationByUser(ctx Ctx, id string) (entities.AdminApplication, error)
	}

	AdminRepo interface {
		GetById(ctx Ctx, userID string) (dbmodels.AdminUser, error)
		Get(ctx Ctx, pagination dbmodels.Pagination) ([]dbmodels.AdminUser, error)
		Create(ctx Ctx, userCreation dbmodels.AdminUserCreation) (dbmodels.AdminUser, error)
		UpdateRole(ctx Ctx, userUpdate dbmodels.AdminUserRoleUpdate) error
		Update(ctx Ctx, userModel dbmodels.AdminUserUpdate) error
		Delete(ctx Ctx, id string) error

		GetApplications(ctx Ctx, filter dbmodels.AdminApplicationFilter) ([]dbmodels.AdminApplication, error)
		CreateApplication(ctx Ctx, ent dbmodels.AdminApplicationCreation) (dbmodels.AdminApplication, error)
		ReviewApplication(ctx Ctx, id uuid.UUID, ent dbmodels.AdminApplicationReview) error
		GetApplicationByUser(ctx Ctx, id string) (dbmodels.AdminApplication, error)
	}
)

func IsSystemManager(admin entities.AdminUser) bool {
	return admin.Role == entities.SystemManagerRole
}

func IsAdmin(admin entities.AdminUser) bool {
	return admin.Role == entities.AdminRole
}

func IsAdminManageOrganization(user entities.AdminUser, organizationID uuid.UUID) bool {
	return IsAdmin(user) && user.OrganizationID != nil && *user.OrganizationID == organizationID
}
