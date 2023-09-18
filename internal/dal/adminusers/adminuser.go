package adminusers

import (
	"context"
	"errors"
	"fmt"
	"washbonus/internal/dal"
	"washbonus/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

const (
	adminResource       = dbmodels.AdminUsersResource
	applicationResource = dbmodels.AdminApplicationsResource
)

func (r *adminUserRepo) GetById(ctx context.Context, userID string) (dbmodels.AdminUser, error) {
	op := "failed to get admin by ID: %w"

	var dbUser dbmodels.AdminUser
	err := r.db.NewSession(nil).
		Select("*").
		From("admin_users").
		Where("NOT deleted AND id = ?", userID).
		LoadOneContext(ctx, &dbUser)

	if err == nil {
		return dbUser, nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.AdminUser{}, dbmodels.ErrNotFound
	}

	return dbmodels.AdminUser{}, fmt.Errorf(op, err)
}

func (r *adminUserRepo) Get(ctx context.Context, pagination dbmodels.Pagination) ([]dbmodels.AdminUser, error) {
	op := "failed to get admins: %w"

	var dbUsers []dbmodels.AdminUser
	_, err := r.db.NewSession(nil).
		Select("*").
		From("admin_users").
		Where("NOT deleted").
		Limit(uint64(pagination.Limit)).
		Offset(uint64(pagination.Offset)).
		LoadContext(ctx, &dbUsers)

	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return dbUsers, nil
}

func (r *adminUserRepo) Create(ctx context.Context, ent dbmodels.AdminUserCreation) (dbmodels.AdminUser, error) {
	op := "failed to create admin: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.AdminUser{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	var dbUser dbmodels.AdminUser
	err = tx.InsertInto("admin_users").
		Columns("id", "email", "name", "organization_id").
		Record(ent).
		Returning("id", "name", "email", "role", "organization_id", "deleted").
		LoadContext(ctx, &dbUser)

	if err != nil {
		return dbmodels.AdminUser{}, fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, adminResource, ent.ID, "create", ent)
	if err != nil {
		return dbmodels.AdminUser{}, fmt.Errorf(op, err)
	}

	return dbUser, tx.Commit()
}

func (r *adminUserRepo) UpdateRole(ctx context.Context, updateUser dbmodels.AdminUserRoleUpdate) error {
	op := "failed to update admin role: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.Update("admin_users").
		Where("NOT deleted AND id = ?", updateUser.ID).
		Set("role", updateUser.Role).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, adminResource, updateUser.ID, "update role", updateUser)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func (r *adminUserRepo) Update(ctx context.Context, userModel dbmodels.AdminUserUpdate) error {
	op := "failed to update admin: %w"

	_, err := r.db.NewSession(nil).
		Update("admin_users").
		Where("NOT deleted AND id = ?", userModel.ID).
		Set("email", userModel.Email).
		Set("name", userModel.Name).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	return err
}

func (r *adminUserRepo) Delete(ctx context.Context, id string) error {
	op := "failed to delete admin user: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.DeleteFrom("admin_users").Where("id = ?", id).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, adminResource, id, "delete", nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func (r *adminUserRepo) GetApplications(ctx context.Context, filter dbmodels.AdminApplicationFilter) ([]dbmodels.AdminApplication, error) {
	op := "failed to get admin applications: %w"

	var dbApps []dbmodels.AdminApplication
	query := r.db.NewSession(nil).
		Select("*").
		From("admin_applications")

	if filter.Status != nil {
		query = query.Where("status = ?", filter.Status)
	}

	_, err := query.
		Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		LoadContext(ctx, &dbApps)

	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return dbApps, nil
}

func (r *adminUserRepo) CreateApplication(ctx context.Context, ent dbmodels.AdminApplicationCreation) (dbmodels.AdminApplication, error) {
	op := "failed to create admin application: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.AdminApplication{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	var dbApp dbmodels.AdminApplication
	err = tx.Select("*").
		From("admin_applications").
		Where("admin_user_id = ?", ent.AdminUserID).
		LoadOneContext(ctx, &dbApp)

	if err == nil {
		return dbApp, nil
	}

	if err != nil && !errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.AdminApplication{}, fmt.Errorf(op, err)
	}

	err = tx.InsertInto("admin_applications").
		Columns("admin_user_id", "name", "email").
		Record(ent).
		Returning("id", "admin_user_id", "name", "email", "status").
		LoadContext(ctx, &dbApp)

	if err != nil {
		return dbmodels.AdminApplication{}, fmt.Errorf(op, err)
	}

	return dbApp, tx.Commit()
}

func (r *adminUserRepo) ReviewApplication(ctx context.Context, id uuid.UUID, ent dbmodels.AdminApplicationReview) error {
	op := "failed to review application: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	var dbApp dbmodels.AdminApplication
	err = tx.Select("*").
		From("admin_applications").
		Where("id = ?", id).
		LoadOneContext(ctx, &dbApp)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return dbmodels.ErrNotFound
		}

		return fmt.Errorf(op, err)
	}

	if ent.Status == dbmodels.Accepted {
		err = createAdminUserFromApplication(ctx, tx, dbApp)
		if err != nil {
			return fmt.Errorf(op, err)
		}

		if ent.OrganizationID != nil {
			_, err = tx.Update("admin_users").
				Set("organization_id", ent.OrganizationID).
				Where("id = ?", dbApp.AdminUserID).
				ExecContext(ctx)
		}
	}

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			err = fmt.Errorf("Organization not found: %w", dbmodels.ErrNotFound)
		}

		return fmt.Errorf(op, err)
	}

	_, err = tx.DeleteFrom("admin_applications").
		Where("id = ?", id).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, applicationResource, id.String(), "review", ent)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func createAdminUserFromApplication(ctx context.Context, tx *dbr.Tx, app dbmodels.AdminApplication) error {
	var user dbmodels.AdminUser
	err := tx.Select("*").
		From("admin_users").
		Where("id = ?", app.AdminUserID).
		LoadOneContext(ctx, &user)

	if err == nil {
		return nil
	}

	if err != nil && !errors.Is(err, dbr.ErrNotFound) {
		return err
	}

	newUser := dbmodels.AdminUser{
		ID:    app.AdminUserID,
		Name:  &app.Name,
		Email: &app.Email,
	}

	_, err = tx.InsertInto("admin_users").
		Columns("id", "name", "email").
		Record(newUser).
		ExecContext(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *adminUserRepo) GetApplicationByUser(ctx context.Context, id string) (dbmodels.AdminApplication, error) {
	op := "failed to get application by user: %w"

	var app dbmodels.AdminApplication
	err := r.db.NewSession(nil).
		Select("*").
		From("admin_applications").
		Where("admin_user_id = ?", id).
		LoadOneContext(ctx, &app)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = dbmodels.ErrNotFound
		}

		return dbmodels.AdminApplication{}, fmt.Errorf(op, err)
	}

	return app, nil
}
