package dal

import (
	"context"
	"errors"
	"wash_admin/internal/conversions"
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity"

	"github.com/gocraft/dbr/v2"
)

func (s *Storage) GetWashUser(ctx context.Context, identity string) (entity.WashUser, error) {
	var dbWashUser dbmodels.WashUser

	err := s.db.NewSession(nil).
		Select("*").
		From("users").
		Where("identity_uid = ?", identity).
		LoadOneContext(ctx, &dbWashUser)

	switch {
	case err == nil:
		return conversions.WashUserFromDB(dbWashUser), err
	case errors.Is(err, dbr.ErrNotFound):
		return entity.WashUser{}, entity.ErrNotFound
	default:
		return entity.WashUser{}, err
	}
}

func (s *Storage) CreateWashUser(ctx context.Context, identity string) (entity.WashUser, error) {
	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return entity.WashUser{}, err
	}

	var dbWashUser dbmodels.WashUser
	err = tx.
		InsertInto("users").
		Columns("identity_uid", "role").
		Values(identity, "user").
		Returning("id", "identity_uid", "role").
		LoadContext(ctx, &dbWashUser)

	if err != nil {
		return entity.WashUser{}, err
	}

	return conversions.WashUserFromDB(dbWashUser), tx.Commit()
}

func (s *Storage) GetOrCreateUserIfNotExists(ctx context.Context, identity string) (entity.WashUser, error) {
	dbWashUser, err := s.GetWashUser(ctx, identity)

	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return s.CreateWashUser(ctx, identity)
		}

		return entity.WashUser{}, err
	}

	return dbWashUser, err
}
