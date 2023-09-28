package conversions

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
	"washbonus/openapi/admin/models"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func WashServerFromDB(dbWashServer dbmodels.WashServer) entities.WashServer {
	return entities.WashServer{
		ID:             dbWashServer.ID,
		Title:          dbWashServer.Title,
		Description:    dbWashServer.Description,
		CreatedBy:      dbWashServer.CreatedBy,
		ServiceKey:     &dbWashServer.ServiceKey,
		GroupID:        dbWashServer.GroupID,
		OrganizationID: dbWashServer.OrganizationID,
	}
}

func WashServerToRest(e entities.WashServer) *models.WashServer {
	return &models.WashServer{
		ID:             e.ID.String(),
		Description:    e.Description,
		Name:           e.Title,
		GroupID:        strfmt.UUID(e.GroupID.String()),
		OrganizationID: strfmt.UUID(e.OrganizationID.String()),
		CreatedBy:      e.CreatedBy,
	}
}

func WashServerToAdminRest(e entities.WashServer) *models.WashServer {
	m := &models.WashServer{
		ID:             e.ID.String(),
		Name:           e.Title,
		Description:    e.Description,
		CreatedBy:      e.CreatedBy,
		GroupID:        strfmt.UUID(e.GroupID.String()),
		OrganizationID: strfmt.UUID(e.OrganizationID.String()),
	}

	if e.ServiceKey != nil {
		m.ServiceKey = *e.ServiceKey
	}

	return m
}

func WashServerListFromDB(washServerList []dbmodels.WashServer) []entities.WashServer {
	res := make([]entities.WashServer, len(washServerList))

	for i, value := range washServerList {
		res[i] = WashServerFromDB(value)
	}

	return res
}

func WashServerListToRest(washServerEntity []entities.WashServer) []*models.WashServer {
	res := make([]*models.WashServer, len(washServerEntity))

	for i, value := range washServerEntity {
		rest := WashServerToRest(value)
		res[i] = rest
	}

	return res
}

func WashServerUpdateFromRest(model models.WashServerUpdate) entities.WashServerUpdate {
	return entities.WashServerUpdate{
		Title:       model.Name,
		Description: model.Description,
	}
}

func WashServerUpdateToDb(e entities.WashServerUpdate) dbmodels.WashServerUpdate {
	return dbmodels.WashServerUpdate{
		Name:        e.Title,
		Description: e.Description,
	}
}

func WashServerCreationFromRest(rest models.WashServerCreation) entities.WashServerCreation {
	ent := entities.WashServerCreation{
		Title:       *rest.Name,
		Description: *rest.Description,
	}

	if rest.GroupID != nil {
		groupID, err := uuid.FromString(string(*rest.GroupID))
		if err == nil {
			ent.GroupID = &groupID
		}
	}

	return ent
}

func WashServerCreationToDb(e entities.WashServerCreation) dbmodels.WashServerCreation {
	return dbmodels.WashServerCreation{
		Title:       e.Title,
		Description: e.Description,
		GroupID:     e.GroupID,
	}
}

func WashServerFilterFromRest(pagination entities.Pagination, organizationID *strfmt.UUID, groupID *strfmt.UUID) entities.WashServerFilter {
	filter := entities.WashServerFilter{
		Pagination: pagination,
	}

	if organizationID != nil {
		orgID, err := uuid.FromString(organizationID.String())
		if err == nil {
			filter.OrganizationID = &orgID
		}
	}

	if groupID != nil {
		grID, err := uuid.FromString(groupID.String())
		if err == nil {
			filter.GroupID = &grID
		}
	}

	return filter
}

func WashServerFilterToDB(filter entities.WashServerFilter) dbmodels.WashServerFilter {
	return dbmodels.WashServerFilter{
		Pagination:     PaginationToDB(filter.Pagination),
		OrganizationID: filter.OrganizationID,
		GroupID:        filter.GroupID,
	}
}
