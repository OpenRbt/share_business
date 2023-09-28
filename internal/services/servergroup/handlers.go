package servergroup

import (
	"context"
	"errors"
	"washbonus/internal/conversions"
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"

	uuid "github.com/satori/go.uuid"
)

func (s *serverGroupService) Get(ctx context.Context, filter entities.ServerGroupFilter) ([]entities.ServerGroup, error) {
	groups, err := s.groupRepo.Get(ctx, conversions.ServerGroupFilterToDB(filter))
	if err != nil {
		return nil, err
	}

	return conversions.ServerGroupsFromDB(groups), nil
}

func (s *serverGroupService) GetById(ctx context.Context, id uuid.UUID) (entities.ServerGroup, error) {
	group, err := s.groupRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ServerGroup{}, entities.ErrNotFound
		}
		return entities.ServerGroup{}, err
	}

	return conversions.ServerGroupFromDB(group), nil
}

func (s *serverGroupService) Create(ctx context.Context, groupCreation entities.ServerGroupCreation) (entities.ServerGroup, error) {
	_, err := s.orgRepo.GetById(ctx, groupCreation.OrganizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ServerGroup{}, entities.ErrNotFound
		}
		return entities.ServerGroup{}, err
	}

	createdGroup, err := s.groupRepo.Create(ctx, conversions.ServerGroupCreationToDb(groupCreation))
	if err != nil {
		return entities.ServerGroup{}, err
	}

	return conversions.ServerGroupFromDB(createdGroup), nil
}

func (s *serverGroupService) Update(ctx context.Context, id uuid.UUID, groupUpdate entities.ServerGroupUpdate) (entities.ServerGroup, error) {
	updatedGroup, err := s.groupRepo.Update(ctx, id, conversions.ServerGroupUpdateToDb(groupUpdate))
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entities.ErrNotFound
		} else if errors.Is(err, dbmodels.ErrBadRequest) {
			err = entities.ErrBadRequest
		}

		return entities.ServerGroup{}, err
	}

	return conversions.ServerGroupFromDB(updatedGroup), nil
}

func (s *serverGroupService) Delete(ctx context.Context, id uuid.UUID) error {
	group, err := s.groupRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ErrNotFound
		}
		return err
	}

	if group.IsDefault {
		return entities.ErrForbidden
	}

	err = s.groupRepo.Delete(ctx, id)
	if errors.Is(err, dbmodels.ErrAlreadyExists) {
		return entities.ErrBadRequest
	}

	return err
}
