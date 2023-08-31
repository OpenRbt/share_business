package conversions

import (
	"fmt"
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/openapi/models"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func OrganizationFromDB(organization dbmodels.Organization) entity.Organization {
	return entity.Organization{
		ID:          organization.ID,
		Name:        organization.Name,
		Description: organization.Description,
		IsDefault:   organization.IsDefault,
		Deleted:     organization.Deleted,
	}
}

func OrganizationToRest(e entity.Organization) *models.Organization {
	return &models.Organization{
		ID:          strfmt.UUID(e.ID.String()),
		Name:        e.Name,
		Description: e.Description,
		IsDefault:   e.IsDefault,
	}
}

func OrganizationsFromDB(organizations []dbmodels.Organization) []entity.Organization {
	res := make([]entity.Organization, len(organizations))

	for i, value := range organizations {
		res[i] = OrganizationFromDB(value)
	}

	return res
}

func OrganizationsToRest(organizations []entity.Organization) []*models.Organization {
	res := make([]*models.Organization, len(organizations))

	for i, value := range organizations {
		res[i] = OrganizationToRest(value)
	}

	return res
}

func OrganizationUpdateToDb(entity entity.OrganizationUpdate) dbmodels.OrganizationUpdate {
	return dbmodels.OrganizationUpdate{
		Name:        entity.Name,
		Description: entity.Description,
	}
}

func OrganizationUpdateFromRest(model models.OrganizationUpdate) entity.OrganizationUpdate {
	return entity.OrganizationUpdate{
		Name:        &model.Name,
		Description: &model.Description,
	}
}

func OrganizationCreationToDb(entity entity.OrganizationCreation) dbmodels.OrganizationCreation {
	return dbmodels.OrganizationCreation{
		Name:        entity.Name,
		Description: entity.Description,
	}
}

func OrganizationCreationFromRest(model models.OrganizationCreation) entity.OrganizationCreation {
	return entity.OrganizationCreation{
		Name:        *model.Name,
		Description: *model.Description,
	}
}

func OrganizationFilterFromRest(pagination entity.Pagination, isManagedByMe bool, organizationIDs []strfmt.UUID) (entity.OrganizationFilter, error) {
	ids := make([]uuid.UUID, 0, len(organizationIDs))

	for _, value := range organizationIDs {
		id, err := uuid.FromString(value.String())
		if err != nil {
			return entity.OrganizationFilter{}, fmt.Errorf("failed to convert UUID: %w", entity.ErrBadRequest)
		}
		ids = append(ids, id)
	}

	return entity.OrganizationFilter{
		OrganizationIDs: ids,
		Pagination:      pagination,
		IsManagedByMe:   isManagedByMe,
	}, nil
}

func OrganizationFilterToDB(filter entity.OrganizationFilter) dbmodels.OrganizationFilter {
	return dbmodels.OrganizationFilter{
		Pagination:      PaginationToDB(filter.Pagination),
		OrganizationIDs: filter.OrganizationIDs,
		IsManagedByMe:   filter.IsManagedByMe,
	}
}
