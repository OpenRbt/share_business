package servergroups

import (
	"context"
	"errors"
	"fmt"
	"washbonus/internal/dal"
	"washbonus/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

const resource = dbmodels.ServerGroupsResource

const GroupColumns = "gr.id, gr.organization_id, gr.name, gr.description, COALESCE(gr.processing_delay, org.processing_delay) AS processing_delay, COALESCE(gr.bonus_percentage, org.bonus_percentage) AS bonus_percentage, COALESCE(gr.utc_offset, org.utc_offset) AS utc_offset, gr.is_default, gr.deleted, gr.version"

func (r *serverGroupRepo) Get(ctx context.Context, filter dbmodels.ServerGroupFilter) ([]dbmodels.ServerGroup, error) {
	op := "failed to get server groups: %w"

	query := r.db.NewSession(nil).
		Select(GroupColumns).
		From(dbr.I("server_groups").As("gr")).
		LeftJoin(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("NOT gr.deleted")

	if filter.OrganizationID != nil {
		query = query.Where("gr.organization_id = ?", filter.OrganizationID)
	}

	var groups []dbmodels.ServerGroup
	_, err := query.
		Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		LoadContext(ctx, &groups)

	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return groups, nil
}

func (r *serverGroupRepo) GetAll(ctx context.Context, pagination dbmodels.Pagination) ([]dbmodels.ServerGroup, error) {
	op := "failed to get all server groups: %w"
	var groups []dbmodels.ServerGroup

	_, err := r.db.NewSession(nil).
		Select(GroupColumns).
		From(dbr.I("server_groups").As("gr")).
		LeftJoin(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("NOT gr.deleted").
		OrderBy("gr.name").
		Limit(uint64(pagination.Limit)).
		Offset(uint64(pagination.Offset)).
		LoadContext(ctx, &groups)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return groups, nil
}

func (r *serverGroupRepo) GetById(ctx context.Context, id uuid.UUID) (dbmodels.ServerGroup, error) {
	op := "failed to get server group by ID: %w"

	var group dbmodels.ServerGroup
	err := r.db.NewSession(nil).
		Select(GroupColumns).
		From(dbr.I("server_groups").As("gr")).
		LeftJoin(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("gr.id = ? AND NOT gr.deleted", id).
		LoadOneContext(ctx, &group)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = dbmodels.ErrNotFound
		}
		return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
	}

	return group, nil
}

func (r *serverGroupRepo) Create(ctx context.Context, model dbmodels.ServerGroupCreation) (dbmodels.ServerGroup, error) {
	op := "failed to create server group: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	var exists int
	err = tx.Select("1").
		From("server_groups").
		Where("NOT deleted AND organization_id = ?", model.OrganizationID).
		Limit(1).
		LoadOneContext(ctx, &exists)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			model.IsDefault = true
		} else {
			return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
		}
	}

	var groupID uuid.UUID
	err = tx.InsertInto("server_groups").
		Columns("name", "description", "utc_offset", "organization_id", "is_default").
		Record(model).
		Returning("id").
		LoadContext(ctx, &groupID)

	if err != nil {
		return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
	}

	var group dbmodels.ServerGroup
	err = tx.Select(GroupColumns).
		From(dbr.I("server_groups").As("gr")).
		LeftJoin(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("gr.id = ? AND NOT gr.deleted", groupID).
		LoadOneContext(ctx, &group)

	if err != nil {
		return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, resource, group.ID.String(), "create", model)
	if err != nil {
		return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
	}

	return group, tx.Commit()
}

func (r *serverGroupRepo) Update(ctx context.Context, id uuid.UUID, model dbmodels.ServerGroupUpdate) (dbmodels.ServerGroup, error) {
	op := "failed to update server group: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	updateMap := dal.ConstructUpdateMap(model)
	if len(updateMap) == 0 {
		return dbmodels.ServerGroup{}, dbmodels.ErrBadRequest
	}

	res, err := tx.Update("server_groups").
		SetMap(updateMap).
		Set("version", dbr.Expr("version + 1")).
		Where("NOT deleted AND id = ?", id).
		ExecContext(ctx)

	if err != nil {
		return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
	}

	if count, err := res.RowsAffected(); err == nil && count == 0 {
		return dbmodels.ServerGroup{}, dbmodels.ErrNotFound
	}

	if err != nil {
		return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
	}

	var group dbmodels.ServerGroup
	err = tx.Select(GroupColumns).
		From(dbr.I("server_groups").As("gr")).
		LeftJoin(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("gr.id = ? AND NOT gr.deleted", id).
		LoadOneContext(ctx, &group)

	if err != nil {
		return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, resource, id.String(), "update", model)
	if err != nil {
		return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
	}

	return group, tx.Commit()
}

func (r *serverGroupRepo) Delete(ctx context.Context, id uuid.UUID) error {
	op := "failed to delete server group: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	var exists int
	err = tx.Select("1").
		From("wash_servers").
		Where("NOT deleted AND group_id = ?", id).
		Limit(1).
		LoadOneContext(ctx, &exists)

	if err != nil && !errors.Is(err, dbr.ErrNotFound) {
		return fmt.Errorf(op, err)
	}

	if exists == 1 {
		return fmt.Errorf("there are exist wash servers: %w", dbmodels.ErrAlreadyExists)
	}

	_, err = tx.Update("server_groups").
		Where("NOT deleted AND NOT is_default AND id = ? ", id).
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

func (r *serverGroupRepo) GetAnyById(ctx context.Context, id uuid.UUID) (dbmodels.ServerGroup, error) {
	op := "failed to get deleted server group by ID: %w"

	var group dbmodels.ServerGroup
	err := r.db.NewSession(nil).
		Select(GroupColumns).
		From(dbr.I("server_groups").As("gr")).
		LeftJoin(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("gr.id = ?", id).
		LoadOneContext(ctx, &group)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = dbmodels.ErrNotFound
		}
		return dbmodels.ServerGroup{}, fmt.Errorf(op, err)
	}

	return group, nil
}
