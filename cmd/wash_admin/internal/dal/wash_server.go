package dal

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"wash_admin/internal/app"
	"wash_admin/internal/conversions"
	"wash_admin/internal/dal/dbmodels"

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

func (s *Storage) RegisterWashServer(ctx context.Context, userID string, newServer app.RegisterWashServer) (app.WashServer, error) {
	var server dbmodels.WashServer

	err := s.db.NewSession(nil).
		InsertInto("wash_servers").
		Columns("title", "description", "service_key", "created_by").
		Record(dbmodels.RegisterWashServer{
			Title:       newServer.Title,
			Description: newServer.Description,
			ServiceKey:  s.generateNewServiceKey(),
			CreatedBy:   userID,
		}).Returning("id", "title", "description", "service_key", "created_by").
		LoadContext(ctx, &server)

	if err != nil {
		return app.WashServer{}, err
	}

	return conversions.WashServerFromDB(server), err
}

func (s *Storage) GetWashServer(ctx context.Context, id uuid.UUID) (app.WashServer, error) {
	var dbWashServer dbmodels.WashServer

	err := s.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("id = ? AND NOT deleted", uuid.NullUUID{UUID: id, Valid: true}).
		LoadOneContext(ctx, &dbWashServer)

	switch {
	case err == nil:
		return conversions.WashServerFromDB(dbWashServer), err
	case errors.Is(err, dbr.ErrNotFound):
		return app.WashServer{}, app.ErrNotFound
	default:
		return app.WashServer{}, err
	}
}

func (s *Storage) UpdateWashServer(ctx context.Context, updateWashServer app.UpdateWashServer) error {
	dbUpdateWashServer := conversions.UpdateWashServerToDb(updateWashServer)

	session := s.db.NewSession(nil)

	updateStatement := session.
		Update("wash_servers").
		Where("id = ?", dbUpdateWashServer.ID)

	if dbUpdateWashServer.Name != nil {
		updateStatement = updateStatement.Set("title", dbUpdateWashServer.Name)
	}
	if dbUpdateWashServer.Description != nil {
		updateStatement = updateStatement.Set("description", dbUpdateWashServer.Description)
	}

	_, err := updateStatement.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
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

func (s *Storage) GetWashServerList(ctx context.Context, pagination app.Pagination) ([]app.WashServer, error) {
	var dbWashServerList []dbmodels.WashServer

	count, err := s.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("NOT DELETED").
		Limit(uint64(pagination.Limit)).
		Offset(uint64(pagination.Offset)).
		LoadContext(ctx, &dbWashServerList)

	if err != nil {
		return []app.WashServer{}, err
	}

	if count == 0 {
		return []app.WashServer{}, app.ErrNotFound
	}

	washServerListFromDB := conversions.WashServerListFromDB(dbWashServerList)

	return washServerListFromDB, nil
}
