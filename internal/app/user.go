package app

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
)

type (
	UserController interface {
		GetById(ctx Ctx, auth Auth, userID string) (entity.User, error)
		Get(ctx Ctx, auth Auth, pagination entity.Pagination) ([]entity.User, error)
		UpdateUserRole(ctx Ctx, auth Auth, userUpdate entity.UserUpdateRole) error
	}

	UserService interface {
		GetById(ctx Ctx, userID string) (entity.User, error)
		Get(ctx Ctx, pagination entity.Pagination) ([]entity.User, error)
		Create(ctx Ctx, userCreation entity.UserCreation) (entity.User, error)
		UpdateUserRole(ctx Ctx, userRole entity.UserUpdateRole) error
		UpdateUser(ctx Ctx, userModel entity.UserUpdate) error
	}

	UserRepo interface {
		GetById(ctx Ctx, userID string) (dbmodels.User, error)
		Get(ctx Ctx, pagination dbmodels.Pagination) ([]dbmodels.User, error)
		Create(ctx Ctx, userCreation dbmodels.UserCreation) (dbmodels.User, error)
		UpdateUserRole(ctx Ctx, userUpdate dbmodels.UserUpdateRole) error
		UpdateUser(ctx Ctx, userModel dbmodels.UserUpdate) error
	}
)

func IsAdmin(user entity.User) bool {
	return user.Role == entity.AdminRole
}

func IsEngineer(user entity.User) bool {
	return user.Role == entity.EngineerRole
}

func IsUser(user entity.User) bool {
	return user.Role == entity.UserRole
}
