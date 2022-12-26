package wash_server

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
)

func (s *repo) GetWashServer(ctx context.Context, id uuid.UUID) (entity.WashServer, error) {
	var res dbmodels.WashServer
	err := s.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("id = ?", uuid.NullUUID{
			UUID:  id,
			Valid: true,
		}).
		LoadOneContext(ctx, &res)
	if err != nil {
		return entity.WashServer{}, err
	}

	return conversions.WashServerFromDB(res), nil
}

func (s *repo) GetWashServerByKey(ctx context.Context, key string) (entity.WashServer, error) {
	var res dbmodels.WashServer
	err := s.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("wash_key = ?", key).
		LoadOneContext(ctx, &res)
	if err != nil {
		return entity.WashServer{}, err
	}

	return conversions.WashServerFromDB(res), nil
}
