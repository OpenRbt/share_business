package app

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

type WashUser struct {
	ID       uuid.UUID
	Identity string
	Role     Role
}

type UpdateUser struct {
	ID       uuid.UUID
	Identity string
	Role     Role
}

func (svc *WashServerSvc) UpdateUser(ctx context.Context, auth *Auth, userUpdate UpdateUser) error {
	user, err := svc.repo.GetOrCreateUserIfNotExists(ctx, auth.UID)

	if err != nil {
		return err
	}

	switch user.Role {
	case AdminRole:

		_, err := svc.repo.GetOrCreateUserIfNotExists(ctx, userUpdate.Identity)

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
