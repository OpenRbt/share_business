package dal

import (
	"context"
	"errors"
	"github.com/gocraft/dbr/v2"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
)

func (s *Storage) GetUser(ctx context.Context, identity string) (entity.User, error) {
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
		return entity.User{}, entity.ErrNotFound
	default:
		return entity.User{}, err
	}
}
