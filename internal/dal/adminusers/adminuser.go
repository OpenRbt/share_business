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
	const op = "failed to get admin by ID: %w"

	var dbUser dbmodels.AdminUser
	err := r.db.NewSession(nil).
		Select("*").
		From("admin_users").
		Where("id = ?", userID).
		LoadOneContext(ctx, &dbUser)

	if err == nil {
		return dbUser, nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.AdminUser{}, dbmodels.ErrNotFound
	}

	return dbmodels.AdminUser{}, fmt.Errorf(op, err)
}

func (r *adminUserRepo) Get(ctx context.Context, filter dbmodels.AdminUserFilter) ([]dbmodels.AdminUser, error) {
	const op = "failed to get admins: %w"

	var dbUsers []dbmodels.AdminUser
	query := r.db.NewSession(nil).
		Select("*").
		From("admin_users")

	if filter.Role != nil {
		query = query.Where("role = ?", filter.Role)
	}

	if filter.IsBlocked != nil {
		if *filter.IsBlocked {
			query = query.Where("role = ?", dbmodels.NoAccessRole)
		} else {
			query = query.Where("role != ?", dbmodels.NoAccessRole)
		}
	}

	_, err := query.Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		LoadContext(ctx, &dbUsers)

	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return dbUsers, nil
}

func (r *adminUserRepo) Create(ctx context.Context, ent dbmodels.AdminUserCreation) (dbmodels.AdminUser, error) {
	const op = "failed to create admin: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.AdminUser{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	var dbUser dbmodels.AdminUser
	err = tx.InsertInto("admin_users").
		Columns("id", "email", "name", "organization_id").
		Record(ent).
		Returning("id", "name", "email", "role", "organization_id").
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
	const op = "failed to update admin role: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.Update("admin_users").
		Where("id = ?", updateUser.ID).
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
	const op = "failed to update admin: %w"

	_, err := r.db.NewSession(nil).
		Update("admin_users").
		Where("id = ?", userModel.ID).
		Set("email", userModel.Email).
		Set("name", userModel.Name).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	return err
}

func (r *adminUserRepo) Block(ctx context.Context, id string) error {
	const op = "failed to block admin user: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.Update("admin_users").
		Where("id = ?", id).
		Set("role", dbmodels.NoAccessRole).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, adminResource, id, "block", nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func (r *adminUserRepo) GetApplications(ctx context.Context, filter dbmodels.AdminApplicationFilter) ([]dbmodels.AdminApplication, error) {
	const op = "failed to get admin applications: %w"

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
	const op = "failed to create admin application: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.AdminApplication{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	var dbApp dbmodels.AdminApplication
	isExists, err := r.loadAdminApplication(ctx, tx, ent, &dbApp)
	if err != nil {
		return dbmodels.AdminApplication{}, fmt.Errorf(op, err)
	}

	if isExists {
		return dbApp, nil
	}

	if err := r.validateAdminUser(ctx, tx, ent); err != nil {
		return dbmodels.AdminApplication{}, fmt.Errorf(op, err)
	}

	if err := r.insertAdminApplication(ctx, tx, ent, &dbApp); err != nil {
		return dbmodels.AdminApplication{}, fmt.Errorf(op, err)
	}

	return dbApp, tx.Commit()
}

func (r *adminUserRepo) loadAdminApplication(ctx context.Context, tx *dbr.Tx, ent dbmodels.AdminApplicationCreation, dbApp *dbmodels.AdminApplication) (bool, error) {
	err := tx.Select("*").
		From("admin_applications").
		Where("admin_user_id = ?", ent.AdminUserID).
		LoadOneContext(ctx, dbApp)

	if err == nil {
		return true, nil
	}
	if !errors.Is(err, dbr.ErrNotFound) {
		return false, err
	}

	return false, nil
}

func (r *adminUserRepo) validateAdminUser(ctx context.Context, tx *dbr.Tx, ent dbmodels.AdminApplicationCreation) error {
	var admin dbmodels.AdminUser
	err := tx.Select("*").
		From("admin_users").
		Where("id = ?", ent.AdminUserID).
		LoadOneContext(ctx, &admin)

	if err == nil && admin.Role != dbmodels.NoAccessRole {
		return fmt.Errorf("admin user %w", dbmodels.ErrAlreadyExists)
	}
	if err != nil && !errors.Is(err, dbr.ErrNotFound) {
		return err
	}

	return nil
}

func (r *adminUserRepo) insertAdminApplication(ctx context.Context, tx *dbr.Tx, ent dbmodels.AdminApplicationCreation, dbApp *dbmodels.AdminApplication) error {
	return tx.InsertInto("admin_applications").
		Columns("admin_user_id", "name", "email").
		Record(ent).
		Returning("id", "admin_user_id", "name", "email", "status").
		LoadContext(ctx, dbApp)
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
		err = r.createAdminUserFromApplication(ctx, tx, dbApp, ent.Role)
		if err != nil {
			return fmt.Errorf(op, err)
		}

		if ent.Role != nil && *ent.Role == dbmodels.AdminRole && ent.OrganizationID == nil {
			return fmt.Errorf("OrganizationID required for admin role: %w", dbmodels.ErrBadRequest)
		}

		if ent.OrganizationID != nil {
			_, err = tx.Update("admin_users").
				Set("organization_id", ent.OrganizationID).
				Where("id = ?", dbApp.AdminUserID).
				ExecContext(ctx)

			if err != nil {
				if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
					err = fmt.Errorf("Organization not found: %w", dbmodels.ErrNotFound)
				}

				return fmt.Errorf(op, err)
			}
		}
	}

	_, err = tx.DeleteFrom("admin_applications").
		Where("id = ?", id).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, applicationResource, id.String(), "review init", dbApp)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, applicationResource, id.String(), "review end", ent)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func (r *adminUserRepo) createAdminUserFromApplication(ctx context.Context, tx *dbr.Tx, app dbmodels.AdminApplication, role *dbmodels.Role) error {
	if role == nil {
		r := dbmodels.AdminRole
		role = &r
	}

	_, err := tx.ExecContext(ctx, `
		INSERT INTO admin_users (id, name, email, role)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id)
		DO UPDATE SET name = EXCLUDED.name, email = EXCLUDED.email, role = EXCLUDED.role
	`, app.AdminUserID, app.Name, app.Email, role)

	return err
}

func (r *adminUserRepo) GetApplicationByID(ctx context.Context, id uuid.UUID) (dbmodels.AdminApplication, error) {
	const op = "failed to get application by ID: %w"

	var app dbmodels.AdminApplication
	err := r.db.NewSession(nil).
		Select("*").
		From("admin_applications").
		Where("id = ?", id).
		LoadOneContext(ctx, &app)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = dbmodels.ErrNotFound
		}

		return dbmodels.AdminApplication{}, fmt.Errorf(op, err)
	}

	return app, nil
}
