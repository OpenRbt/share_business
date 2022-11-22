package dal

import (
	"context"
	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
)

type Storage struct {
	db *dbr.Connection
	l  *zap.SugaredLogger
}

func (s *Storage) GetWashServerByKey(ctx context.Context, key string) (entity.WashServer, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) GetWashServer(ctx context.Context, id uuid.UUID) (entity.WashServer, error) {
	var res dbmodels.WashServer
	err := s.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		LoadOneContext(ctx, &res)
	if err != nil {
		return entity.WashServer{}, err
	}

	return conversions.WashServerFromDB(res), nil
}

func New(db *dbr.Connection, logger *zap.SugaredLogger) *Storage {
	return &Storage{
		db: db,
		l:  logger,
	}
}
