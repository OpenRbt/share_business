package washServer

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"washBonus/internal/dal"
	"washBonus/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

func (r *repo) generateNewServiceKey() string {
	data := make([]byte, 10)

	_, err := rand.Read(data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", sha256.Sum256(data))
}

func (r *repo) CreateWashServer(ctx context.Context, userID string, creationEntity dbmodels.WashServerCreation) (dbmodels.WashServer, error) {
	var err error
	defer dal.LogOptionalError(r.l, "wash_server", err)

	var server dbmodels.WashServer

	if !creationEntity.GroupID.Valid {
		err = r.db.NewSession(nil).
			Select("gr.id").
			From(dbr.I("server_groups").As("gr")).
			Join(dbr.I("organizations").As("org"), "org.id = gr.organization_id").
			Where("org.is_default AND gr.is_default").
			LoadOneContext(ctx, &creationEntity.GroupID.UUID)

		if err != nil {
			return dbmodels.WashServer{}, err
		}

		creationEntity.GroupID.Valid = true
	}

	err = r.db.NewSession(nil).
		InsertInto("wash_servers").
		Columns("title", "description", "service_key", "created_by", "group_id").
		Record(dbmodels.WashServerCreation{
			Title:       creationEntity.Title,
			Description: creationEntity.Description,
			ServiceKey:  r.generateNewServiceKey(),
			CreatedBy:   userID,
			GroupID:     creationEntity.GroupID,
		}).Returning("id").
		LoadContext(ctx, &server)

	if err != nil {
		return dbmodels.WashServer{}, err
	}

	err = r.db.NewSession(nil).Select("ser.id, ser.title, ser.description, ser.service_key, ser.created_by, ser.group_id, org.id organization_id").
		From(dbr.I("wash_servers").As("ser")).
		Join(dbr.I("server_groups").As("gr"), "ser.group_id = gr.id").
		Join(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("ser.id = ? AND NOT ser.deleted", server.ID).
		LoadOneContext(ctx, &server)

	return server, err
}

func (r *repo) GetWashServerById(ctx context.Context, serverID uuid.UUID) (dbmodels.WashServer, error) {
	var err error
	defer dal.LogOptionalError(r.l, "wash_server", err)

	var dbWashServer dbmodels.WashServer

	err = r.db.NewSession(nil).Select("ser.id, ser.title, ser.description, ser.service_key, ser.created_by, ser.group_id, org.id organization_id").
		From(dbr.I("wash_servers").As("ser")).
		Join(dbr.I("server_groups").As("gr"), "ser.group_id = gr.id").
		Join(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("ser.id = ? AND NOT ser.deleted", serverID).
		LoadOneContext(ctx, &dbWashServer)

	if errors.Is(err, dbr.ErrNotFound) {
		return dbWashServer, dbmodels.ErrNotFound
	}

	return dbWashServer, err
}

func (r *repo) UpdateWashServer(ctx context.Context, serverID uuid.UUID, updateEntity dbmodels.WashServerUpdate) (dbmodels.WashServer, error) {
	var err error
	defer dal.LogOptionalError(r.l, "wash_server", err)

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.WashServer{}, err
	}
	defer tx.RollbackUnlessCommitted()

	updateMap := make(map[string]interface{})

	if updateEntity.Name != nil && *updateEntity.Name != "" {
		updateMap["title"] = updateEntity.Name
	}
	if updateEntity.Description != nil && *updateEntity.Description != "" {
		updateMap["description"] = updateEntity.Description
	}

	if len(updateMap) == 0 {
		return dbmodels.WashServer{}, dbmodels.ErrBadValue
	}

	updateStatement := tx.Update("wash_servers").SetMap(updateMap).Where("id = ?", serverID)
	_, err = updateStatement.ExecContext(ctx)
	if err != nil {
		return dbmodels.WashServer{}, err
	}

	var washServer dbmodels.WashServer
	err = tx.Select("ser.id, ser.title, ser.description, ser.service_key, ser.created_by, ser.group_id, org.id organization_id").
		From(dbr.I("wash_servers").As("ser")).
		Join(dbr.I("server_groups").As("gr"), "ser.group_id = gr.id").
		Join(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("ser.id = ? AND NOT ser.deleted", serverID).
		LoadOneContext(ctx, &washServer)
	if err != nil {
		return dbmodels.WashServer{}, err
	}

	return washServer, tx.Commit()
}

func (r *repo) DeleteWashServer(ctx context.Context, serverID uuid.UUID) error {
	var err error
	defer dal.LogOptionalError(r.l, "wash_server", err)

	_, err = r.db.NewSession(nil).
		Update("wash_servers").
		Where("id = ? AND NOT deleted", serverID).
		Set("deleted", true).
		ExecContext(ctx)

	return err
}

func (r *repo) GetWashServers(ctx context.Context, filter dbmodels.WashServerFilter) ([]dbmodels.WashServer, error) {
	var err error
	defer dal.LogOptionalError(r.l, "wash_server", err)

	query := r.db.NewSession(nil).
		Select("ser.id, ser.title, ser.description, ser.service_key, ser.created_by, ser.group_id, org.id organization_id").
		From(dbr.I("wash_servers").As("ser")).
		Join(dbr.I("server_groups").As("gr"), "ser.group_id = gr.id").
		Join(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("NOT ser.deleted")

	return getServers(ctx, query, filter)
}

func (r *repo) GetForManager(ctx context.Context, userID string, filter dbmodels.WashServerFilter) ([]dbmodels.WashServer, error) {
	var err error
	defer dal.LogOptionalError(r.l, "wash_server", err)

	query := r.db.NewSession(nil).
		Select("ser.id, ser.title, ser.description, ser.service_key, ser.created_by, ser.group_id, org.id organization_id").
		From(dbr.I("wash_servers").As("ser")).
		Join(dbr.I("server_groups").As("gr"), "ser.group_id = gr.id").
		Join(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Join(dbr.I("organization_managers").As("man"), "org.id = man.organization_id").
		Where("NOT ser.deleted AND man.user_id = ?", userID)

	return getServers(ctx, query, filter)
}

func getServers(ctx context.Context, query *dbr.SelectStmt, filter dbmodels.WashServerFilter) ([]dbmodels.WashServer, error) {
	var dbWashServerList []dbmodels.WashServer

	if filter.OrganizationID != uuid.Nil {
		query.Where("org.id = ?", filter.OrganizationID)
	}

	if filter.GroupID != uuid.Nil {
		query.Where("gr.id = ?", filter.GroupID)
	}

	_, err := query.
		Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		LoadContext(ctx, &dbWashServerList)

	return dbWashServerList, err
}

func (r *repo) AssignToServerGroup(ctx context.Context, serverID uuid.UUID, groupID uuid.UUID) error {
	var err error
	defer dal.LogOptionalError(r.l, "wash_server", err)

	_, err = r.db.NewSession(nil).
		Update("wash_servers").
		Where("id = ? AND NOT deleted", serverID).
		Set("group_id", groupID).
		ExecContext(ctx)

	return err
}
