package adminuser

import (
	"context"
	"errors"
	"fmt"
	"washbonus/internal/conversions"
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"

	uuid "github.com/satori/go.uuid"
)

func (s *adminService) Get(ctx context.Context, filter entities.AdminUserFilter) ([]entities.AdminUser, error) {
	usersFromDB, err := s.adminRepo.Get(ctx, conversions.AdminUserFilterToDB(filter))
	if err != nil {
		return nil, err
	}

	return conversions.AdminUsersFromDb(usersFromDB), nil
}

func (s *adminService) GetAll(ctx context.Context, pagination entities.Pagination) ([]entities.AdminUser, error) {
	usersFromDB, err := s.adminRepo.GetAll(ctx, conversions.PaginationToDB(pagination))
	if err != nil {
		return nil, err
	}

	return conversions.AdminUsersFromDb(usersFromDB), nil
}

func (s *adminService) GetById(ctx context.Context, userID string) (entities.AdminUser, error) {
	userFromDB, err := s.adminRepo.GetById(ctx, userID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entities.ErrNotFound
		}
		return entities.AdminUser{}, err
	}

	return conversions.AdminUserFromDb(userFromDB), nil
}

func (s *adminService) Create(ctx context.Context, userCreation entities.AdminUserCreation) (entities.AdminUser, error) {
	dbUser := conversions.AdminUserCreationToDB(userCreation)

	user, err := s.adminRepo.Create(ctx, dbUser)
	if err != nil {
		return entities.AdminUser{}, err
	}

	return conversions.AdminUserFromDb(user), nil
}

func (s *adminService) UpdateRole(ctx context.Context, userRole entities.AdminUserRoleUpdate) error {
	_, err := s.adminRepo.GetById(ctx, userRole.ID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ErrNotFound
		}
		return err
	}

	return s.adminRepo.UpdateRole(ctx, conversions.AdminUserRoleUpdateToDB(userRole))
}

func (s *adminService) Update(ctx context.Context, userModel entities.AdminUserUpdate) error {
	_, err := s.adminRepo.GetById(ctx, userModel.ID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ErrNotFound
		}
		return err
	}

	return s.adminRepo.Update(ctx, conversions.AdminUserUpdateToDB(userModel))
}

func (s *adminService) GetApplications(ctx context.Context, filter entities.AdminApplicationFilter) ([]entities.AdminApplication, error) {
	appsFromDB, err := s.adminRepo.GetApplications(ctx, conversions.AdminApplicationFilterToDB(filter))
	if err != nil {
		return nil, err
	}

	return conversions.AdminApplicationsFromDb(appsFromDB), nil
}

func (s *adminService) Block(ctx context.Context, id string) error {
	_, err := s.adminRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ErrNotFound
		}
		return err
	}

	return s.adminRepo.Block(ctx, id)
}

func (s *adminService) CreateApplication(ctx context.Context, ent entities.AdminApplicationCreation) (entities.AdminApplication, error) {
	dbApp := conversions.AdminApplicationCreationToDB(ent)

	app, err := s.adminRepo.CreateApplication(ctx, dbApp)
	if err != nil {
		if errors.Is(err, dbmodels.ErrAlreadyExists) {
			err = fmt.Errorf(err.Error()+": %w", entities.ErrBadRequest)
		}

		return entities.AdminApplication{}, err
	}

	return conversions.AdminApplicationFromDb(app), nil
}

func (s *adminService) ReviewApplication(ctx context.Context, id uuid.UUID, ent entities.AdminApplicationReview) error {
	dbReview := conversions.AdminApplicationReviewToDB(ent)

	err := s.adminRepo.ReviewApplication(ctx, id, dbReview)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ErrNotFound
		} else if errors.Is(err, dbmodels.ErrBadRequest) {
			return entities.ErrBadRequest
		}

		return err
	}

	return nil
}

func (s *adminService) GetApplicationByID(ctx context.Context, id uuid.UUID) (entities.AdminApplication, error) {
	app, err := s.adminRepo.GetApplicationByID(ctx, id)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entities.ErrNotFound
		}

		return entities.AdminApplication{}, err
	}

	return conversions.AdminApplicationFromDb(app), nil
}
