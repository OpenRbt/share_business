package app

import (
	"context"
	"washBonus/internal/entity"

	"github.com/shopspring/decimal"
)

type (
	UserController interface {
		Get(ctx context.Context, authorizedUserID string, userID string) (entity.User, error)
		UpdateUserRole(ctx context.Context, authorizedUserID string, userUpdate entity.UpdateUser) error
	}

	UserService interface {
		Create(ctx context.Context, userID string) (user entity.User, err error)
		Get(ctx context.Context, userID string) (user entity.User, err error)
		GetOrCreate(ctx context.Context, userID string) (user entity.User, err error)
		AddBonuses(ctx context.Context, amount decimal.Decimal, userID string) (err error)
		UpdateUserRole(ctx context.Context, authorizedUserID string, userUpdate entity.UpdateUser) error
	}

	UserRepo interface {
		GetByID(ctx context.Context, userID string) (user entity.User, err error)
		Create(ctx context.Context, userID string) (user entity.User, err error)
		AddBonuses(ctx context.Context, amount decimal.Decimal, userID string) (err error)
		GetBalance(ctx context.Context, userID string) (balance decimal.Decimal, err error)
		UpdateUserRole(ctx context.Context, userUpdate entity.UpdateUser) error
	}
)
