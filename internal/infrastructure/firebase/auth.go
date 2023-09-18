package firebase

import (
	"context"
	"errors"
	"strings"
	"washbonus/internal/app"
	"washbonus/internal/entities"

	opErrors "github.com/go-openapi/errors"
)

var ErrUnauthorized = opErrors.New(401, "unauthorized")

func (svc *FirebaseService) BonusAuth(bearer string) (*app.Auth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	jwtToken := strings.TrimSpace(strings.Replace(bearer, "Bearer", "", 1))
	if jwtToken == "" {
		return nil, ErrUnauthorized
	}

	token, err := svc.auth.VerifyIDToken(context.Background(), jwtToken)
	if err != nil {
		return nil, ErrUnauthorized
	}

	fbUser, err := svc.auth.GetUser(ctx, token.UID)
	if err != nil {
		return nil, ErrUnauthorized
	}

	user, err := svc.userSvc.GetById(ctx, fbUser.UID)
	if err != nil {
		if errors.Is(err, entities.ErrNotFound) {
			userCreation := entities.UserCreation{
				ID:    fbUser.UID,
				Email: fbUser.Email,
				Name:  fbUser.DisplayName,
			}

			user, err = svc.userSvc.Create(ctx, userCreation)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if user.Email == nil {
		err := svc.userSvc.UpdateUser(ctx, entities.UserUpdate{
			ID:    fbUser.UID,
			Email: fbUser.Email,
			Name:  fbUser.DisplayName,
		})
		if err != nil {
			return nil, err
		}
	}

	return &app.Auth{
		User:         user,
		Disabled:     fbUser.Disabled,
		UserMetadata: (*app.AuthUserMeta)(fbUser.UserMetadata),
	}, nil
}

func (svc *FirebaseService) AdminAuth(bearer string) (*app.AdminAuth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	jwtToken := strings.TrimSpace(strings.Replace(bearer, "Bearer", "", 1))
	if jwtToken == "" {
		return nil, ErrUnauthorized
	}

	token, err := svc.auth.VerifyIDToken(context.Background(), jwtToken)
	if err != nil {
		return nil, ErrUnauthorized
	}

	fbUser, err := svc.auth.GetUser(ctx, token.UID)
	if err != nil {
		return nil, ErrUnauthorized
	}

	user, err := svc.adminSvc.GetById(ctx, fbUser.UID)
	if err != nil {
		if errors.Is(err, entities.ErrNotFound) {
			return nil, ErrUnauthorized
		} else {
			return nil, err
		}
	}

	if user.Email == nil {
		err := svc.adminSvc.Update(ctx, entities.AdminUserUpdate{
			ID:    fbUser.UID,
			Email: &fbUser.Email,
			Name:  &fbUser.DisplayName,
		})

		if err != nil {
			return nil, err
		}
	}

	return &app.AdminAuth{
		User:         user,
		Disabled:     fbUser.Disabled,
		UserMetadata: (*app.AuthUserMeta)(fbUser.UserMetadata),
	}, nil
}
