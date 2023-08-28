package app

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
)

type (
	UserController interface {
		GetById(ctx Ctx, authUser entity.User, userID string) (entity.User, error)
		Get(ctx Ctx, authUser entity.User, pagination entity.Pagination) ([]entity.User, error)
		UpdateUserRole(ctx Ctx, authUser entity.User, userUpdate entity.UserUpdate) error
	}

	UserService interface {
		Create(ctx Ctx, userID string) (user entity.User, err error)
		GetById(ctx Ctx, userID string) (user entity.User, err error)
		Get(ctx Ctx, pagination entity.Pagination) ([]entity.User, error)
		GetOrCreate(ctx Ctx, userID string) (user entity.User, err error)
		UpdateUserRole(ctx Ctx, userUpdate entity.UserUpdate) error
	}

	UserRepo interface {
		GetById(ctx Ctx, userID string) (user dbmodels.User, err error)
		Get(ctx Ctx, pagination dbmodels.Pagination) ([]dbmodels.User, error)
		Create(ctx Ctx, userID string) (user dbmodels.User, err error)
		UpdateUserRole(ctx Ctx, userUpdate dbmodels.UserUpdate) error
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
