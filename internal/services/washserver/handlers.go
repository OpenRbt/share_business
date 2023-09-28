package washserver

import (
	"context"
	"errors"
	"fmt"
	"washbonus/internal/conversions"
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"

	uuid "github.com/satori/go.uuid"
)

func (s *washService) GetWashServerById(ctx context.Context, serverID uuid.UUID) (entities.WashServer, error) {
	server, err := s.washServerRepo.GetWashServerById(ctx, serverID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entities.ErrNotFound
		}
		return entities.WashServer{}, err
	}

	return conversions.WashServerFromDB(server), nil
}

func (s *washService) GetWashServers(ctx context.Context, filter entities.WashServerFilter) ([]entities.WashServer, error) {
	servers, err := s.washServerRepo.GetWashServers(ctx, conversions.WashServerFilterToDB(filter))
	if err != nil {
		return nil, err
	}

	return conversions.WashServerListFromDB(servers), nil
}

func (s *washService) CreateWashServer(ctx context.Context, userID string, creationEntity entities.WashServerCreation) (entities.WashServer, error) {
	server, err := s.washServerRepo.CreateWashServer(ctx, userID, conversions.WashServerCreationToDb(creationEntity))
	if err != nil {
		return entities.WashServer{}, err
	}

	return conversions.WashServerFromDB(server), nil
}

func (s *washService) UpdateWashServer(ctx context.Context, serverID uuid.UUID, updateEntity entities.WashServerUpdate) (entities.WashServer, error) {
	server, err := s.washServerRepo.UpdateWashServer(ctx, serverID, conversions.WashServerUpdateToDb(updateEntity))
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entities.ErrNotFound
		} else if errors.Is(err, dbmodels.ErrBadRequest) {
			err = entities.ErrForbidden
		}
		return entities.WashServer{}, err
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
			return fmt.Errorf("wash server %w", entities.ErrNotFound)
		}
		return err
	}

	if server.GroupID == groupID {
		return nil
	}

	group, err := s.serverGroupRepo.GetById(ctx, groupID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return fmt.Errorf("server group %w", entities.ErrNotFound)
		}
		return err
	}

	if server.OrganizationID != group.OrganizationID {
		return fmt.Errorf("can't assign server to another organization: %w", entities.ErrForbidden)
	}

	return s.washServerRepo.AssignToServerGroup(ctx, serverID, groupID)
}
