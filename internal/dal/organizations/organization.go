package organizations

import (
	"context"
	"errors"
	"fmt"
	"washbonus/internal/dal"
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/dal/servergroups"

	"github.com/gocraft/dbr/v2"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var OrgColumns = []string{"id", "name", "display_name", "description", "utc_offset", "is_default", "processing_delay", "bonus_percentage", "deleted", "version"}

const resource = dbmodels.OrganizationsResource

func (r *repo) Get(ctx context.Context, filter dbmodels.OrganizationFilter) ([]dbmodels.Organization, error) {
	var orgs []dbmodels.Organization

	query := r.db.NewSession(nil).
		Select(OrgColumns...).
		From("organizations").
		Where("NOT deleted")

	if len(filter.OrganizationIDs) > 0 {
		query = query.Where("id IN ?", filter.OrganizationIDs)
	}

	_, err := query.
		Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		LoadContext(ctx, &orgs)

	if err != nil {
		return nil, fmt.Errorf("failed to load organizations: %w", err)
	}

	return orgs, nil
}

func (r *repo) GetAll(ctx context.Context, pagination dbmodels.Pagination) ([]dbmodels.Organization, error) {
	var orgs []dbmodels.Organization

	_, err := r.db.NewSession(nil).
		Select(OrgColumns...).
		From("organizations").
		OrderBy("name").
		Limit(uint64(pagination.Limit)).
		Offset(uint64(pagination.Offset)).
		LoadContext(ctx, &orgs)
	if err != nil {
		return nil, fmt.Errorf("failed to load organizations: %w", err)
	}

	return orgs, nil
}

func (r *repo) GetById(ctx context.Context, id uuid.UUID) (dbmodels.Organization, error) {
	var org dbmodels.Organization
	err := r.db.NewSession(nil).
		Select(OrgColumns...).
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

func (r *repo) GetAnyByID(ctx context.Context, id uuid.UUID) (dbmodels.Organization, error) {
	var org dbmodels.Organization
	err := r.db.NewSession(nil).
		Select(OrgColumns...).
		From("organizations").
		Where("id = ?", id).
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
	op := "failed to create organization: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	columns := []string{"name", "display_name", "description"}
	values := []interface{}{model.Name, model.DisplayName, model.Description}

	if model.ReportsProcessingDelayMinutes != nil {
		columns = append(columns, "processing_delay")
		values = append(values, model.ReportsProcessingDelayMinutes)
	}

	if model.BonusPercentage != nil {
		columns = append(columns, "bonus_percentage")
		values = append(values, model.BonusPercentage)
	}

	if model.UTCOffset != nil {
		columns = append(columns, "utc_offset")
		values = append(values, model.UTCOffset)
	}

	var org dbmodels.Organization
	err = tx.InsertInto("organizations").
		Columns(columns...).
		Values(values...).
		Returning(OrgColumns...).
		LoadContext(ctx, &org)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			err = fmt.Errorf("Display name occupied: %w", dbmodels.ErrBadRequest)
		}

		return dbmodels.Organization{}, fmt.Errorf(op, err)
	}

	_, err = tx.InsertInto("server_groups").
		Columns("organization_id", "name", "description", "is_default").
		Values(org.ID, "Default", fmt.Sprintf("Default server group for organization %s", org.DisplayName), true).
		ExecContext(ctx)

	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, resource, org.ID.String(), "create", model)
	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf(op, err)
	}

	return org, tx.Commit()
}

func (r *repo) Update(ctx context.Context, id uuid.UUID, model dbmodels.OrganizationUpdate) (dbmodels.Organization, error) {
	op := "failed to update organization: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	updateMap := dal.ConstructUpdateMap(model)
	if len(updateMap) == 0 {
		return dbmodels.Organization{}, dbmodels.ErrBadRequest
	}

	res, err := tx.Update("organizations").
		SetMap(updateMap).
		Set("version", dbr.Expr("version + 1")).
		Where("NOT deleted AND id = ?", id).
		ExecContext(ctx)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			err = fmt.Errorf("Display name occupied: %w", dbmodels.ErrBadRequest)
		}

		return dbmodels.Organization{}, fmt.Errorf(op, err)
	}

	if count, err := res.RowsAffected(); err == nil && count == 0 {
		return dbmodels.Organization{}, dbmodels.ErrNotFound
	}

	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf(op, err)
	}

	var org dbmodels.Organization
	err = tx.Select(OrgColumns...).
		From("organizations").
		Where("id = ?", id).
		LoadOneContext(ctx, &org)

	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, resource, id.String(), "update", model)
	if err != nil {
		return dbmodels.Organization{}, fmt.Errorf(op, err)
	}

	return org, tx.Commit()
}

