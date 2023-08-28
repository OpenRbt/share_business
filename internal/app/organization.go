package app

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
)

type (
	OrganizationController interface {
		Get(ctx Ctx, authUser entity.User, filter entity.OrganizationFilter) ([]entity.Organization, error)
		GetById(ctx Ctx, authUser entity.User, id uuid.UUID) (entity.Organization, error)
		Create(ctx Ctx, authUser entity.User, ent entity.OrganizationCreation) (entity.Organization, error)
		Update(ctx Ctx, authUser entity.User, id uuid.UUID, ent entity.OrganizationUpdate) (entity.Organization, error)
		Delete(ctx Ctx, authUser entity.User, id uuid.UUID) error

		AssignManager(ctx Ctx, authUser entity.User, organizationID uuid.UUID, userID string) error
		RemoveManager(ctx Ctx, authUser entity.User, organizationID uuid.UUID, userID string) error
	}

	OrganizationService interface {
		Get(ctx Ctx, filter entity.OrganizationFilter) ([]entity.Organization, error)
		GetForManager(ctx Ctx, userID string, filter entity.OrganizationFilter) ([]entity.Organization, error)
		GetById(ctx Ctx, id uuid.UUID) (entity.Organization, error)
		Create(ctx Ctx, ent entity.OrganizationCreation) (entity.Organization, error)
		Update(ctx Ctx, id uuid.UUID, ent entity.OrganizationUpdate) (entity.Organization, error)
		Delete(ctx Ctx, id uuid.UUID) error

		AssignManager(ctx Ctx, organizationID uuid.UUID, userID string) error
		RemoveManager(ctx Ctx, organizationID uuid.UUID, userID string) error
		IsUserManager(ctx Ctx, organizationID uuid.UUID, userID string) (bool, error)
	}

	OrganizationRepo interface {
		Get(ctx Ctx, filter dbmodels.OrganizationFilter) ([]dbmodels.Organization, error)
		GetForManager(ctx Ctx, userID string, filter dbmodels.OrganizationFilter) ([]dbmodels.Organization, error)
		GetById(ctx Ctx, id uuid.UUID) (dbmodels.Organization, error)
		Create(ctx Ctx, model dbmodels.OrganizationCreation) (dbmodels.Organization, error)
		Update(ctx Ctx, id uuid.UUID, model dbmodels.OrganizationUpdate) (dbmodels.Organization, error)
		Delete(ctx Ctx, id uuid.UUID) error

		AssignManager(ctx Ctx, organizationID uuid.UUID, userID string) error
		RemoveManager(ctx Ctx, organizationID uuid.UUID, userID string) error
		IsUserManager(ctx Ctx, organizationID uuid.UUID, userID string) (bool, error)
	}
)
