package serverGroups

import (
	"context"
	"errors"
	"washBonus/internal/conversions"
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *serverGroupService) Get(ctx context.Context, userID string, filter entity.ServerGroupFilter) ([]entity.ServerGroup, error) {
	groups, err := s.groupRepo.Get(ctx, userID, conversions.ServerGroupFilterToDB(filter))
	if err != nil {
		return nil, err
	}

	return conversions.ServerGroupsFromDB(groups), nil
}

func (s *serverGroupService) GetById(ctx context.Context, id uuid.UUID) (entity.ServerGroup, error) {
	group, err := s.groupRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.ServerGroup{}, entity.ErrNotFound
		}
		return entity.ServerGroup{}, err
	}

	return conversions.ServerGroupFromDB(group), nil
}

func (s *serverGroupService) Create(ctx context.Context, groupCreation entity.ServerGroupCreation) (entity.ServerGroup, error) {
	org, err := s.orgRepo.GetById(ctx, groupCreation.OrganizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.ServerGroup{}, entity.ErrNotFound
		}
		return entity.ServerGroup{}, err
	}

	if org.IsDefault {
		return entity.ServerGroup{}, err
	}

	createdGroup, err := s.groupRepo.Create(ctx, conversions.ServerGroupCreationToDb(groupCreation))
	if err != nil {
		return entity.ServerGroup{}, err
	}

	return conversions.ServerGroupFromDB(createdGroup), nil
}

func (s *serverGroupService) Update(ctx context.Context, id uuid.UUID, groupUpdate entity.ServerGroupUpdate) (entity.ServerGroup, error) {
	updatedGroup, err := s.groupRepo.Update(ctx, id, conversions.ServerGroupUpdateToDb(groupUpdate))
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entity.ErrNotFound
		} else if errors.Is(err, dbmodels.ErrBadValue) {
			err = entity.ErrAccessDenied
		}

		return entity.ServerGroup{}, err
	}

	return conversions.ServerGroupFromDB(updatedGroup), nil
}

func (s *serverGroupService) Delete(ctx context.Context, id uuid.UUID) error {
	group, err := s.groupRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.ErrNotFound
		}
		return err
	}

	if group.IsDefault {
		return entity.ErrAccessDenied
	}

	return s.groupRepo.Delete(ctx, id)
}
