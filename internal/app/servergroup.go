package app

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"

	uuid "github.com/satori/go.uuid"
)

type (
	ServerGroupController interface {
		Get(ctx Ctx, auth AdminAuth, filter entities.ServerGroupFilter) ([]entities.ServerGroup, error)
		GetById(ctx Ctx, auth AdminAuth, id uuid.UUID) (entities.ServerGroup, error)
		Create(ctx Ctx, auth AdminAuth, ent entities.ServerGroupCreation) (entities.ServerGroup, error)
		Update(ctx Ctx, auth AdminAuth, id uuid.UUID, ent entities.ServerGroupUpdate) (entities.ServerGroup, error)
		Delete(ctx Ctx, auth AdminAuth, id uuid.UUID) error
	}

	ServerGroupService interface {
		Get(ctx Ctx, filter entities.ServerGroupFilter) ([]entities.ServerGroup, error)
		GetById(ctx Ctx, id uuid.UUID) (entities.ServerGroup, error)
		Create(ctx Ctx, ent entities.ServerGroupCreation) (entities.ServerGroup, error)
		Update(ctx Ctx, id uuid.UUID, ent entities.ServerGroupUpdate) (entities.ServerGroup, error)
		Delete(ctx Ctx, id uuid.UUID) error
	}

	ServerGroupRepo interface {
		Get(ctx Ctx, filter dbmodels.ServerGroupFilter) ([]dbmodels.ServerGroup, error)
		GetById(ctx Ctx, id uuid.UUID) (dbmodels.ServerGroup, error)
		Create(ctx Ctx, model dbmodels.ServerGroupCreation) (dbmodels.ServerGroup, error)
		Update(ctx Ctx, id uuid.UUID, model dbmodels.ServerGroupUpdate) (dbmodels.ServerGroup, error)
		Delete(ctx Ctx, id uuid.UUID) error
	}
)
