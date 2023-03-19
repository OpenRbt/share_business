package conversions

import (
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/entity/vo"
	models2 "wash_bonus/internal/infrastructure/rabbit/models"
	"wash_bonus/openapi/models"

	uuid "github.com/satori/go.uuid"
)

func WashServerFromDB(dbWashServer dbmodels.WashServer) entity.WashServer {
	return entity.WashServer{
		Id:          dbWashServer.ID.UUID,
		Title:       dbWashServer.Title,
		Description: dbWashServer.Description,
	}
}

func WashServerCreationFromRabbit(m models2.ServerRegistered) (e entity.WashServer, err error) {
	id, err := uuid.FromString(m.ID)
	if err != nil {
		return
	}

	e.Id = id
	e.Title = m.Title
	e.Description = m.Description

	return
}

func WashServerToRest(e entity.WashServer) *models.WashServer {
	return &models.WashServer{
		Description: e.Description,
		Name:        e.Title,
	}
}

func WashServerUpdateFromRabbit(m models2.ServerUpdate) (v vo.WashServerUpdate, err error) {
	id, err := uuid.FromString(m.ID)
	if err != nil {
		return
	}

	v.ID = id
	v.Title = m.Title
	v.Description = m.Description

	v.Deleted = m.Deleted

	return
}
