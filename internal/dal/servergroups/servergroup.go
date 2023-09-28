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

func (r *serverGroupRepo) Get(ctx context.Context, filter dbmodels.ServerGroupFilter) ([]dbmodels.ServerGroup, error) {
	op := "failed to get server groups: %w"

	query := r.db.NewSession(nil).
		Select("gr.*").
		From(dbr.I("server_groups").As("gr")).
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

func (r *serverGroupRepo) GetById(ctx context.Context, id uuid.UUID) (dbmodels.ServerGroup, error) {
	op := "failed to get server group by ID: %w"

	var group dbmodels.ServerGroup
	err := r.db.NewSession(nil).
		Select("*").
		From("server_groups").
		Where("id = ? AND NOT deleted", id).
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

	var group dbmodels.ServerGroup
	err = tx.InsertInto("server_groups").
		Columns("name", "description", "organization_id", "is_default").
		Record(model).
		Returning("id", "name", "description", "organization_id", "is_default", "deleted").
		LoadContext(ctx, &group)

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
	err = tx.Select("*").
		From("server_groups").
		Where("id = ?", id).
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
