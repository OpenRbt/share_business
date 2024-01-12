package conversions

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
	rabbitEntities "washbonus/internal/infrastructure/rabbit/entities"
	"washbonus/openapi/admin/models"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func ServerGroupFromDB(group dbmodels.ServerGroup) entities.ServerGroup {
	return entities.ServerGroup{
		ID:                            group.ID,
		OrganizationID:                group.OrganizationID,
		Name:                          group.Name,
		Description:                   group.Description,
		ReportsProcessingDelayMinutes: group.ReportsProcessingDelayMinutes,
		BonusPercentage:               group.BonusPercentage,
		UTCOffset:                     group.UTCOffset,
		IsDefault:                     group.IsDefault,
		Deleted:                       group.Deleted,
		Version:                       group.Version,
	}
}

func ServerGroupToRest(group entities.ServerGroup) *models.ServerGroup {
	return &models.ServerGroup{
		ID:                            strfmt.UUID(group.ID.String()),
		OrganizationID:                strfmt.UUID(group.OrganizationID.String()),
		Name:                          group.Name,
		Description:                   group.Description,
		UtcOffset:                     &group.UTCOffset,
		IsDefault:                     group.IsDefault,
		ReportsProcessingDelayMinutes: &group.ReportsProcessingDelayMinutes,
		BonusPercentage:               &group.BonusPercentage,
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
		Name:                          groupUpdate.Name,
		Description:                   groupUpdate.Description,
		UTCOffset:                     groupUpdate.UTCOffset,
		ReportsProcessingDelayMinutes: groupUpdate.ReportsProcessingDelayMinutes,
		BonusPercentage:               groupUpdate.BonusPercentage,
	}
}

func ServerGroupUpdateFromRest(groupUpdate models.ServerGroupUpdate) entities.ServerGroupUpdate {
	return entities.ServerGroupUpdate{
		Name:                          groupUpdate.Name,
		Description:                   groupUpdate.Description,
		UTCOffset:                     groupUpdate.UtcOffset,
		ReportsProcessingDelayMinutes: groupUpdate.ReportsProcessingDelayMinutes,
		BonusPercentage:               groupUpdate.BonusPercentage,
	}
}

func ServerGroupCreationToDb(groupCreation entities.ServerGroupCreation) dbmodels.ServerGroupCreation {
	return dbmodels.ServerGroupCreation{
		OrganizationID:                groupCreation.OrganizationID,
		Name:                          groupCreation.Name,
		Description:                   groupCreation.Description,
		UTCOffset:                     groupCreation.UTCOffset,
		ReportsProcessingDelayMinutes: groupCreation.ReportsProcessingDelayMinutes,
		BonusPercentage:               groupCreation.BonusPercentage,
	}
}

func ServerGroupCreationFromRest(model models.ServerGroupCreation) entities.ServerGroupCreation {
	return entities.ServerGroupCreation{
		OrganizationID:                uuid.FromStringOrNil((*model.OrganizationID).String()),
		Name:                          *model.Name,
		Description:                   *model.Description,
		UTCOffset:                     model.UtcOffset,
		ReportsProcessingDelayMinutes: model.ReportsProcessingDelayMinutes,
		BonusPercentage:               model.BonusPercentage,
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

func ServerGroupToRabbit(group entities.ServerGroup) rabbitEntities.ServerGroup {
	return rabbitEntities.ServerGroup{
		ID:             group.ID.String(),
		OrganizationID: group.OrganizationID.String(),
		Name:           group.Name,
		Description:    group.Description,
		UTCOffset:      group.UTCOffset,
		IsDefault:      group.IsDefault,
		Deleted:        group.Deleted,
		Version:        group.Version,
	}
}

func ServerGroupsToRabbit(groups []entities.ServerGroup) []rabbitEntities.ServerGroup {
	res := make([]rabbitEntities.ServerGroup, len(groups))

	for i, value := range groups {
		res[i] = ServerGroupToRabbit(value)
	}

	return res
}
