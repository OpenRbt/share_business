package conversions

import (
	"fmt"
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
	"washbonus/openapi/admin/models"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func OrganizationFromDB(organization dbmodels.Organization) entities.Organization {
	return entities.Organization{
		ID:                            organization.ID,
		Name:                          organization.Name,
		DisplayName:                   organization.DisplayName,
		Description:                   organization.Description,
		IsDefault:                     organization.IsDefault,
		ReportsProcessingDelayMinutes: organization.ReportsProcessingDelayMinutes,
		BonusPercentage:               organization.BonusPercentage,
		Deleted:                       organization.Deleted,
	}
}

func OrganizationToRest(e entities.Organization) *models.Organization {
	return &models.Organization{
		ID:                            strfmt.UUID(e.ID.String()),
		Name:                          e.Name,
		DisplayName:                   e.DisplayName,
		Description:                   e.Description,
		IsDefault:                     e.IsDefault,
		ReportsProcessingDelayMinutes: &e.ReportsProcessingDelayMinutes,
		BonusPercentage:               &e.BonusPercentage,
	}
}

func OrganizationsFromDB(organizations []dbmodels.Organization) []entities.Organization {
	res := make([]entities.Organization, len(organizations))

	for i, value := range organizations {
		res[i] = OrganizationFromDB(value)
	}

	return res
}

func OrganizationsToRest(organizations []entities.Organization) []*models.Organization {
	res := make([]*models.Organization, len(organizations))

	for i, value := range organizations {
		res[i] = OrganizationToRest(value)
	}

	return res
}

func OrganizationUpdateToDb(e entities.OrganizationUpdate) dbmodels.OrganizationUpdate {
	return dbmodels.OrganizationUpdate{
		Name:                          e.Name,
		Description:                   e.Description,
		DisplayName:                   e.DisplayName,
		ReportsProcessingDelayMinutes: e.ReportsProcessingDelayMinutes,
		BonusPercentage:               e.BonusPercentage,
	}
}

func OrganizationUpdateFromRest(model models.OrganizationUpdate) entities.OrganizationUpdate {
	return entities.OrganizationUpdate{
		Name:                          model.Name,
		DisplayName:                   model.DisplayName,
		Description:                   model.Description,
		ReportsProcessingDelayMinutes: model.ReportsProcessingDelayMinutes,
		BonusPercentage:               model.BonusPercentage,
	}
}

func OrganizationCreationToDb(e entities.OrganizationCreation) dbmodels.OrganizationCreation {
	mod := dbmodels.OrganizationCreation{
		Name:            e.Name,
		DisplayName:     e.DisplayName,
		Description:     e.Description,
		BonusPercentage: e.BonusPercentage,
	}

	if e.ReportsProcessingDelayMinutes != nil {
		processingDelayMinutes := fmt.Sprintf("%d minutes", *e.ReportsProcessingDelayMinutes)
		mod.ReportsProcessingDelayMinutes = &processingDelayMinutes
	}

	return mod
}

func OrganizationCreationFromRest(model models.OrganizationCreation) entities.OrganizationCreation {
	return entities.OrganizationCreation{
		Name:                          *model.Name,
		DisplayName:                   model.DisplayName,
		Description:                   *model.Description,
		ReportsProcessingDelayMinutes: model.ReportsProcessingDelayMinutes,
		BonusPercentage:               model.BonusPercentage,
	}
}

func OrganizationFilterFromRest(pagination entities.Pagination, organizationIDs []strfmt.UUID) (entities.OrganizationFilter, error) {
	ids := make([]uuid.UUID, 0, len(organizationIDs))

	for _, value := range organizationIDs {
		id, err := uuid.FromString(value.String())
		if err != nil {
			return entities.OrganizationFilter{}, fmt.Errorf("%w: wrong uuid", entities.ErrBadRequest)
		}
		ids = append(ids, id)
	}

	return entities.OrganizationFilter{
		OrganizationIDs: ids,
		Pagination:      pagination,
	}, nil
}

func OrganizationFilterToDB(filter entities.OrganizationFilter) dbmodels.OrganizationFilter {
	return dbmodels.OrganizationFilter{
		Pagination:      PaginationToDB(filter.Pagination),
		OrganizationIDs: filter.OrganizationIDs,
	}
}
