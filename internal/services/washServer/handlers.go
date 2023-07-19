package washServer

import (
	"context"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *washService) GetWashServerById(ctx context.Context, serverID uuid.UUID) (entity.WashServer, error) {
	return s.washServerRepo.GetWashServerById(ctx, serverID)
}

func (s *washService) GetWashServers(ctx context.Context, pagination entity.Pagination) ([]entity.WashServer, error) {
	return s.washServerRepo.GetWashServers(ctx, pagination)
}

func (s *washService) CreateWashServer(ctx context.Context, userID string, creationEntity entity.CreateWashServer) (entity.WashServer, error) {
	return s.washServerRepo.CreateWashServer(ctx, userID, creationEntity)
}

func (s *washService) UpdateWashServer(ctx context.Context, serverID uuid.UUID, updateEntity entity.UpdateWashServer) (entity.WashServer, error) {
	return s.washServerRepo.UpdateWashServer(ctx, serverID, updateEntity)
}

func (s *washService) DeleteWashServer(ctx context.Context, serverID uuid.UUID) error {
	return s.washServerRepo.DeleteWashServer(ctx, serverID)
}
