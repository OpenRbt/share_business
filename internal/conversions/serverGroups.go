package conversions

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/openapi/models"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func ServerGroupFromDB(group dbmodels.ServerGroup) entity.ServerGroup {
	return entity.ServerGroup{
		ID:             group.ID,
		Name:           group.Name,
		Description:    group.Description,
		OrganizationID: group.OrganizationID,
		IsDefault:      group.IsDefault,
		Deleted:        group.Deleted,
	}
}

func ServerGroupToRest(group entity.ServerGroup) *models.ServerGroup {
	return &models.ServerGroup{
		ID:             strfmt.UUID(group.ID.String()),
		Name:           group.Name,
		Description:    group.Description,
		OrganizationID: strfmt.UUID(group.OrganizationID.String()),
		IsDefault:      group.IsDefault,
	}
}

func ServerGroupsFromDB(groups []dbmodels.ServerGroup) []entity.ServerGroup {
	res := make([]entity.ServerGroup, len(groups))

	for i, value := range groups {
		res[i] = ServerGroupFromDB(value)
	}

	return res
}

func ServerGroupsToRest(groups []entity.ServerGroup) []*models.ServerGroup {
	res := make([]*models.ServerGroup, len(groups))

	for i, value := range groups {
		rest := ServerGroupToRest(value)
		res[i] = rest
	}

	return res
}

func ServerGroupUpdateToDb(groupUpdate entity.ServerGroupUpdate) dbmodels.ServerGroupUpdate {
	return dbmodels.ServerGroupUpdate{
		Name:        groupUpdate.Name,
		Description: groupUpdate.Description,
		IsDefault:   groupUpdate.IsDefault,
	}
}

func ServerGroupUpdateFromRest(groupUpdate models.ServerGroupUpdate) entity.ServerGroupUpdate {
	return entity.ServerGroupUpdate{
		Name:        groupUpdate.Name,
		Description: groupUpdate.Description,
		IsDefault:   groupUpdate.IsDefault,
	}
}

func ServerGroupCreationToDb(groupCreation entity.ServerGroupCreation) dbmodels.ServerGroupCreation {
	return dbmodels.ServerGroupCreation{
		Name:           groupCreation.Name,
		Description:    groupCreation.Description,
		OrganizationID: groupCreation.OrganizationID,
	}
}

func ServerGroupCreationFromRest(model models.ServerGroupCreation) entity.ServerGroupCreation {
	return entity.ServerGroupCreation{
		Name:           *model.Name,
		Description:    *model.Description,
		OrganizationID: uuid.FromStringOrNil((*model.OrganizationID).String()),
	}
}

func ServerGroupFilterFromRest(pagination entity.Pagination, isManagedByMe bool, organizationID *strfmt.UUID) entity.ServerGroupFilter {
	filter := entity.ServerGroupFilter{
		Pagination:    pagination,
		IsManagedByMe: isManagedByMe,
	}

	if organizationID != nil {
		orgID, err := uuid.FromString(organizationID.String())
		if err == nil {
			filter.OrganizationID = orgID
		}
	}

	return filter
}

func ServerGroupFilterToDB(filter entity.ServerGroupFilter) dbmodels.ServerGroupFilter {
	return dbmodels.ServerGroupFilter{
		Pagination:     PaginationToDB(filter.Pagination),
		OrganizationID: filter.OrganizationID,
		IsManagedByMe:  filter.IsManagedByMe,
	}
}
