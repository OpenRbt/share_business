package app

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
)

type (
	WashServerController interface {
		GetWashServerById(ctx Ctx, auth Auth, serverID uuid.UUID) (entity.WashServer, error)
		GetWashServers(ctx Ctx, auth Auth, filter entity.WashServerFilter) ([]entity.WashServer, error)
		CreateWashServer(ctx Ctx, auth Auth, creationEntity entity.WashServerCreation) (entity.WashServer, error)
		UpdateWashServer(ctx Ctx, auth Auth, serverID uuid.UUID, updateEntity entity.WashServerUpdate) (entity.WashServer, error)
		DeleteWashServer(ctx Ctx, auth Auth, serverID uuid.UUID) error

		AssignToServerGroup(ctx Ctx, auth Auth, serverID uuid.UUID, groupID uuid.UUID) error
	}

	WashServerService interface {
		GetWashServerById(ctx Ctx, serverID uuid.UUID) (entity.WashServer, error)
		GetWashServers(ctx Ctx, userID string, filter entity.WashServerFilter) ([]entity.WashServer, error)
		CreateWashServer(ctx Ctx, userID string, creationEntity entity.WashServerCreation) (entity.WashServer, error)
		UpdateWashServer(ctx Ctx, serverID uuid.UUID, updateEntity entity.WashServerUpdate) (entity.WashServer, error)
		DeleteWashServer(ctx Ctx, serverID uuid.UUID) error

		AssignToServerGroup(ctx Ctx, serverID uuid.UUID, groupID uuid.UUID) error
	}

	WashServerRepo interface {
		GetWashServerById(ctx Ctx, serverID uuid.UUID) (dbmodels.WashServer, error)
		GetWashServers(ctx Ctx, userID string, filter dbmodels.WashServerFilter) ([]dbmodels.WashServer, error)
		CreateWashServer(ctx Ctx, userID string, creationEntity dbmodels.WashServerCreation) (dbmodels.WashServer, error)
		UpdateWashServer(ctx Ctx, serverID uuid.UUID, updateEntity dbmodels.WashServerUpdate) (dbmodels.WashServer, error)
		DeleteWashServer(ctx Ctx, serverID uuid.UUID) error

		AssignToServerGroup(ctx Ctx, serverID uuid.UUID, groupID uuid.UUID) error
	}
)
