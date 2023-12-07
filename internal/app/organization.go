package app

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"

	uuid "github.com/satori/go.uuid"
)

type (
	OrganizationController interface {
		Get(ctx Ctx, auth AdminAuth, filter entities.OrganizationFilter) ([]entities.Organization, error)
		GetById(ctx Ctx, auth AdminAuth, id uuid.UUID) (entities.Organization, error)
		Create(ctx Ctx, auth AdminAuth, ent entities.OrganizationCreation) (entities.Organization, error)
		Update(ctx Ctx, auth AdminAuth, id uuid.UUID, ent entities.OrganizationUpdate) (entities.Organization, error)
		Delete(ctx Ctx, auth AdminAuth, id uuid.UUID) error

		AssignManager(ctx Ctx, auth AdminAuth, organizationID uuid.UUID, userID string) error
		RemoveManager(ctx Ctx, auth AdminAuth, organizationID uuid.UUID, userID string) error
	}

	OrganizationService interface {
		Get(ctx Ctx, filter entities.OrganizationFilter) ([]entities.Organization, error)
		GetAll(ctx Ctx, pagination entities.Pagination) ([]entities.Organization, error)
		GetById(ctx Ctx, id uuid.UUID) (entities.Organization, error)
		Create(ctx Ctx, ent entities.OrganizationCreation) (entities.Organization, error)
		Update(ctx Ctx, id uuid.UUID, ent entities.OrganizationUpdate) (entities.Organization, error)
		Delete(ctx Ctx, id uuid.UUID) error

		AssignManager(ctx Ctx, organizationID uuid.UUID, userID string) error
		RemoveManager(ctx Ctx, organizationID uuid.UUID, userID string) error

		GetDefaultGroupByOrganizationId(ctx Ctx, id uuid.UUID) (entities.ServerGroup, error)
		GetAdminUsersByOrganizationID(ctx Ctx, id uuid.UUID) ([]entities.AdminUser, error)
		GetDeletedByID(ctx Ctx, id uuid.UUID) (entities.Organization, error)
	}

	OrganizationRepo interface {
		Get(ctx Ctx, filter dbmodels.OrganizationFilter) ([]dbmodels.Organization, error)
		GetAll(ctx Ctx, pagination dbmodels.Pagination) ([]dbmodels.Organization, error)
		GetById(ctx Ctx, id uuid.UUID) (dbmodels.Organization, error)
		Create(ctx Ctx, model dbmodels.OrganizationCreation) (dbmodels.Organization, error)
		Update(ctx Ctx, id uuid.UUID, model dbmodels.OrganizationUpdate) (dbmodels.Organization, error)
		Delete(ctx Ctx, id uuid.UUID) error

		AssignManager(ctx Ctx, organizationID uuid.UUID, userID string) error
		RemoveManager(ctx Ctx, organizationID uuid.UUID, userID string) error

		GetDefaultGroupByOrganizationId(ctx Ctx, id uuid.UUID) (dbmodels.ServerGroup, error)
		GetAdminUsersByOrganizationID(ctx Ctx, id uuid.UUID) ([]dbmodels.AdminUser, error)
		GetDeletedByID(ctx Ctx, id uuid.UUID) (dbmodels.Organization, error)
	}
)
