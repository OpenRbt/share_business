package dto

import (
	"wash-bonus/internal/api/restapi/models"
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/app/entity/vo"
	"wash-bonus/internal/dal/dbmodel"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func WashServersFromDB(ss []dbmodel.WashServer) []entity.WashServer {
	res := make([]entity.WashServer, len(ss))

	for i, s := range ss {
		res[i] = *WashServerFromDB(&s)
	}

	return res
}

func WashServerFromDB(s *dbmodel.WashServer) *entity.WashServer {
	return &entity.WashServer{
		CreatedAt:   s.CreatedAt,
		ModifiedAt:  s.ModifiedAt,
		ID:          s.ID,
		OwnerID:     s.OwnerID,
		ServiceKey:  s.ServiceKey,
		Name:        s.Name,
		Description: s.Description,
	}
}

func WashServersToRest(ss []entity.WashServer) []*models.WashServer {
	apiWashServers := make([]*models.WashServer, len(ss))

	for i, s := range ss {
		apiWashServer := WashServerToRest(s)
		apiWashServers[i] = apiWashServer
	}
	return apiWashServers
}

func WashServerToRest(s entity.WashServer) *models.WashServer {
	return &models.WashServer{
		ID:         s.ID.String(),
		Key:        s.ServiceKey,
		CreatedAt:  (*strfmt.DateTime)(s.CreatedAt),
		ModifiedAt: (*strfmt.DateTime)(s.ModifiedAt),
		Name:       s.Name,
	}
}

func WashServerFromRestAdd(s *models.WashServerAdd) entity.WashServer {
	return entity.WashServer{
		Name:       s.Name,
		ServiceKey: s.Key,
	}
}

func WashServerFromRestUpdate(s models.WashServer) vo.WashServerUpdate {
	// !!! Исправить заполнение структуры
	return vo.WashServerUpdate{
		ServiceKey:  s.Key,
		Name:        s.Name,
		Description: "",
		OwnerID:     uuid.UUID{},
	}
}
