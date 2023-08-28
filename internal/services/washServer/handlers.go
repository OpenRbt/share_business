package washServer

import (
	"context"
	"errors"
	"washBonus/internal/conversions"
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *washService) GetWashServerById(ctx context.Context, serverID uuid.UUID) (entity.WashServer, error) {
	server, err := s.washServerRepo.GetWashServerById(ctx, serverID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entity.ErrNotFound
		}
		return entity.WashServer{}, err
	}

	return conversions.WashServerFromDB(server), nil
}

func (s *washService) GetWashServers(ctx context.Context, filter entity.WashServerFilter) ([]entity.WashServer, error) {
	servers, err := s.washServerRepo.GetWashServers(ctx, conversions.WashServerFilterToDB(filter))
	if err != nil {
		return nil, err
	}

	return conversions.WashServerListFromDB(servers), nil
}

func (s *washService) GetForManager(ctx context.Context, userID string, filter entity.WashServerFilter) ([]entity.WashServer, error) {
	servers, err := s.washServerRepo.GetForManager(ctx, userID, conversions.WashServerFilterToDB(filter))
	if err != nil {
		return nil, err
	}

	return conversions.WashServerListFromDB(servers), nil
}

func (s *washService) CreateWashServer(ctx context.Context, userID string, creationEntity entity.WashServerCreation) (entity.WashServer, error) {
	server, err := s.washServerRepo.CreateWashServer(ctx, userID, conversions.WashServerCreationToDb(creationEntity))
	if err != nil {
		return entity.WashServer{}, err
	}

	return conversions.WashServerFromDB(server), nil
}

func (s *washService) UpdateWashServer(ctx context.Context, serverID uuid.UUID, updateEntity entity.WashServerUpdate) (entity.WashServer, error) {
	server, err := s.washServerRepo.UpdateWashServer(ctx, serverID, conversions.WashServerUpdateToDb(updateEntity))
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entity.ErrNotFound
		}
		return entity.WashServer{}, err
	}

	return conversions.WashServerFromDB(server), nil
}

func (s *washService) DeleteWashServer(ctx context.Context, serverID uuid.UUID) error {
	return s.washServerRepo.DeleteWashServer(ctx, serverID)
}

func (s *washService) AssignToServerGroup(ctx context.Context, serverID uuid.UUID, groupID uuid.UUID) error {
	server, err := s.washServerRepo.GetWashServerById(ctx, serverID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.ErrNotFound
		}
		return err
	}

	if server.GroupID == groupID {
		return entity.ErrBadRequest
	}

	_, err = s.serverGroupRepo.GetById(ctx, groupID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.ErrNotFound
		}
		return err
	}

	return s.washServerRepo.AssignToServerGroup(ctx, serverID, groupID)
}
