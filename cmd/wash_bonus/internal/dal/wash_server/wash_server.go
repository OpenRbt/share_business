package wash_server

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/dal"
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/entity/vo"
)

func (r *repo) GetWashServer(ctx context.Context, id uuid.UUID) (e entity.WashServer, err error) {
	defer dal.LogOptionalError(r.l, "wash_server", err)

	var res dbmodels.WashServer

	err = r.db.NewSession(nil).
		Select("*").
		From("wash_servers").
		Where("id = ?", uuid.NullUUID{
			UUID:  id,
			Valid: true,
		}).
		LoadOneContext(ctx, &res)
	if err != nil {
		return
	}

	e = conversions.WashServerFromDB(res)

	return
}

func (r *repo) CreateWashServer(ctx context.Context, server entity.WashServer) (e entity.WashServer, err error) {
	defer dal.LogOptionalError(r.l, "wash_server", err)

	var id uuid.NullUUID

	err = r.db.NewSession(nil).
		InsertInto("wash_servers").
		Record(dbmodels.WashServer{
			ID:          uuid.NullUUID{UUID: server.Id, Valid: true},
			Title:       server.Title,
			Description: server.Description,
		}).
		Returning("id").
		LoadContext(ctx, &id)
	if err != nil {
		return
	}

	e, err = r.GetWashServer(ctx, id.UUID)

	return
}

func (r *repo) UpdateWashServer(ctx context.Context, update vo.WashServerUpdate) (err error) {
	defer dal.LogOptionalError(r.l, "wash_server", err)

	updateStmt := r.db.NewSession(nil).
		Update("wash_servers").
		Where("id = ?", uuid.NullUUID{
			UUID:  update.ID,
			Valid: true,
		})

	if update.Title != nil {
		updateStmt.Set("title", *update.Title)
	}

	if update.Description != nil {
		updateStmt.Set("description", *update.Description)
	}

	if update.Deleted != nil {
		updateStmt.Set("deleted", *update.Deleted)
	}

	_, err = updateStmt.ExecContext(ctx)

	return
}
