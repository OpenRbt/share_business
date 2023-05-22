package dal

import (
	"context"
	"errors"
	"wash_admin/internal/app"
	"wash_admin/internal/conversions"
	"wash_admin/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
)

func (s *Storage) GetWashUser(ctx context.Context, identity string) (app.WashUser, error) {
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
		return app.WashUser{}, app.ErrNotFound
	default:
		return app.WashUser{}, err
	}
}

func (s *Storage) CreateWashUser(ctx context.Context, identity string) (app.WashUser, error) {
	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return app.WashUser{}, err
	}

	var dbWashUser dbmodels.WashUser
	err = tx.
		InsertInto("users").
		Columns("identity_uid", "role").
		Values(identity, "user").
		Returning("id", "identity_uid", "role").
		LoadContext(ctx, &dbWashUser)

	if err != nil {
		return app.WashUser{}, err
	}

	return conversions.WashUserFromDB(dbWashUser), tx.Commit()
}

func (s *Storage) GetOrCreateUserIfNotExists(ctx context.Context, identity string) (app.WashUser, error) {
	dbWashUser, err := s.GetWashUser(ctx, identity)

	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			return s.CreateWashUser(ctx, identity)
		}

		return app.WashUser{}, err
	}

	return dbWashUser, err
}

func (s *Storage) UpdateUserRole(ctx context.Context, updateUser app.UpdateUser) error {
	dbUpdateUser := conversions.WashUserToDB(updateUser)

	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	updateStatement := tx.
		Update("users").
		Where("id = ?", dbUpdateUser.ID).Set("role", dbUpdateUser.Role)

	_, err = updateStatement.ExecContext(ctx)

	if err != nil {
		return err
	}

	return tx.Commit()
}
