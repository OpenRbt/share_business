package organizations

import (
	"context"
	"washBonus/internal/dal"
	"washBonus/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

func (r *serverGroupRepo) Get(ctx context.Context, filter dbmodels.ServerGroupFilter) ([]dbmodels.ServerGroup, error) {
	var err error
	defer dal.LogOptionalError(r.l, "server_group", err)

	query := r.db.NewSession(nil).
		Select("*").
		From("server_groups").
		Where("NOT deleted")

	if filter.OrganizationID != uuid.Nil {
		query = query.Where("organization_id = ?", filter.OrganizationID)
	}

	return getGroups(ctx, query, filter)
}

func (r *serverGroupRepo) GetForManager(ctx context.Context, userID string, filter dbmodels.ServerGroupFilter) ([]dbmodels.ServerGroup, error) {
	var err error
	defer dal.LogOptionalError(r.l, "server_group", err)

	query := r.db.NewSession(nil).
		Select("gr.*").
		From(dbr.I("server_groups").As("gr")).
		Join(dbr.I("organization_managers").As("man"), "gr.organization_id = man.organization_id").
		Where("NOT gr.deleted AND man.user_id = ?", userID)

	if filter.OrganizationID != uuid.Nil {
		query = query.Where("gr.organization_id = ?", filter.OrganizationID)
	}

	return getGroups(ctx, query, filter)
}

func getGroups(ctx context.Context, query *dbr.SelectStmt, filter dbmodels.ServerGroupFilter) ([]dbmodels.ServerGroup, error) {
	var groups []dbmodels.ServerGroup

	_, err := query.
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
		return dbmodels.ServerGroup{}, err
	}

	return group, nil
}

func (r *serverGroupRepo) Create(ctx context.Context, model dbmodels.ServerGroupCreation) (dbmodels.ServerGroup, error) {
	var err error
	defer dal.LogOptionalError(r.l, "server_group", err)

	var group dbmodels.ServerGroup

	err = r.db.NewSession(nil).
		InsertInto("server_groups").
		Columns("name", "description", "organization_id").
		Record(model).
		Returning("id", "name", "description", "organization_id", "is_default", "deleted").
		LoadContext(ctx, &group)

	if err != nil {
		return dbmodels.ServerGroup{}, err
	}

	return group, err
}

func (r *serverGroupRepo) Update(ctx context.Context, id uuid.UUID, model dbmodels.ServerGroupUpdate) (dbmodels.ServerGroup, error) {
	var err error
	defer dal.LogOptionalError(r.l, "server_group", err)

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.ServerGroup{}, err
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
		return dbmodels.ServerGroup{}, dbmodels.ErrBadValue
	}

	updateStatement := tx.Update("server_groups").SetMap(updateMap).Where("id = ? AND NOT deleted", id)
	_, err = updateStatement.ExecContext(ctx)
	if err != nil {
		return dbmodels.ServerGroup{}, err
	}

	var group dbmodels.ServerGroup
	err = tx.Select("*").From("server_groups").Where("id = ?", id).LoadOneContext(ctx, &group)
	if err != nil {
		return dbmodels.ServerGroup{}, err
	}

	return group, tx.Commit()
}

func (r *serverGroupRepo) Delete(ctx context.Context, id uuid.UUID) error {
	var err error
	defer dal.LogOptionalError(r.l, "server_group", err)

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	var orgID uuid.UUID
	err = tx.Select("org.id").
		From(dbr.I("organizations").As("org")).
		Join(dbr.I("server_groups").As("gr"), "gr.organization_id = org.id").
		Where("gr.id = ?", id).
		LoadOneContext(ctx, &orgID)
	if err != nil {
		return err
	}

	var defaultGroupID uuid.UUID
	err = tx.Select("gr.id").
		From(dbr.I("organizations").As("org")).
		Join(dbr.I("server_groups").As("gr"), "gr.organization_id = org.id").
		Where("org.id = ? AND gr.is_default", orgID).
		LoadOneContext(ctx, &defaultGroupID)
	if err != nil {
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
		Where("id = ? AND NOT deleted AND NOT is_default", id).
		Set("deleted", true).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
