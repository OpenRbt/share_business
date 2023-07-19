package washServer

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"washBonus/internal/conversions"
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"

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

func (r *repo) CreateWashServer(ctx context.Context, userID string, creationEntity entity.CreateWashServer) (entity.WashServer, error) {
	var server dbmodels.WashServer

	err := r.db.NewSession(nil).
		InsertInto("wash_servers").
		Columns("title", "description", "service_key", "created_by").
		Record(dbmodels.RegisterWashServer{
			Title:       creationEntity.Title,
			Description: creationEntity.Description,
			ServiceKey:  r.generateNewServiceKey(),
			CreatedBy:   userID,
		}).Returning("id", "title", "description", "service_key", "created_by").
		LoadContext(ctx, &server)

	if err != nil {
		return entity.WashServer{}, err
	}

	return conversions.WashServerFromDB(server), err
}

func (r *repo) GetWashServerById(ctx context.Context, serverID uuid.UUID) (entity.WashServer, error) {
	var dbWashServer dbmodels.WashServer

	err := r.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("id = ? AND NOT deleted", uuid.NullUUID{UUID: serverID, Valid: true}).
		LoadOneContext(ctx, &dbWashServer)

	switch {
	case err == nil:
		return conversions.WashServerFromDB(dbWashServer), err
	case errors.Is(err, dbr.ErrNotFound):
		return entity.WashServer{}, entity.ErrNotFound
	default:
		return entity.WashServer{}, err
	}
}

func (r *repo) UpdateWashServer(ctx context.Context, serverID uuid.UUID, updateEntity entity.UpdateWashServer) (entity.WashServer, error) {
	dbUpdateWashServer := conversions.UpdateWashServerToDb(updateEntity)
	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return entity.WashServer{}, err
	}

	updateMap := make(map[string]interface{})

	if dbUpdateWashServer.Name != nil && *dbUpdateWashServer.Name != "" {
		updateMap["title"] = dbUpdateWashServer.Name
	}
	if dbUpdateWashServer.Description != nil && *dbUpdateWashServer.Description != "" {
		updateMap["description"] = dbUpdateWashServer.Description
	}

	if len(updateMap) == 0 {
		return entity.WashServer{}, entity.ErrBadRequest
	}

	updateStatement := tx.Update("wash_servers").SetMap(updateMap).Where("id = ?", serverID)
	_, err = updateStatement.ExecContext(ctx)
	if err != nil {
		return entity.WashServer{}, err
	}

	var washServer entity.WashServer
	err = tx.Select("*").From("wash_servers").Where("id = ?", serverID).LoadOneContext(ctx, &washServer)
	if err != nil {
		return entity.WashServer{}, err
	}

	return washServer, tx.Commit()
}

func (r *repo) DeleteWashServer(ctx context.Context, serverID uuid.UUID) error {
	tx, err := r.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	deleteStatement := tx.
		Update("wash_servers").
		Where("id = ? AND NOT DELETED", serverID).
		Set("deleted", true)

	_, err = deleteStatement.ExecContext(ctx)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repo) GetWashServers(ctx context.Context, pagination entity.Pagination) ([]entity.WashServer, error) {
	var dbWashServerList []dbmodels.WashServer

	_, err := r.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("NOT DELETED").
		Limit(uint64(pagination.Limit)).
		Offset(uint64(pagination.Offset)).
		LoadContext(ctx, &dbWashServerList)

	if err != nil {
		return []entity.WashServer{}, err
	}

	washServerListFromDB := conversions.WashServerListFromDB(dbWashServerList)

	return washServerListFromDB, nil
}
