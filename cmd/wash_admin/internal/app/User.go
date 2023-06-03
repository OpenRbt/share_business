package app

import (
	"context"
)

type User struct {
	ID   string
	Role Role
}

type UpdateUser struct {
	ID   string
	Role Role
}

func (svc *WashServerSvc) UpdateUserRole(ctx context.Context, auth *Auth, userUpdate UpdateUser) error {
	user, err := svc.repo.GetOrCreateUserIfNotExists(ctx, auth.UID)
	if err != nil {
		return err
	}

	switch user.Role {
	case AdminRole:
		_, err = svc.repo.GetUser(ctx, userUpdate.ID)
		if err != nil {
			return err
		}

		err = svc.repo.UpdateUserRole(ctx, userUpdate)
		if err != nil {
			return err
		}

		return nil
	default:
		return ErrAccessDenied
	}
}
