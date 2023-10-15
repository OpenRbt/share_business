package washservers

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"washbonus/internal/dal"
	"washbonus/internal/dal/dbmodels"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

const resource = dbmodels.WashServersResource

func (r *repo) generateNewServiceKey() string {
	data := make([]byte, 10)

	_, err := rand.Read(data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", sha256.Sum256(data))
}

func (r *repo) CreateWashServer(ctx context.Context, userID string, creationEntity dbmodels.WashServerCreation) (dbmodels.WashServer, error) {
	op := "failed to create wash server: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.WashServer{}, fmt.Errorf(op, err)
	}
	defer tx.RollbackUnlessCommitted()

	var server dbmodels.WashServer
	if creationEntity.GroupID == nil {
		err := tx.Select("gr.id").
			From(dbr.I("server_groups").As("gr")).
			Join(dbr.I("organizations").As("org"), "org.id = gr.organization_id").
			Where("org.is_default AND gr.is_default").
			LoadOneContext(ctx, &creationEntity.GroupID)

		if err != nil {
			return dbmodels.WashServer{}, fmt.Errorf(op, err)
		}
	}

	err = tx.InsertInto("wash_servers").
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
		return dbmodels.WashServer{}, fmt.Errorf(op, err)
	}

	err = tx.
		Select("ser.id, ser.title, ser.description, ser.service_key, ser.created_by, ser.group_id, org.id organization_id").
		From(dbr.I("wash_servers").As("ser")).
		Join(dbr.I("server_groups").As("gr"), "ser.group_id = gr.id").
		Join(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("ser.id = ? AND NOT ser.deleted", server.ID).
		LoadOneContext(ctx, &server)

	if err != nil {
		return dbmodels.WashServer{}, fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, resource, server.ID.String(), "create", creationEntity)
	if err != nil {
		return dbmodels.WashServer{}, fmt.Errorf(op, err)
	}

	return server, tx.Commit()
}

func (r *repo) GetWashServerById(ctx context.Context, serverID uuid.UUID) (dbmodels.WashServer, error) {
	op := "failed to get wash server by ID: %w"

	var dbWashServer dbmodels.WashServer
	err := r.db.NewSession(nil).
		Select("ser.id, ser.title, ser.description, ser.service_key, ser.created_by, ser.group_id, org.id organization_id").
		From(dbr.I("wash_servers").As("ser")).
		Join(dbr.I("server_groups").As("gr"), "ser.group_id = gr.id").
		Join(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("ser.id = ? AND NOT ser.deleted", serverID).
		LoadOneContext(ctx, &dbWashServer)

	if err == nil {
		return dbWashServer, nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return dbmodels.WashServer{}, dbmodels.ErrNotFound
	}

	return dbmodels.WashServer{}, fmt.Errorf(op, err)
}

func (r *repo) UpdateWashServer(ctx context.Context, serverID uuid.UUID, updateEntity dbmodels.WashServerUpdate) (dbmodels.WashServer, error) {
	op := "failed to update wash server: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return dbmodels.WashServer{}, err
	}
	defer tx.RollbackUnlessCommitted()

	updateMap := dal.ConstructUpdateMap(updateEntity)
	if len(updateMap) == 0 {
		return dbmodels.WashServer{}, dbmodels.ErrBadRequest
	}

	_, err = tx.Update("wash_servers").
		SetMap(updateMap).
		Where("id = ?", serverID).
		ExecContext(ctx)

	if err != nil {
		return dbmodels.WashServer{}, fmt.Errorf(op, err)
	}

	var washServer dbmodels.WashServer
	err = tx.
		Select("ser.id, ser.title, ser.description, ser.service_key, ser.created_by, ser.group_id, org.id organization_id").
		From(dbr.I("wash_servers").As("ser")).
		Join(dbr.I("server_groups").As("gr"), "ser.group_id = gr.id").
		Join(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("ser.id = ? AND NOT ser.deleted", serverID).
		LoadOneContext(ctx, &washServer)

	if err != nil {
		return dbmodels.WashServer{}, fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, resource, serverID.String(), "update", updateEntity)
	if err != nil {
		return dbmodels.WashServer{}, fmt.Errorf(op, err)
	}

	return washServer, tx.Commit()
}

func (r *repo) DeleteWashServer(ctx context.Context, serverID uuid.UUID) error {
	op := "failed to delete wash server: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.Update("wash_servers").
		Where("id = ? AND NOT deleted", serverID).
		Set("deleted", true).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, resource, serverID.String(), "delete", nil)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}

func (r *repo) GetWashServers(ctx context.Context, filter dbmodels.WashServerFilter) ([]dbmodels.WashServer, error) {
	op := "failed to get wash servers: %w"

	var dbWashServerList []dbmodels.WashServer
	query := r.db.NewSession(nil).
		Select("ser.id, ser.title, ser.description, ser.service_key, ser.created_by, ser.group_id, org.id organization_id").
		From(dbr.I("wash_servers").As("ser")).
		Join(dbr.I("server_groups").As("gr"), "ser.group_id = gr.id").
		Join(dbr.I("organizations").As("org"), "gr.organization_id = org.id").
		Where("NOT ser.deleted")

	if filter.OrganizationID != nil {
		query = query.Where("org.id = ?", filter.OrganizationID)
	}

	if filter.GroupID != nil {
		query = query.Where("gr.id = ?", filter.GroupID)
	}

	_, err := query.
		Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		LoadContext(ctx, &dbWashServerList)

	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return dbWashServerList, nil
}

func (r *repo) AssignToServerGroup(ctx context.Context, serverID uuid.UUID, groupID uuid.UUID) error {
	op := "failed to assign wash server to group: %w"

	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.Update("wash_servers").
		Where("id = ? AND NOT deleted", serverID).
		Set("group_id", groupID).
		ExecContext(ctx)

	if err != nil {
		return fmt.Errorf(op, err)
	}

	err = dal.WriteAuditLog(ctx, tx, resource, serverID.String(), "assign server to group", struct{ GroupID string }{GroupID: groupID.String()})
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return tx.Commit()
}
