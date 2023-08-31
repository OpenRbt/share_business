package firebase

import (
	"context"
	"errors"
	"strings"
	"washBonus/internal/app"
	"washBonus/internal/entity"
)

func (svc *FirebaseService) Auth(bearer string) (*app.Auth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	idToken := strings.TrimSpace(strings.Replace(bearer, "Bearer", "", 1))

	if idToken == "" {
		return nil, entity.ErrAccessDenied
	}

	token, err := svc.auth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, entity.ErrAccessDenied
	}

	fbUser, err := svc.auth.GetUser(ctx, token.UID)
	if err != nil {
		return nil, entity.ErrAccessDenied
	}

	user, err := svc.userSvc.GetById(ctx, fbUser.UID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			userCreation := entity.UserCreation{
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

	return &app.Auth{
		User:         user,
		Disabled:     fbUser.Disabled,
		UserMetadata: (*app.AuthUserMeta)(fbUser.UserMetadata),
	}, nil
}
