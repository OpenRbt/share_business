package organizations

import (
	"context"
	"errors"
	"fmt"
	"washBonus/internal/dal"
	"washBonus/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

func (r *repo) Get(ctx context.Context, userID string, filter dbmodels.OrganizationFilter) ([]dbmodels.Organization, error) {
	var err error
	defer dal.LogOptionalError(r.l, "organization", err)

	var orgs []dbmodels.Organization

	query := r.db.NewSession(nil).
		Select("org.*").
		From(dbr.I("organizations").As("org")).
		Where("NOT org.deleted")

	if len(filter.OrganizationIDs) > 0 {
		query = query.Where("org.id IN ?", filter.OrganizationIDs)
	}

	if filter.IsManagedByMe {
		query = query.LeftJoin(dbr.I("organization_managers").As("man"), "org.id = man.organization_id")
		query = query.Where("man.user_id = ?", userID)
	}

	_, err = query.
		Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		LoadContext(ctx, &orgs)
	if err != nil {
		return nil, fmt.Errorf("failed to load organizations: %w", err)
	}

	return orgs, nil
}

func (r *repo) GetById(ctx context.Context, id uuid.UUID) (dbmodels.Organization, error) {
	var err error
	defer dal.LogOptionalError(r.l, "organization", err)

	var org dbmodels.Organization

	err = r.db.NewSession(nil).
		Select("*").
		From("organizations").
		Where("id = ? AND NOT deleted", id).
		LoadOneContext(ctx, &org)

	if err == nil {
		return org, nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.Organization{}, dbmodels.ErrNotFound
	}

	return dbmodels.Organization{}, fmt.Errorf("failed to load organization: %w", err)
}

func (r *repo) Create(ctx context.Context, model dbmodels.OrganizationCreation) (dbmodels.Organization, error) {
	var err error
	defer dal.LogOptionalError(r.l, "organization", err)

	var org dbmodels.Organization

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf("failed to create organization: %w", err)
	}
	defer tx.RollbackUnlessCommitted()

	err = tx.InsertInto("organizations").
		Columns("name", "description").
		Record(model).
		Returning("id", "name", "description", "is_default", "deleted").
		LoadContext(ctx, &org)

	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf("failed to create organization: %w", err)
	}

	_, err = tx.InsertInto("server_groups").
		Columns("name", "description", "organization_id", "is_default").
		Values("Default group for "+org.Name, org.Description, org.ID, true).
		ExecContext(ctx)

	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf("failed to create organization: %w", err)
	}

	return org, tx.Commit()
}

func (r *repo) Update(ctx context.Context, id uuid.UUID, model dbmodels.OrganizationUpdate) (dbmodels.Organization, error) {
	var err error
	defer dal.LogOptionalError(r.l, "organization", err)

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf("failed to update organization: %w", err)
	}
	defer tx.RollbackUnlessCommitted()

	updateMap := make(map[string]interface{})

	if model.Name != nil && *model.Name != "" {
		updateMap["name"] = model.Name
	}
	if model.Description != nil && *model.Description != "" {
		updateMap["description"] = model.Description
	}

	if len(updateMap) == 0 {
		return dbmodels.Organization{}, dbmodels.ErrBadValue
	}

	updateStatement := tx.Update("organizations").SetMap(updateMap).Where("id = ? AND NOT deleted", id)
	_, err = updateStatement.ExecContext(ctx)
	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf("failed to update organization: %w", err)
	}

	var org dbmodels.Organization
	err = tx.Select("*").From("organizations").Where("id = ?", id).LoadOneContext(ctx, &org)
	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf("failed to update organization: %w", err)
	}

	return org, tx.Commit()
}

func (r *repo) Delete(ctx context.Context, id uuid.UUID) error {
	var err error
	defer dal.LogOptionalError(r.l, "organization", err)

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to delete organization: %w", err)
	}
	defer tx.RollbackUnlessCommitted()

	var defaultGroupID uuid.UUID
	err = tx.Select("gr.id").
		From(dbr.I("organizations").As("org")).
		Join(dbr.I("server_groups").As("gr"), "gr.organization_id = org.id").
		Where("org.is_default AND gr.is_default").
		LoadOneContext(ctx, &defaultGroupID)
	if err != nil {
		return fmt.Errorf("failed to delete organization: %w", err)
	}

	groupsSubquery := tx.Select("id").
		From("server_groups").
		Where("organization_id = ?", id)

	_, err = tx.Update("wash_servers").
		Where("group_id IN ?", groupsSubquery).
		Set("group_id", defaultGroupID).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete organization: %w", err)
	}

	_, err = tx.
		Update("organizations").
		Where("id = ? AND NOT deleted AND NOT is_default", id).
		Set("deleted", true).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete organization: %w", err)
	}

	return tx.Commit()
}

func (r *repo) AssignManager(ctx context.Context, organizationID uuid.UUID, userID string) error {
	var err error
	defer dal.LogOptionalError(r.l, "organization", err)

	_, err = r.db.NewSession(nil).
		InsertInto("organization_managers").
		Columns("user_id", "organization_id").
		Values(userID, organizationID).
		ExecContext(ctx)

	if err == nil {
		return nil
	}

	if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
		return dbmodels.ErrAlreadyExists
	}

	return fmt.Errorf("failed to assign manager to organization: %w", err)
}

func (r *repo) RemoveManager(ctx context.Context, organizationID uuid.UUID, userID string) error {
	var err error
	defer dal.LogOptionalError(r.l, "organization", err)

	res, err := r.db.NewSession(nil).
		DeleteFrom("organization_managers").
		Where("user_id = ? AND organization_id = ?", userID, organizationID).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf("failed to remove manager from organization: %w", err)
	}

	if count, err := res.RowsAffected(); err == nil && count == 0 {
		return dbmodels.ErrNotFound
	}

	return fmt.Errorf("failed to remove manager from organization: %w", err)
}

func (r *repo) IsUserManager(ctx context.Context, organizationID uuid.UUID, userID string) (bool, error) {
	var isUserManager bool
	var err error
	defer dal.LogOptionalError(r.l, "organization", err)

	_, err = r.db.NewSession(nil).
		Select("true").
		From("organization_managers").
		Where("user_id = ? AND organization_id = ?", userID, organizationID).
		LoadContext(ctx, &isUserManager)

	if errors.Is(err, dbr.ErrNotFound) {
		return false, nil
	}

	return isUserManager, err
}
