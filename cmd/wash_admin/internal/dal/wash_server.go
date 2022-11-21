package dal

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"wash_admin/internal/conversions"
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity"
	"wash_admin/internal/entity/vo"

	"github.com/gocraft/dbr/v2"
	uuid "github.com/satori/go.uuid"
)

func (s *Storage) GetWashServer(ctx context.Context, ownerId uuid.UUID, id uuid.UUID) (entity.WashServer, error) {
	var dbWashServer dbmodels.WashServer

	err := s.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("id = ? AND owner = ?", uuid.NullUUID{UUID: id, Valid: true}, uuid.NullUUID{UUID: ownerId, Valid: true}).
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

func (s *Storage) AddWashServer(ctx context.Context, addWashServer vo.AddWashServer, ownerId uuid.UUID) error {
	dbAddWashServer := conversions.AddWashServerToDB(addWashServer)

	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	var washId uuid.NullUUID

	err = tx.
		InsertInto("wash_servers").
		Columns("name", "description").
		Record(dbAddWashServer).
		Returning("id").
		LoadContext(ctx, &washId)

	if err != nil {
		return err
	}

	washKey := fmt.Sprintf("%s:%s", ownerId, washId.UUID)
	h := sha256.New()
	h.Write([]byte(washKey))

	sha256Hash := hex.EncodeToString(h.Sum(nil))

	_, err = tx.
		Update("wash_servers").
		Set("wash_key", sha256Hash).
		Where("id = ?", washId).
		ExecContext(ctx)

	switch {
	case errors.Is(err, dbr.ErrNotFound):
		return entity.ErrNotFound
	default:
		return tx.Commit()
	}
}

func (s *Storage) UpdateWashServer(ctx context.Context, updateWashServer vo.UpdateWashServer) error {
	dbUpdateWashServer := conversions.UpdateWashServerToDb(updateWashServer)

	tx, err := s.db.NewSession(nil).BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	updateStatement := tx.Update("wash_servers").Where("id = ?", dbUpdateWashServer.ID)

	if dbUpdateWashServer.Name != nil {
		updateStatement = updateStatement.Set("name", dbUpdateWashServer.Name)
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
	//TODO: Реализовать метод Delete на уровне БД
	panic("Реализовать метод Delete на уровне БД")
}
