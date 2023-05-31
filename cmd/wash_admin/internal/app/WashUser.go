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

func (svc *WashServerSvc) UpdateUserRole(ctx context.Context, auth *Auth, userUpdate UpdateUser) error {
	user, err := svc.repo.GetOrCreateUserIfNotExists(ctx, auth.UID)
	if err != nil {
		return err
	}

	if !checkAccess(user, roleAdmin) {
		return ErrAccessDenied
	}

	_, err = svc.repo.GetWashUser(ctx, userUpdate.Identity)
	if err != nil {
		return err
	}

	return svc.repo.UpdateUserRole(ctx, userUpdate)
}
