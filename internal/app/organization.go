package app

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
)

type (
	OrganizationController interface {
		Get(ctx Ctx, auth Auth, filter entity.OrganizationFilter) ([]entity.Organization, error)
		GetById(ctx Ctx, auth Auth, id uuid.UUID) (entity.Organization, error)
		Create(ctx Ctx, auth Auth, ent entity.OrganizationCreation) (entity.Organization, error)
		Update(ctx Ctx, auth Auth, id uuid.UUID, ent entity.OrganizationUpdate) (entity.Organization, error)
		Delete(ctx Ctx, auth Auth, id uuid.UUID) error

		AssignManager(ctx Ctx, auth Auth, organizationID uuid.UUID, userID string) error
		RemoveManager(ctx Ctx, auth Auth, organizationID uuid.UUID, userID string) error

		GetSettingsForOrganization(ctx Ctx, auth Auth, organizationID uuid.UUID) (entity.OrganizationSettings, error)
		UpdateSettingsForOrganization(ctx Ctx, auth Auth, organizationID uuid.UUID, e entity.OrganizationSettingsUpdate) (entity.OrganizationSettings, error)
	}

	OrganizationService interface {
		Get(ctx Ctx, userID string, filter entity.OrganizationFilter) ([]entity.Organization, error)
		GetById(ctx Ctx, id uuid.UUID) (entity.Organization, error)
		Create(ctx Ctx, ent entity.OrganizationCreation) (entity.Organization, error)
		Update(ctx Ctx, id uuid.UUID, ent entity.OrganizationUpdate) (entity.Organization, error)
		Delete(ctx Ctx, id uuid.UUID) error

		AssignManager(ctx Ctx, organizationID uuid.UUID, userID string) error
		RemoveManager(ctx Ctx, organizationID uuid.UUID, userID string) error
		IsUserManager(ctx Ctx, organizationID uuid.UUID, userID string) (bool, error)

		GetSettingsForOrganization(ctx Ctx, organizationID uuid.UUID) (entity.OrganizationSettings, error)
		UpdateSettingsForOrganization(ctx Ctx, organizationID uuid.UUID, e entity.OrganizationSettingsUpdate) (entity.OrganizationSettings, error)
	}

	OrganizationRepo interface {
		Get(ctx Ctx, userID string, filter dbmodels.OrganizationFilter) ([]dbmodels.Organization, error)
		GetById(ctx Ctx, id uuid.UUID) (dbmodels.Organization, error)
		Create(ctx Ctx, model dbmodels.OrganizationCreation) (dbmodels.Organization, error)
		Update(ctx Ctx, id uuid.UUID, model dbmodels.OrganizationUpdate) (dbmodels.Organization, error)
		Delete(ctx Ctx, id uuid.UUID) error

		AssignManager(ctx Ctx, organizationID uuid.UUID, userID string) error
		RemoveManager(ctx Ctx, organizationID uuid.UUID, userID string) error
		IsUserManager(ctx Ctx, organizationID uuid.UUID, userID string) (bool, error)

		GetSettingsForOrganization(ctx Ctx, organizationID uuid.UUID) (dbmodels.OrganizationSettings, error)
		UpdateSettingsForOrganization(ctx Ctx, organizationID uuid.UUID, model dbmodels.OrganizationSettingsUpdate) (dbmodels.OrganizationSettings, error)
	}
)
