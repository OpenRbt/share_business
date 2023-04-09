package dal

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"wash_admin/internal/conversions"
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity"
	"wash_admin/internal/entity/vo"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

func (s *Storage) generateNewServiceKey() string {
	data := make([]byte, 10)

	_, err := rand.Read(data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", sha256.Sum256(data))
}

func (s *Storage) RegisterWashServer(ctx context.Context, owner uuid.UUID, newWashServer vo.RegisterWashServer) (entity.WashServer, error) {
	var registredServer dbmodels.WashServer

	err := s.db.NewSession(nil).
		InsertInto("wash_servers").
		Columns("title", "description", "owner", "service_key").
		Record(dbmodels.RegisterWashServer{
			Title:       newWashServer.Title,
			Description: newWashServer.Description,
			Owner: uuid.NullUUID{
				UUID:  owner,
				Valid: true,
			},
			ServiceKey: s.generateNewServiceKey(),
		}).Returning("id", "title", "description", "owner", "service_key").
		LoadContext(ctx, &registredServer)

	if err != nil {
		return entity.WashServer{}, err
	}

	return conversions.WashServerFromDB(registredServer), err
}

func (s *Storage) GetWashServer(ctx context.Context, ownerId uuid.UUID, id uuid.UUID) (entity.WashServer, error) {
	var dbWashServer dbmodels.WashServer

	err := s.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("id = ? AND owner = ? AND NOT deleted", uuid.NullUUID{UUID: id, Valid: true}, uuid.NullUUID{UUID: ownerId, Valid: true}).
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

func (s *Storage) UpdateWashServer(ctx context.Context, updateWashServer vo.UpdateWashServer) error {
	dbUpdateWashServer := conversions.UpdateWashServerToDb(updateWashServer)

	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	updateStatement := tx.
		Update("wash_servers").
		Where("id = ?", dbUpdateWashServer.ID)

	if dbUpdateWashServer.Name != nil {
		updateStatement = updateStatement.Set("title", dbUpdateWashServer.Name)
	}
	if dbUpdateWashServer.Description != nil {
		updateStatement = updateStatement.Set("description", dbUpdateWashServer.Description)
	}

	_, err = updateStatement.ExecContext(ctx)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Storage) DeleteWashServer(ctx context.Context, id uuid.UUID) error {
	dbDeleteWashServer := conversions.DeleteWashServerToDB(id)

	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	deleteStatement := tx.
		Update("wash_servers").
		Where("id = ? AND NOT DELETED", dbDeleteWashServer.ID).
		Set("deleted", true)

	_, err = deleteStatement.ExecContext(ctx)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Storage) GetWashServerList(ctx context.Context, ownerId uuid.UUID, pagination vo.Pagination) ([]entity.WashServer, error) {
	var dbWashServerList []dbmodels.WashServer

	count, err := s.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("NOT DELETED AND owner = ?", uuid.NullUUID{UUID: ownerId, Valid: true}).
		Limit(uint64(pagination.Limit)).
		Offset(uint64(pagination.Offset)).
		LoadContext(ctx, &dbWashServerList)

	if err != nil {
		return []entity.WashServer{}, err
	}

	if count == 0 {
		return []entity.WashServer{}, dbr.ErrNotFound
	}

	washServerListFromDB := conversions.WashServerListFromDB(dbWashServerList)

	return washServerListFromDB, nil
}
