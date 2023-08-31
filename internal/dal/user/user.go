package user

import (
	"context"
	"errors"
	"washBonus/internal/dal"
	"washBonus/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
)

func (r *userRepo) GetById(ctx context.Context, userID string) (dbmodels.User, error) {
	var err error
	defer dal.LogOptionalError(r.l, "user", err)

	var dbUser dbmodels.User

	err = r.db.NewSession(nil).
		Select("users.*, ARRAY_TO_JSON(ARRAY_AGG(man.organization_id)) AS organization_ids").
		From("users").
		LeftJoin(dbr.I("organization_managers").As("man"), "users.id = man.user_id").
		Where("users.id = ?", userID).
		GroupBy("users.id").
		LoadOneContext(ctx, &dbUser)

	if err == nil {
		return dbUser, nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.User{}, dbmodels.ErrNotFound
	}

	return dbmodels.User{}, err
}

func (r *userRepo) Get(ctx context.Context, pagination dbmodels.Pagination) ([]dbmodels.User, error) {
	var err error
	defer dal.LogOptionalError(r.l, "user", err)

	var dbUsers []dbmodels.User
	_, err = r.db.NewSession(nil).
		Select("users.*, ARRAY_TO_JSON(ARRAY_AGG(man.organization_id)) AS organization_ids").
		From("users").
		LeftJoin(dbr.I("organization_managers").As("man"), "users.id = man.user_id").
		GroupBy("users.id").
		Limit(uint64(pagination.Limit)).
		Offset(uint64(pagination.Offset)).
		LoadContext(ctx, &dbUsers)

	return dbUsers, err
}

func (r *userRepo) Create(ctx context.Context, ent dbmodels.UserCreation) (dbmodels.User, error) {
	var err error
	defer dal.LogOptionalError(r.l, "user", err)

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.User{}, err
	}
	defer tx.RollbackUnlessCommitted()

	var dbUser dbmodels.User

	err = tx.InsertInto("users").
		Columns("id", "email", "name", "role").
		Values(ent.ID, ent.Email, ent.Name, "user").
		Returning("id", "role", "deleted").
		LoadContext(ctx, &dbUser)
	if err != nil {
		return dbmodels.User{}, err
	}

	var org dbmodels.Organization
	err = tx.Select("*").From("organizations").Where("is_default").LoadOneContext(ctx, &org)
	if err != nil {
		return dbmodels.User{}, err
	}

	_, err = tx.InsertInto("wallets").
		Columns("user_id", "organization_id", "is_default").
		Values(ent.ID, org.ID, true).
		ExecContext(ctx)

	if err != nil {
		return dbmodels.User{}, err
	}

	return dbUser, tx.Commit()
}

func (r *userRepo) UpdateUserRole(ctx context.Context, updateUser dbmodels.UserUpdateRole) error {
	var err error
	defer dal.LogOptionalError(r.l, "user", err)

	_, err = r.db.NewSession(nil).
		Update("users").
		Where("id = ?", updateUser.ID).
		Set("role", updateUser.Role).
		ExecContext(ctx)

	return err
}

func (r *userRepo) UpdateUser(ctx context.Context, userModel dbmodels.UserUpdate) error {
	var err error
	defer dal.LogOptionalError(r.l, "user", err)

	_, err = r.db.NewSession(nil).
		Update("users").
		Where("id = ?", userModel.ID).
		Set("email", userModel.Email).
		Set("name", userModel.Name).
		ExecContext(ctx)

	return err
}
