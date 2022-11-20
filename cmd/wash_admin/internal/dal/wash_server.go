package dal

import (
	"context"
	"errors"
	"wash_admin/internal/conversions"
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

func (s *Storage) GetWashServer(ctx context.Context, ownerId uuid.UUID, id uuid.UUID) (entity.WashServer, error) {
	var dbWashServer dbmodels.WashServer

	err := s.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("id = ? AND owner = ?", uuid.NullUUID{UUID: id, Valid: true}, uuid.NullUUID{UUID: ownerId, Valid: true}).
		LoadOneContext(ctx, &dbWashServer)

	switch {
	case err == nil:
		return conversions.WashServerFromDB(dbWashServer), err
	case errors.Is(err, dbr.ErrNotFound):
		return entity.WashServer{}, entity.ErrNotFound
	default:
		return entity.WashServer{}, err
	}
}