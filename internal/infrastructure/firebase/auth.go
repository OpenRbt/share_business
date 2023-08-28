package firebase

import (
	"context"
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

	user, err := svc.auth.GetUser(ctx, token.UID)
	if err != nil {
		return nil, entity.ErrAccessDenied
	}

	dbUser, err := svc.userSvc.GetOrCreate(ctx, user.UID)
	if err != nil {
		return nil, err
	}

	return &app.Auth{
		UID:          user.UID,
		Disabled:     user.Disabled,
		User:         dbUser,
		UserMetadata: (*app.AuthUserMeta)(user.UserMetadata),
	}, nil
}
