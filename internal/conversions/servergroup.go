package conversions

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
	"washbonus/openapi/admin/models"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func ServerGroupFromDB(group dbmodels.ServerGroup) entities.ServerGroup {
	return entities.ServerGroup{
		ID:             group.ID,
		Name:           group.Name,
		Description:    group.Description,
		OrganizationID: group.OrganizationID,
		IsDefault:      group.IsDefault,
		Deleted:        group.Deleted,
	}
}

func ServerGroupToRest(group entities.ServerGroup) *models.ServerGroup {
	return &models.ServerGroup{
		ID:             strfmt.UUID(group.ID.String()),
		Name:           group.Name,
		Description:    group.Description,
		OrganizationID: strfmt.UUID(group.OrganizationID.String()),
		IsDefault:      group.IsDefault,
	}
}

func ServerGroupsFromDB(groups []dbmodels.ServerGroup) []entities.ServerGroup {
	res := make([]entities.ServerGroup, len(groups))

	for i, value := range groups {
		res[i] = ServerGroupFromDB(value)
	}

	return res
}

func ServerGroupsToRest(groups []entities.ServerGroup) []*models.ServerGroup {
	res := make([]*models.ServerGroup, len(groups))

	for i, value := range groups {
		rest := ServerGroupToRest(value)
		res[i] = rest
	}

	return res
}

func ServerGroupUpdateToDb(groupUpdate entities.ServerGroupUpdate) dbmodels.ServerGroupUpdate {
	return dbmodels.ServerGroupUpdate{
		Name:        groupUpdate.Name,
		Description: groupUpdate.Description,
	}
}

func ServerGroupUpdateFromRest(groupUpdate models.ServerGroupUpdate) entities.ServerGroupUpdate {
	return entities.ServerGroupUpdate{
		Name:        groupUpdate.Name,
		Description: groupUpdate.Description,
	}
}

func ServerGroupCreationToDb(groupCreation entities.ServerGroupCreation) dbmodels.ServerGroupCreation {
	return dbmodels.ServerGroupCreation{
		Name:           groupCreation.Name,
		Description:    groupCreation.Description,
		OrganizationID: groupCreation.OrganizationID,
	}
}

func ServerGroupCreationFromRest(model models.ServerGroupCreation) entities.ServerGroupCreation {
	return entities.ServerGroupCreation{
		Name:           *model.Name,
		Description:    *model.Description,
		OrganizationID: uuid.FromStringOrNil((*model.OrganizationID).String()),
	}
}

func ServerGroupFilterFromRest(pagination entities.Pagination, organizationID *strfmt.UUID) entities.ServerGroupFilter {
	filter := entities.ServerGroupFilter{
		Pagination: pagination,
	}

	if organizationID != nil {
		orgID, err := uuid.FromString(organizationID.String())
		if err == nil {
			filter.OrganizationID = &orgID
		}
	}

	return filter
}

func ServerGroupFilterToDB(filter entities.ServerGroupFilter) dbmodels.ServerGroupFilter {
	return dbmodels.ServerGroupFilter{
		Pagination:     PaginationToDB(filter.Pagination),
		OrganizationID: filter.OrganizationID,
	}
}
