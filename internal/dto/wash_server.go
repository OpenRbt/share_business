package dto

import (
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/dal/dbmodel"
)

func WashServersFromDB(ss []dbmodel.WashServer) []entity.WashServer {
	res := make([]entity.WashServer, len(ss))

	for i, s := range ss {
		res[i] = WashServerFromDB(s)
	}

	return res
}

func WashServerFromDB(s dbmodel.WashServer) entity.WashServer {
	return entity.WashServer{
		CreatedAt:   s.CreatedAt,
		ModifiedAt:  s.ModifiedAt,
		ID:          s.ID,
		OwnerID:     s.OwnerID,
		ServiceKey:  s.ServiceKey,
		Name:        s.Name,
		Description: s.Description,
	}
}
