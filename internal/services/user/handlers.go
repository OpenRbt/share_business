package user

import (
	"context"
	"errors"
	"washBonus/internal/conversions"
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
)

func (s *userService) Get(ctx context.Context, pagination entity.Pagination) ([]entity.User, error) {
	usersFromDB, err := s.userRepo.Get(ctx, conversions.PaginationToDB(pagination))
	if err != nil {
		return nil, err
	}

	return conversions.UsersFromDb(usersFromDB), nil
}

func (s *userService) GetById(ctx context.Context, userID string) (entity.User, error) {
	userFromDB, err := s.userRepo.GetById(ctx, userID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entity.ErrNotFound
		}
		return entity.User{}, err
	}

	return conversions.UserFromDb(userFromDB), nil
}

func (s *userService) Create(ctx context.Context, userCreation entity.UserCreation) (entity.User, error) {
	dbUser := conversions.UserCreationToDB(userCreation)

	user, err := s.userRepo.Create(ctx, dbUser)
	if err != nil {
		return entity.User{}, err
	}

	return conversions.UserFromDb(user), nil
}

func (s *userService) UpdateUserRole(ctx context.Context, userUpdate entity.UserUpdate) error {
	_, err := s.userRepo.GetById(ctx, userUpdate.ID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.ErrNotFound
		}
		return err
	}

	return s.userRepo.UpdateUserRole(ctx, conversions.UserUpdateToDB(userUpdate))
}