func (r *repo) Delete(ctx context.Context, id uuid.UUID) error {
	op := "failed to delete organization: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	if ok, err := areExistNotDefaultGroups(ctx, tx, id); err == nil && ok {
		return fmt.Errorf("there are not deleted groups: %w", dbmodels.ErrBadRequest)
	}

	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = deleteDefaultServerGroup(ctx, tx, id)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = deleteOrganizationWallets(ctx, tx, id)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = blockOrganizationAdmins(ctx, tx, id)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	_, err = tx.
		Update("organizations").
		Where("NOT deleted AND NOT is_default AND id = ?", id).
		Set("deleted", true).
		Set("version", dbr.Expr("version + 1")).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, resource, id.String(), "delete", nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func areExistNotDefaultGroups(ctx context.Context, tx *dbr.Tx, orgID uuid.UUID) (bool, error) {
	var exists int
	err := tx.Select("1").
		From("server_groups").
		Where("NOT deleted AND NOT is_default AND organization_id = ?", orgID).
		Limit(1).
		LoadOneContext(ctx, &exists)

	if err != nil && !errors.Is(err, dbr.ErrNotFound) {
		return false, err
	}

	return exists == 1, nil
}

func deleteDefaultServerGroup(ctx context.Context, tx *dbr.Tx, orgID uuid.UUID) error {
	defaultGroupId := dbr.Select("id").
		From("server_groups").
		Where("NOT deleted AND is_default AND organization_id = ?", orgID)

	var exists int
	err := tx.Select("1").
		From("wash_servers").
		Where("NOT deleted AND group_id = ?", defaultGroupId).
		Limit(1).
		LoadOneContext(ctx, &exists)

	if err != nil && !errors.Is(err, dbr.ErrNotFound) {
		return err
	}

	if exists == 1 {
		return fmt.Errorf("there are not deleted wash servers: %w", dbmodels.ErrBadRequest)
	}

	_, err = tx.Update("server_groups").
		Set("deleted", true).
		Set("version", dbr.Expr("version + 1")).
		Where("is_default AND organization_id = ?", orgID).
		ExecContext(ctx)

	return err
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

func blockOrganizationAdmins(ctx context.Context, tx *dbr.Tx, id uuid.UUID) error {
	_, err := tx.Update("admin_users").
		Set("role", dbmodels.NoAccessRole).
		Set("version", dbr.Expr("version + 1")).
		Where("organization_id = ?", id).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf("failed to block organization users: %w", err)
	}

	return nil
}

func (r *repo) AssignManager(ctx context.Context, organizationID uuid.UUID, userID string) error {
	op := "failed to assign manager to organization: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	orgIDSubquery := dbr.Select("id").
		From("organizations").
		Where("NOT deleted AND id = ?", organizationID)

	res, err := tx.Update("admin_users").
		Set("organization_id", orgIDSubquery).
		Set("version", dbr.Expr("version + 1")).
		Where("id = ?", userID).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	if count, err := res.RowsAffected(); err == nil && count == 0 {
		return dbmodels.ErrNotFound
	}

	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, resource, organizationID.String(), "assign manager", struct{ UserID string }{UserID: userID})
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func (r *repo) RemoveManager(ctx context.Context, organizationID uuid.UUID, userID string) error {
	op := "failed to remove manager from organization: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	res, err := tx.Update("admin_users").
		Set("organization_id", nil).
		Set("version", dbr.Expr("version + 1")).
		Where("organization_id = ? AND id = ?", organizationID, userID).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	if count, err := res.RowsAffected(); err == nil && count == 0 {
		return dbmodels.ErrNotFound
	}

	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, resource, organizationID.String(), "remove manager", struct{ UserID string }{UserID: userID})
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func (r *repo) GetDefaultGroupByOrganizationId(ctx context.Context, id uuid.UUID) (dbmodels.ServerGroup, error) {
	op := "failed to get default server group by organization ID: %w"

	var group dbmodels.ServerGroup
	err := r.db.NewSession(nil).
		Select(servergroups.GroupColumns).
		From(dbr.I("server_groups").As("gr")).
		LeftJoin(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("gr.organization_id = ?", id).
		LoadOneContext(ctx, &group)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = dbmodels.ErrNotFound
		}
		return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
	}

	return group, nil
}

func (r *repo) GetAdminUsersByOrganizationID(ctx context.Context, id uuid.UUID) ([]dbmodels.AdminUser, error) {
	const op = "failed to get admins by organization ID: %w"

	var dbUsers []dbmodels.AdminUser
	_, err := r.db.NewSession(nil).
		Select("*").
		From("admin_users").
		Where("organization_id = ?", id).
		LoadContext(ctx, &dbUsers)

	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return dbUsers, nil
}
