package organization

import (
	"context"
	"errors"
	"washbonus/internal/app"
	"washbonus/internal/conversions"
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"

	uuid "github.com/satori/go.uuid"
)

func (s *organizationService) Get(ctx context.Context, filter entities.OrganizationFilter) ([]entities.Organization, error) {
	orgs, err := s.organizationRepo.Get(ctx, conversions.OrganizationFilterToDB(filter))
	if err != nil {
		return nil, err
	}

	return conversions.OrganizationsFromDB(orgs), nil
}

func (s *organizationService) GetAll(ctx context.Context, pagination entities.Pagination) ([]entities.Organization, error) {
	orgs, err := s.organizationRepo.GetAll(ctx, conversions.PaginationToDB(pagination))
	if err != nil {
		return nil, err
	}

	return conversions.OrganizationsFromDB(orgs), nil
}

func (s *organizationService) GetById(ctx context.Context, id uuid.UUID) (entities.Organization, error) {
	org, err := s.organizationRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entities.ErrNotFound
		}
		return entities.Organization{}, err
	}

	return conversions.OrganizationFromDB(org), nil
}

func (s *organizationService) Create(ctx context.Context, ent entities.OrganizationCreation) (entities.Organization, error) {
	org, err := s.organizationRepo.Create(ctx, conversions.OrganizationCreationToDb(ent))
	if err != nil {
		if errors.Is(err, dbmodels.ErrBadRequest) {
			return entities.Organization{}, entities.ErrBadRequest
		}
		return entities.Organization{}, err
	}

	return conversions.OrganizationFromDB(org), nil
}

func (s *organizationService) Update(ctx context.Context, id uuid.UUID, ent entities.OrganizationUpdate) (entities.Organization, error) {
	_, err := s.organizationRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entities.ErrNotFound
		}

		return entities.Organization{}, err
	}

	updatedOrg, err := s.organizationRepo.Update(ctx, id, conversions.OrganizationUpdateToDb(ent))
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.Organization{}, entities.ErrNotFound
		} else if errors.Is(err, dbmodels.ErrBadRequest) {
			return entities.Organization{}, entities.ErrBadRequest
		}
		return entities.Organization{}, err
	}

	return conversions.OrganizationFromDB(updatedOrg), nil
}

func (s *organizationService) Delete(ctx context.Context, id uuid.UUID) error {
	org, err := s.organizationRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ErrNotFound
		}
		return err
	}

	if org.IsDefault {
		return entities.ErrForbidden
	}

	err = s.organizationRepo.Delete(ctx, id)
	if errors.Is(err, dbmodels.ErrBadRequest) {
		return entities.ErrBadRequest
	}

	return err
}

func (s *organizationService) AssignManager(ctx context.Context, organizationID uuid.UUID, userID string) error {
	_, err := s.organizationRepo.GetById(ctx, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ErrNotFound
		}

		return err
	}

	dbUser, err := s.adminRepo.GetById(ctx, userID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ErrNotFound
		}

		return err
	}

	user := conversions.AdminUserFromDb(dbUser)
	if !app.IsAdmin(user) {
		return entities.ErrBadRequest
	}

	return s.organizationRepo.AssignManager(ctx, organizationID, userID)
}

func (s *organizationService) RemoveManager(ctx context.Context, organizationID uuid.UUID, userID string) error {
	_, err := s.organizationRepo.GetById(ctx, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ErrNotFound
		}

		return err
	}

	dbUser, err := s.adminRepo.GetById(ctx, userID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ErrNotFound
		}

		return err
	}

	user := conversions.AdminUserFromDb(dbUser)
	if !app.IsAdmin(user) {
		return entities.ErrBadRequest
	}

	err = s.organizationRepo.RemoveManager(ctx, organizationID, userID)
	if errors.Is(err, dbmodels.ErrNotFound) {
		return entities.ErrNotFound
	}

	return err
}

func (s *organizationService) GetDefaultGroupByOrganizationId(ctx context.Context, id uuid.UUID) (entities.ServerGroup, error) {
	group, err := s.organizationRepo.GetDefaultGroupByOrganizationId(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entities.ErrNotFound
		}

		return entities.ServerGroup{}, err
	}

	return conversions.ServerGroupFromDB(group), nil
}

func (s *organizationService) GetAdminUsersByOrganizationID(ctx context.Context, id uuid.UUID) ([]entities.AdminUser, error) {
	users, err := s.organizationRepo.GetAdminUsersByOrganizationID(ctx, id)
	if err != nil {
		return nil, err
	}

	return conversions.AdminUsersFromDb(users), nil
}

func (s *organizationService) GetAnyByID(ctx context.Context, id uuid.UUID) (entities.Organization, error) {
	org, err := s.organizationRepo.GetAnyByID(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entities.ErrNotFound
		}
		return entities.Organization{}, err
	}

	return conversions.OrganizationFromDB(org), nil
}
