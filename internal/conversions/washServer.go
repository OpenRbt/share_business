package conversions

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/openapi/models"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func WashServerFromDB(dbWashServer dbmodels.WashServer) entity.WashServer {
	return entity.WashServer{
		ID:             dbWashServer.ID,
		Title:          dbWashServer.Title,
		Description:    dbWashServer.Description,
		CreatedBy:      dbWashServer.CreatedBy,
		ServiceKey:     dbWashServer.ServiceKey,
		GroupID:        dbWashServer.GroupID,
		OrganizationID: dbWashServer.OrganizationID,
	}
}

func WashServerToRest(e entity.WashServer) *models.WashServer {
	return &models.WashServer{
		ID:             e.ID.String(),
		Description:    e.Description,
		Name:           e.Title,
		GroupID:        strfmt.UUID(e.GroupID.String()),
		OrganizationID: strfmt.UUID(e.OrganizationID.String()),
	}
}

func WashServerToAdminRest(e entity.WashServer) *models.WashServer {
	return &models.WashServer{
		ID:             e.ID.String(),
		Name:           e.Title,
		Description:    e.Description,
		CreatedBy:      e.CreatedBy,
		ServiceKey:     e.ServiceKey,
		GroupID:        strfmt.UUID(e.GroupID.String()),
		OrganizationID: strfmt.UUID(e.OrganizationID.String()),
	}
}

func WashServerListFromDB(washServerList []dbmodels.WashServer) []entity.WashServer {
	res := make([]entity.WashServer, len(washServerList))

	for i, value := range washServerList {
		res[i] = WashServerFromDB(value)
	}

	return res
}

func WashServerListToRest(washServerEntity []entity.WashServer) []*models.WashServer {
	res := make([]*models.WashServer, len(washServerEntity))

	for i, value := range washServerEntity {
		rest := WashServerToRest(value)
		res[i] = rest
	}

	return res
}

func WashServerUpdateFromRest(model models.WashServerUpdate) entity.WashServerUpdate {
	return entity.WashServerUpdate{
		Title:       &model.Name,
		Description: &model.Description,
	}
}

func WashServerUpdateToDb(entity entity.WashServerUpdate) dbmodels.WashServerUpdate {
	return dbmodels.WashServerUpdate{
		Name:        entity.Title,
		Description: entity.Description,
	}
}

func WashServerCreationFromRest(rest models.WashServerCreation) entity.WashServerCreation {
	ent := entity.WashServerCreation{
		Title:       *rest.Name,
		Description: *rest.Description,
	}

	if rest.GroupID != nil {
		groupID, err := uuid.FromString(string(*rest.GroupID))
		if err == nil {
			ent.GroupID = uuid.NullUUID{UUID: groupID, Valid: true}
		}
	}

	return ent
}

func WashServerCreationToDb(entity entity.WashServerCreation) dbmodels.WashServerCreation {
	return dbmodels.WashServerCreation{
		Title:       entity.Title,
		Description: entity.Description,
		GroupID:     entity.GroupID,
	}
}

func WashServerFilterFromRest(pagination entity.Pagination, isManagedByMe bool, organizationID *strfmt.UUID, groupID *strfmt.UUID) entity.WashServerFilter {
	filter := entity.WashServerFilter{
		Pagination:    pagination,
		IsManagedByMe: isManagedByMe,
	}

	if organizationID != nil {
		orgID, err := uuid.FromString(organizationID.String())
		if err == nil {
			filter.OrganizationID = orgID
		}
	}

	if groupID != nil {
		groupID, err := uuid.FromString(groupID.String())
		if err == nil {
			filter.GroupID = groupID
		}
	}

	return filter
}

func WashServerFilterToDB(filter entity.WashServerFilter) dbmodels.WashServerFilter {
	return dbmodels.WashServerFilter{
		Pagination:     PaginationToDB(filter.Pagination),
		OrganizationID: filter.OrganizationID,
		GroupID:        filter.GroupID,
		IsManagedByMe:  filter.IsManagedByMe,
	}
}
