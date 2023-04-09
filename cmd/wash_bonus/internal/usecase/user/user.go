package user

import (
	"context"
	"errors"
	"wash_bonus/internal/entity"
)

func (u useCase) Get(ctx context.Context, userID string) (user entity.User, err error) {
	user, err = u.UserSvc.Get(ctx, userID)
	if err != nil && errors.Is(err, entity.ErrNotFound) {
		user, err = u.UserSvc.Create(ctx, userID)
	}
	return
}
