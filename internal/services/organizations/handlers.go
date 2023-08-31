package organizations

import (
	"context"
	"errors"
	"washBonus/internal/app"
	"washBonus/internal/conversions"
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
)

func (s *organizationService) Get(ctx context.Context, userID string, filter entity.OrganizationFilter) ([]entity.Organization, error) {
	orgs, err := s.organizationRepo.Get(ctx, userID, conversions.OrganizationFilterToDB(filter))
	if err != nil {
		return nil, err
	}

	return conversions.OrganizationsFromDB(orgs), nil
}

func (s *organizationService) GetById(ctx context.Context, id uuid.UUID) (entity.Organization, error) {
	org, err := s.organizationRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.Organization{}, entity.ErrNotFound
		}
		return entity.Organization{}, err
	}

	return conversions.OrganizationFromDB(org), nil
}

func (s *organizationService) Create(ctx context.Context, ent entity.OrganizationCreation) (entity.Organization, error) {
	org, err := s.organizationRepo.Create(ctx, conversions.OrganizationCreationToDb(ent))
	if err != nil {
		return entity.Organization{}, err
	}

	return conversions.OrganizationFromDB(org), nil
}

func (s *organizationService) Update(ctx context.Context, id uuid.UUID, ent entity.OrganizationUpdate) (entity.Organization, error) {
	org, err := s.organizationRepo.Update(ctx, id, conversions.OrganizationUpdateToDb(ent))
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.Organization{}, entity.ErrNotFound
		}
		return entity.Organization{}, err
	}

	return conversions.OrganizationFromDB(org), nil
}

func (s *organizationService) Delete(ctx context.Context, id uuid.UUID) error {
	org, err := s.organizationRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.ErrNotFound
		}
		return err
	}

	if org.IsDefault {
		return entity.ErrAccessDenied
	}

	return s.organizationRepo.Delete(ctx, id)
}

func (s *organizationService) AssignManager(ctx context.Context, organizationID uuid.UUID, userID string) error {
	_, err := s.organizationRepo.GetById(ctx, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.ErrNotFound
		}

		return err
	}

	dbUser, err := s.userRepo.GetById(ctx, userID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.ErrNotFound
		}

		return err
	}

	user := conversions.UserFromDb(dbUser)
	if !app.IsEngineer(user) {
		return entity.ErrBadRequest
	}

	err = s.organizationRepo.AssignManager(ctx, organizationID, userID)
	if errors.Is(err, dbmodels.ErrAlreadyExists) {
		return entity.ErrAlreadyExists
	}

	return err
}

func (s *organizationService) RemoveManager(ctx context.Context, organizationID uuid.UUID, userID string) error {
	_, err := s.organizationRepo.GetById(ctx, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.ErrNotFound
		}

		return err
	}

	dbUser, err := s.userRepo.GetById(ctx, userID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.ErrNotFound
		}

		return err
	}

	user := conversions.UserFromDb(dbUser)
	if !app.IsEngineer(user) {
		return entity.ErrBadRequest
	}

	err = s.organizationRepo.RemoveManager(ctx, organizationID, userID)
	if errors.Is(err, dbmodels.ErrNotFound) {
		return entity.ErrNotFound
	}

	return err
}

func (s *organizationService) IsUserManager(ctx context.Context, organizationID uuid.UUID, userID string) (bool, error) {
	return s.organizationRepo.IsUserManager(ctx, organizationID, userID)
}
