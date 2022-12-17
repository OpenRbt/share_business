package dal

import (
	"context"
	"errors"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"

	"github.com/gocraft/dbr/v2"
)

func (s *Storage) GetProfileOrCreateIfNotExists(ctx context.Context, identity string) (entity.User, error) {
	var dbUser dbmodels.User

	err := s.db.NewSession(nil).
		Select("*").
		From("users").
		Where("identity = ?", identity).
		LoadOneContext(ctx, &dbUser)

	switch {

	case err == nil:
		return conversions.UserFromDb(dbUser), err
	case errors.Is(err, dbr.ErrNotFound):
		return s.CreateProfile(ctx, identity)
	default:
		return entity.User{}, err
	}
}

func (s *Storage) CreateProfile(ctx context.Context, identity string) (entity.User, error) {
	var dbUser dbmodels.User

	err := s.db.NewSession(nil).
		InsertInto("users").
		Columns("identity").
		Values(identity).
		Returning("id", "identity", "balance", "active").
		LoadContext(ctx, &dbUser)

	if err != nil {
		return entity.User{}, err
	}

	return conversions.UserFromDb(dbUser), nil
}
