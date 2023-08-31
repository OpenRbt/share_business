package app

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
)

type (
	ServerGroupController interface {
		Get(ctx Ctx, auth Auth, filter entity.ServerGroupFilter) ([]entity.ServerGroup, error)
		GetById(ctx Ctx, auth Auth, id uuid.UUID) (entity.ServerGroup, error)
		Create(ctx Ctx, auth Auth, ent entity.ServerGroupCreation) (entity.ServerGroup, error)
		Update(ctx Ctx, auth Auth, id uuid.UUID, ent entity.ServerGroupUpdate) (entity.ServerGroup, error)
		Delete(ctx Ctx, auth Auth, id uuid.UUID) error
	}

	ServerGroupService interface {
		Get(ctx Ctx, userID string, filter entity.ServerGroupFilter) ([]entity.ServerGroup, error)
		GetById(ctx Ctx, id uuid.UUID) (entity.ServerGroup, error)
		Create(ctx Ctx, ent entity.ServerGroupCreation) (entity.ServerGroup, error)
		Update(ctx Ctx, id uuid.UUID, ent entity.ServerGroupUpdate) (entity.ServerGroup, error)
		Delete(ctx Ctx, id uuid.UUID) error
	}

	ServerGroupRepo interface {
		Get(ctx Ctx, userID string, filter dbmodels.ServerGroupFilter) ([]dbmodels.ServerGroup, error)
		GetById(ctx Ctx, id uuid.UUID) (dbmodels.ServerGroup, error)
		Create(ctx Ctx, model dbmodels.ServerGroupCreation) (dbmodels.ServerGroup, error)
		Update(ctx Ctx, id uuid.UUID, model dbmodels.ServerGroupUpdate) (dbmodels.ServerGroup, error)
		Delete(ctx Ctx, id uuid.UUID) error
	}
)
