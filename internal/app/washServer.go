package app

import (
	"context"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
)

type (
	WashServerController interface {
		GetWashServerById(ctx context.Context, userID string, serverID uuid.UUID) (entity.WashServer, error)
		GetWashServers(ctx context.Context, userID string, pagination entity.Pagination) ([]entity.WashServer, error)
		CreateWashServer(ctx context.Context, userID string, creationEntity entity.CreateWashServer) (entity.WashServer, error)
		UpdateWashServer(ctx context.Context, userID string, serverID uuid.UUID, updateEntity entity.UpdateWashServer) (entity.WashServer, error)
		DeleteWashServer(ctx context.Context, userID string, serverID uuid.UUID) error
	}

	WashServerService interface {
		GetWashServerById(ctx context.Context, serverID uuid.UUID) (entity.WashServer, error)
		GetWashServers(ctx context.Context, pagination entity.Pagination) ([]entity.WashServer, error)
		CreateWashServer(ctx context.Context, userID string, creationEntity entity.CreateWashServer) (entity.WashServer, error)
		UpdateWashServer(ctx context.Context, serverID uuid.UUID, updateEntity entity.UpdateWashServer) (entity.WashServer, error)
		DeleteWashServer(ctx context.Context, serverID uuid.UUID) error
	}

	WashServerRepo interface {
		GetWashServerById(ctx context.Context, serverID uuid.UUID) (entity.WashServer, error)
		GetWashServers(ctx context.Context, pagination entity.Pagination) ([]entity.WashServer, error)
		CreateWashServer(ctx context.Context, userID string, creationEntity entity.CreateWashServer) (entity.WashServer, error)
		UpdateWashServer(ctx context.Context, serverID uuid.UUID, updateEntity entity.UpdateWashServer) (entity.WashServer, error)
		DeleteWashServer(ctx context.Context, serverID uuid.UUID) error
	}
)
