package firebase

import (
	"context"
	"strings"
	"wash_bonus/internal/app"
	"wash_bonus/internal/entity"
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

	return &app.Auth{
		UID:          user.UID,
		Disabled:     user.Disabled,
		UserMetadata: (*app.AuthUserMeta)(user.UserMetadata),
	}, nil
}
