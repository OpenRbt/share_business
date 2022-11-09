package dto

import (
	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/app/entity/vo"
	"wash-bonus/internal/dal/dbmodel"
	"wash-bonus/openapi/models"
)

func WashServersFromDB(ss []dbmodel.WashServer) []entity.WashServer {
	res := make([]entity.WashServer, len(ss))

	for i, s := range ss {
		res[i] = *WashServerFromDB(s)
	}

	return res
}

func WashServerFromDB(s dbmodel.WashServer) *entity.WashServer {
	return &entity.WashServer{
		CreatedAt:   s.CreatedAt,
		ModifiedAt:  s.ModifiedAt,
		ID:          s.ID,
		Owner:       entity.User{ID: s.OwnerID},
		ServiceKey:  s.ServiceKey.String,
		Name:        s.Name,
		Description: s.Description,
	}
}

func WashServersToRest(ss []entity.WashServer) []*models.WashServer {
	apiWashServers := make([]*models.WashServer, len(ss))

	for i, s := range ss {
		apiWashServer := WashServerToRest(&s)
		apiWashServers[i] = apiWashServer
	}

	return apiWashServers
}

func WashServerToRest(s *entity.WashServer) *models.WashServer {
	return &models.WashServer{
		ID:          s.ID.String(),
		ServiceKey:  s.ServiceKey,
		CreatedAt:   (*strfmt.DateTime)(s.CreatedAt),
		ModifiedAt:  (*strfmt.DateTime)(s.ModifiedAt),
		Name:        s.Name,
		Description: s.Description,
		OwnerID:     s.Owner.ID.String(),
	}
}

func WashServerFromRestAdd(s *models.WashServerAdd) (*entity.WashServer, error) {
	id, err := uuid.FromString(s.OwnerID)
	if err != nil {
		return nil, err
	}

	return &entity.WashServer{
		Name:        s.Name,
		Description: s.Description,
		Owner:       entity.User{ID: id},
	}, nil
}

func WashServerFromRestUpdate(s *models.WashServerUpdate) (*vo.WashServerUpdate, error) {
	id, err := uuid.FromString(s.OwnerID)
	if err != nil {
		return nil, err
	}

	return &vo.WashServerUpdate{
		ServiceKey:  s.ServiceKey,
		Name:        s.Name,
		Description: s.Description,
		OwnerID:     id,
	}, nil
}
