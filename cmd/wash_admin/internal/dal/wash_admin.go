package dal

import (
	"context"
	"errors"
	"wash_admin/internal/conversions"
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity"

	"github.com/gocraft/dbr/v2"
)

func (s *Storage) GetWashAdmin(ctx context.Context, identity string) (entity.WashAdmin, error) {
	var dbWashAdmin dbmodels.WashAdmin

	err := s.db.NewSession(nil).
		Select("*").
		From("users").
		Where("identity_uid = ?", identity).
		LoadOneContext(ctx, &dbWashAdmin)

	switch {
	case err == nil:
		return conversions.WashAdminFromDB(dbWashAdmin), err
	case errors.Is(err, dbr.ErrNotFound):
		return entity.WashAdmin{}, entity.ErrNotFound
	default:
		return entity.WashAdmin{}, err
	}
}

func (s *Storage) CreateWashAdmin(ctx context.Context, identity string) (entity.WashAdmin, error) {
	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return entity.WashAdmin{}, err
	}

	var dbWashAdmin dbmodels.WashAdmin
	err = tx.
		InsertInto("users").
		Columns("identity_uid", "role").
		Values(identity, "user").
		Returning("id", "identity_uid", "role").
		LoadContext(ctx, &dbWashAdmin)

	if err != nil {
		return entity.WashAdmin{}, err
	}

	return conversions.WashAdminFromDB(dbWashAdmin), tx.Commit()
}

func (s *Storage) GetOrCreateAdminIfNotExists(ctx context.Context, identity string) (entity.WashAdmin, error) {
	dbWashAdmin, err := s.GetWashAdmin(ctx, identity)

	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return s.CreateWashAdmin(ctx, identity)
		}

		return entity.WashAdmin{}, err
	}

	return dbWashAdmin, err
}
