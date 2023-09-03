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

	err = r.db.NewSession(nil).
		InsertInto("organizations").
		Columns("name", "description").
		Record(model).
		Returning("id", "name", "description", "is_default", "deleted").
		LoadContext(ctx, &org)

	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf("failed to create organization: %w", err)
	}

	return org, nil
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

	if model.Name != nil {
		updateMap["name"] = model.Name
	}
	if model.Description != nil {
		updateMap["description"] = model.Description
	}

	if len(updateMap) == 0 {
		return dbmodels.Organization{}, dbmodels.ErrBadValue
	}

	_, err = tx.Update("organizations").
		SetMap(updateMap).
		Where("id = ? AND NOT deleted", id).
		ExecContext(ctx)
	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf("failed to update organization: %w", err)
	}

	var org dbmodels.Organization
	err = tx.Select("*").
		From("organizations").
		Where("id = ?", id).
		LoadOneContext(ctx, &org)
	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf("failed to update organization: %w", err)
	}

	return org, tx.Commit()
}

func (r *repo) Delete(ctx context.Context, id uuid.UUID) error {
	var err error
	defer dal.LogOptionalError(r.l, "organization", err)
	op := "failed to delete organization: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	err = resetGroupsToDefaultOrganiziation(ctx, tx, id)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = deleteOrganizationWallets(ctx, tx, id)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	_, err = tx.
		Update("organizations").
		Where("id = ? AND NOT deleted AND NOT is_default", id).
		Set("deleted", true).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func resetGroupsToDefaultOrganiziation(ctx context.Context, tx *dbr.Tx, id uuid.UUID) error {
	defaultGroupIDSubquery := tx.Select("gr.id").
		From(dbr.I("organizations").As("org")).
		Join(dbr.I("server_groups").As("gr"), "gr.organization_id = org.id").
		Where("org.is_default AND gr.is_default").
		Limit(1)

	resetedGroupsSubquery := tx.Select("id").
		From("server_groups").
		Where("organization_id = ?", id)

	_, err := tx.Update("wash_servers").
		Where("group_id IN ?", resetedGroupsSubquery).
		Set("group_id", defaultGroupIDSubquery).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func deleteOrganizationWallets(ctx context.Context, tx *dbr.Tx, id uuid.UUID) error {
	defaultOrgIDSubquery := tx.Select("id").
		From("organizations").
		Where("is_default").
		Limit(1)

	deletedWalletsBalancesSubquery := tx.Select("user_id, balance").
		From("wallets").
		Where("organization_id = ?", id)

	var userBalances []struct {
		UserID  string
		Balance float64
	}
	_, err := deletedWalletsBalancesSubquery.LoadContext(ctx, &userBalances)
	if err != nil {
		return err
	}

	for _, userBalance := range userBalances {
		_, err := tx.Update("wallets").
			Set("balance", dbr.Expr("balance + ?", userBalance.Balance)).
			Where("user_id = ? AND organization_id = ?", userBalance.UserID, defaultOrgIDSubquery).
			ExecContext(ctx)
		if err != nil {
			return err
		}
	}

	_, err = tx.Update("wallets").
		Where("organization_id = ?", id).
		Set("deleted", true).
		Set("balance", 0).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
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

func (r *repo) GetSettingsForOrganization(ctx context.Context, organizationID uuid.UUID) (dbmodels.OrganizationSettings, error) {
	var err error
	defer dal.LogOptionalError(r.l, "organization_settings", err)

	var settings dbmodels.OrganizationSettings

	err = r.db.NewSession(nil).
		SelectBySql(`
			SELECT id, organization_id, FLOOR(EXTRACT(EPOCH FROM processing_delay) / 60) AS processing_delay, bonus_percentage 
			FROM organization_settings 
			WHERE organization_id = ?
			LIMIT 1
		`, organizationID).
		LoadOneContext(ctx, &settings)

	if err == nil {
		return settings, nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.OrganizationSettings{}, dbmodels.ErrNotFound
	}

	return dbmodels.OrganizationSettings{}, fmt.Errorf("failed to load organization settings: %w", err)
}

func (r *repo) UpdateSettingsForOrganization(ctx context.Context, organizationID uuid.UUID, model dbmodels.OrganizationSettingsUpdate) (dbmodels.OrganizationSettings, error) {
	var err error
	defer dal.LogOptionalError(r.l, "organization_settings", err)

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.OrganizationSettings{}, fmt.Errorf("failed to update organization settings: %w", err)
	}
	defer tx.RollbackUnlessCommitted()

	updateMap := make(map[string]interface{})

	if model.ReportsProcessingDelayMinutes != nil {
		updateMap["processing_delay"] = fmt.Sprintf("%d minutes", *model.ReportsProcessingDelayMinutes)
	}

	if model.BonusPercentage != nil {
		updateMap["bonus_percentage"] = model.BonusPercentage
	}

	if len(updateMap) == 0 {
		return dbmodels.OrganizationSettings{}, dbmodels.ErrBadValue
	}

	_, err = tx.Update("organization_settings").
		SetMap(updateMap).
		Where("organization_id = ?", organizationID).
		ExecContext(ctx)
	if err != nil {
		return dbmodels.OrganizationSettings{}, fmt.Errorf("failed to update organization settings: %w", err)
	}

	var settings dbmodels.OrganizationSettings
	err = tx.SelectBySql(`
		SELECT id, organization_id, FLOOR(EXTRACT(EPOCH FROM processing_delay) / 60) AS processing_delay, bonus_percentage 
		FROM organization_settings 
		WHERE organization_id = ?
	`, organizationID).
		LoadOneContext(ctx, &settings)

	if err != nil {
		return dbmodels.OrganizationSettings{}, fmt.Errorf("failed to update organization settings: %w", err)
	}

	return settings, tx.Commit()
}
