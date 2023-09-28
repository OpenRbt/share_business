package user

import (
	"context"
	"errors"
	"washbonus/internal/conversions"
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
)

func (s *userService) GetById(ctx context.Context, userID string) (entities.User, error) {
	userFromDB, err := s.userRepo.GetById(ctx, userID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			err = entities.ErrNotFound
		}
		return entities.User{}, err
	}

	return conversions.UserFromDb(userFromDB), nil
}

func (s *userService) Create(ctx context.Context, userCreation entities.UserCreation) (entities.User, error) {
	dbUser := conversions.UserCreationToDB(userCreation)

	user, err := s.userRepo.Create(ctx, dbUser)
	if err != nil {
		return entities.User{}, err
	}

	return conversions.UserFromDb(user), nil
}

func (s *userService) UpdateUser(ctx context.Context, userModel entities.UserUpdate) error {
	_, err := s.userRepo.GetById(ctx, userModel.ID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.ErrNotFound
		}
		return err
	}

	return s.userRepo.UpdateUser(ctx, conversions.UserUpdateToDB(userModel))
}
