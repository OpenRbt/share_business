package app

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
)

type (
	UserController interface {
		GetById(ctx Ctx, auth Auth, userID string) (entities.User, error)
	}

	UserService interface {
		GetById(ctx Ctx, userID string) (entities.User, error)
		Create(ctx Ctx, userCreation entities.UserCreation) (entities.User, error)
		UpdateUser(ctx Ctx, userModel entities.UserUpdate) error
	}

	UserRepo interface {
		GetById(ctx Ctx, userID string) (dbmodels.User, error)
		Create(ctx Ctx, userCreation dbmodels.UserCreation) (dbmodels.User, error)
		UpdateUser(ctx Ctx, userModel dbmodels.UserUpdate) error
	}
)
