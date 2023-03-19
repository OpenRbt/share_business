package wash_server

import (
	"context"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/entity/vo"

	uuid "github.com/satori/go.uuid"
)

func (s *service) GetWashServer(ctx context.Context, id uuid.UUID) (entity.WashServer, error) {
	return s.repo.GetWashServer(ctx, id)
}

func (s *service) GetWashServers(ctx context.Context) ([]entity.WashServer, error) {
	return s.repo.GetWashServers(ctx)
}

func (s *service) CreateWashServer(ctx context.Context, server entity.WashServer) (entity.WashServer, error) {
	return s.repo.CreateWashServer(ctx, server)
}

func (s *service) UpdateWashServer(ctx context.Context, update vo.WashServerUpdate) error {
	return s.repo.UpdateWashServer(ctx, update)
}
