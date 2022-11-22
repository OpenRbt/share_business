package wash_server

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"wash_bonus/internal/entity"
)

func (s *service) GetWashServer(ctx context.Context, id uuid.UUID) (entity.WashServer, error) {
	return s.repo.GetWashServer(ctx, id)
}
func (s *service) GetWashServerByKey(ctx context.Context, key string) (entity.WashServer, error) {
	//TODO implement me
	panic("implement me")
}
