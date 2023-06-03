package dal

import (
	"context"
	"errors"
	"wash_admin/internal/app"
	"wash_admin/internal/conversions"
	"wash_admin/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
)

func (s *Storage) GetUser(ctx context.Context, identity string) (app.User, error) {
	var dbWashUser dbmodels.User

	err := s.db.NewSession(nil).
		Select("*").
		From("users").
		Where("id = ?", identity).
		LoadOneContext(ctx, &dbWashUser)

	switch {
	case err == nil:
		return conversions.WashUserFromDB(dbWashUser), err
	case errors.Is(err, dbr.ErrNotFound):
		return app.User{}, app.ErrNotFound
	default:
		return app.User{}, err
	}
}

func (s *Storage) CreateWashUser(ctx context.Context, id string) (app.User, error) {
	var dbWashUser dbmodels.User

	err := s.db.NewSession(nil).
		InsertInto("users").
		Columns("id", "role").
		Values(id, "user").
		Returning("id", "role").
		LoadContext(ctx, &dbWashUser)

	if err != nil {
		return app.User{}, nil
	}

	return conversions.WashUserFromDB(dbWashUser), nil
}

func (s *Storage) GetOrCreateUserIfNotExists(ctx context.Context, identity string) (app.User, error) {
	dbWashUser, err := s.GetUser(ctx, identity)

	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			return s.CreateWashUser(ctx, identity)
		}

		return app.User{}, err
	}

	return dbWashUser, err
}

func (s *Storage) UpdateUserRole(ctx context.Context, updateUser app.UpdateUser) error {
	dbUpdateUser := conversions.WashUserToDB(updateUser)

	_, err := s.db.NewSession(nil).
		Update("users").
		Where("id = ?", dbUpdateUser.ID).
		Set("role", dbUpdateUser.Role).
		ExecContext(ctx)

	return err
}
