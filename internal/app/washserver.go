package app

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"

	uuid "github.com/satori/go.uuid"
)

type (
	WashServerController interface {
		GetWashServerById(ctx Ctx, auth AdminAuth, serverID uuid.UUID) (entities.WashServer, error)
		GetWashServers(ctx Ctx, auth AdminAuth, filter entities.WashServerFilter) ([]entities.WashServer, error)
		CreateWashServer(ctx Ctx, auth AdminAuth, creationEntity entities.WashServerCreation) (entities.WashServer, error)
		UpdateWashServer(ctx Ctx, auth AdminAuth, serverID uuid.UUID, updateEntity entities.WashServerUpdate) (entities.WashServer, error)
		DeleteWashServer(ctx Ctx, auth AdminAuth, serverID uuid.UUID) error

		AssignToServerGroup(ctx Ctx, auth AdminAuth, serverID uuid.UUID, groupID uuid.UUID) error
	}

	WashServerService interface {
		GetWashServerById(ctx Ctx, serverID uuid.UUID) (entities.WashServer, error)
		GetWashServers(ctx Ctx, filter entities.WashServerFilter) ([]entities.WashServer, error)
		CreateWashServer(ctx Ctx, userID string, creationEntity entities.WashServerCreation) (entities.WashServer, error)
		UpdateWashServer(ctx Ctx, serverID uuid.UUID, updateEntity entities.WashServerUpdate) (entities.WashServer, error)
		DeleteWashServer(ctx Ctx, serverID uuid.UUID) error

		AssignToServerGroup(ctx Ctx, serverID uuid.UUID, groupID uuid.UUID) error
	}

	WashServerRepo interface {
		GetWashServerById(ctx Ctx, serverID uuid.UUID) (dbmodels.WashServer, error)
		GetWashServers(ctx Ctx, filter dbmodels.WashServerFilter) ([]dbmodels.WashServer, error)
		CreateWashServer(ctx Ctx, userID string, creationEntity dbmodels.WashServerCreation) (dbmodels.WashServer, error)
		UpdateWashServer(ctx Ctx, serverID uuid.UUID, updateEntity dbmodels.WashServerUpdate) (dbmodels.WashServer, error)
		DeleteWashServer(ctx Ctx, serverID uuid.UUID) error

		AssignToServerGroup(ctx Ctx, serverID uuid.UUID, groupID uuid.UUID) error
	}
)
