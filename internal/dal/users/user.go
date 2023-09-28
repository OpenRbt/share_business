package users

import (
	"context"
	"errors"
	"fmt"
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/dal/organizations"

	"github.com/gocraft/dbr/v2"
)

func (r *userRepo) GetById(ctx context.Context, userID string) (dbmodels.User, error) {
	op := "failed to get user by ID: %w"

	var dbUser dbmodels.User
	err := r.db.NewSession(nil).
		Select("*").
		From("users").
		Where("NOT deleted AND id = ?", userID).
		LoadOneContext(ctx, &dbUser)

	if err == nil {
		return dbUser, nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.User{}, dbmodels.ErrNotFound
	}

	return dbmodels.User{}, fmt.Errorf(op, err)
}

func (r *userRepo) Get(ctx context.Context, pagination dbmodels.Pagination) ([]dbmodels.User, error) {
	op := "failed to get users: %w"

	var dbUsers []dbmodels.User
	_, err := r.db.NewSession(nil).
		Select("*").
		From("users").
		Where("NOT deleted").
		Limit(uint64(pagination.Limit)).
		Offset(uint64(pagination.Offset)).
		LoadContext(ctx, &dbUsers)

	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return dbUsers, nil
}

func (r *userRepo) Create(ctx context.Context, ent dbmodels.UserCreation) (dbmodels.User, error) {
	op := "failed to create user: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.User{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	var dbUser dbmodels.User

	err = tx.InsertInto("users").
		Columns("id", "email", "name").
		Values(ent.ID, ent.Email, ent.Name).
		Returning("id", "email", "name", "deleted").
		LoadContext(ctx, &dbUser)

	if err != nil {
		return dbmodels.User{}, fmt.Errorf(op, err)
	}

	var org dbmodels.Organization
	err = tx.Select(organizations.OrgColumns...).From("organizations").Where("is_default").LoadOneContext(ctx, &org)
	if err != nil {
		return dbmodels.User{}, fmt.Errorf(op, err)
	}

	_, err = tx.InsertInto("wallets").
		Columns("user_id", "organization_id", "is_default").
		Values(ent.ID, org.ID, true).
		ExecContext(ctx)

	if err != nil {
		return dbmodels.User{}, fmt.Errorf(op, err)
	}

	return dbUser, tx.Commit()
}

func (r *userRepo) UpdateUser(ctx context.Context, userModel dbmodels.UserUpdate) error {
	op := "failed to update user: %w"

	_, err := r.db.NewSession(nil).
		Update("users").
		Where("id = ?", userModel.ID).
		Set("email", userModel.Email).
		Set("name", userModel.Name).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	return err
}
