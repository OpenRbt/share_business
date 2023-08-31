package organizations

import (
	"context"
	"errors"
	"washBonus/internal/dal"
	"washBonus/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

func (r *serverGroupRepo) Get(ctx context.Context, userID string, filter dbmodels.ServerGroupFilter) ([]dbmodels.ServerGroup, error) {
	var err error
	defer dal.LogOptionalError(r.l, "server_group", err)

	var groups []dbmodels.ServerGroup

	query := r.db.NewSession(nil).
		Select("gr.*").
		From(dbr.I("server_groups").As("gr")).
		Where("NOT gr.deleted")

	if filter.OrganizationID != uuid.Nil {
		query = query.Where("gr.organization_id = ?", filter.OrganizationID)
	}

	if filter.IsManagedByMe {
		query = query.LeftJoin(dbr.I("organization_managers").As("man"), "gr.organization_id = man.organization_id")
		query = query.Where("man.user_id = ?", userID)
	}

	_, err = query.
		Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		LoadContext(ctx, &groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *serverGroupRepo) GetById(ctx context.Context, id uuid.UUID) (dbmodels.ServerGroup, error) {
	var err error
	defer dal.LogOptionalError(r.l, "server_group", err)

	var group dbmodels.ServerGroup

	err = r.db.NewSession(nil).
		Select("*").
		From("server_groups").
		Where("id = ? AND NOT deleted", id).
		LoadOneContext(ctx, &group)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			err = dbmodels.ErrNotFound
		}
		return dbmodels.ServerGroup{}, err
	}

	return group, nil
}

func (r *serverGroupRepo) Create(ctx context.Context, model dbmodels.ServerGroupCreation) (dbmodels.ServerGroup, error) {
	var err error
	defer dal.LogOptionalError(r.l, "server_group", err)

	var group dbmodels.ServerGroup

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return group, err
	}
	defer tx.RollbackUnlessCommitted()

	var groupCount int

	err = tx.Select("1").
		From("server_groups").
		Where("NOT deleted AND organization_id = ?", model.OrganizationID).
		Limit(1).
		LoadOneContext(ctx, &groupCount)

	if errors.Is(err, dbr.ErrNotFound) {
		groupCount = 0
	} else if err != nil {
		return dbmodels.ServerGroup{}, err
	}

	model.IsDefault = groupCount == 0

	err = tx.InsertInto("server_groups").
		Columns("name", "description", "organization_id", "is_default").
		Record(model).
		Returning("id", "name", "description", "organization_id", "is_default", "deleted").
		LoadContext(ctx, &group)

	if err != nil {
		return dbmodels.ServerGroup{}, err
	}

	return group, tx.Commit()
}

func (r *serverGroupRepo) Update(ctx context.Context, id uuid.UUID, model dbmodels.ServerGroupUpdate) (dbmodels.ServerGroup, error) {
	var err error
	defer dal.LogOptionalError(r.l, "server_group", err)

	var group dbmodels.ServerGroup

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return group, err
	}
	defer tx.RollbackUnlessCommitted()

	isDefaultOrg, err := isDefaultOrganizationByGroup(ctx, tx, id)
	if err != nil {
		return group, err
	}

	if isDefaultOrg {
		return group, dbmodels.ErrBadValue
	}

	updateMap, err := buildUpdateMap(model)
	if err != nil {
		return group, err
	}

	if model.IsDefault != nil {
		err = setDefaultGroup(ctx, tx, id)
		if err != nil {
			return group, err
		}
	}

	_, err = tx.Update("server_groups").
		SetMap(updateMap).
		Where("NOT deleted AND id = ?", id).
		ExecContext(ctx)
	if err != nil {
		return group, err
	}

	err = tx.Select("*").
		From("server_groups").
		Where("id = ?", id).
		LoadOneContext(ctx, &group)
	if err != nil {
		return group, err
	}

	return group, tx.Commit()
}

func isDefaultOrganizationByGroup(ctx context.Context, tx *dbr.Tx, groupID uuid.UUID) (bool, error) {
	var isDefault bool

	err := tx.Select("true").
		From(dbr.I("organizations").As("org")).
		Join(dbr.I("server_groups").As("gr"), "gr.organization_id = org.id").
		Where("org.is_default AND gr.id = ?", groupID).
		LoadOneContext(ctx, &isDefault)

	if !errors.Is(err, dbr.ErrNotFound) {
		return false, err
	}

	return isDefault, nil
}

func buildUpdateMap(model dbmodels.ServerGroupUpdate) (map[string]interface{}, error) {
	updateMap := make(map[string]interface{})

	if model.Name != nil {
		updateMap["name"] = model.Name
	}

	if model.Description != nil {
		updateMap["description"] = model.Description
	}

	if len(updateMap) == 0 {
		return nil, dbmodels.ErrBadValue
	}

	return updateMap, nil
}

func setDefaultGroup(ctx context.Context, tx *dbr.Tx, newDefaultGroupID uuid.UUID) error {
	orgIDSubquery := tx.Select("organization_id").
		From("server_groups").
		Where("id = ?", newDefaultGroupID)

	_, err := tx.Update("server_groups").
		Set("is_default", false).
		Where("is_default").
		Where("organization_id = ?", orgIDSubquery).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Update("server_groups").
		Set("is_default", true).
		Where("id = ?", newDefaultGroupID).
		ExecContext(ctx)

	return err
}

func (r *serverGroupRepo) Delete(ctx context.Context, id uuid.UUID) error {
	var err error
	defer dal.LogOptionalError(r.l, "server_group", err)

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	var defaultGroupID uuid.UUID

	orgIDSubquery := tx.Select("org.id").
		From(dbr.I("organizations").As("org")).
		LeftJoin(dbr.I("server_groups").As("gr"), "gr.organization_id = org.id").
		Where("NOT org.is_default AND gr.id = ?", id)

	err = tx.Select("id").
		From("server_groups").
		Where("is_default AND organization_id = ?", orgIDSubquery).
		LoadOneContext(ctx, &defaultGroupID)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return dbmodels.ErrNotFound
		}
		return err
	}

	_, err = tx.Update("wash_servers").
		Where("group_id = ?", id).
		Set("group_id", defaultGroupID).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Update("server_groups").
		Where("NOT deleted AND id = ? ", id).
		Set("deleted", true).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
