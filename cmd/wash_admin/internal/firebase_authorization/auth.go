package firebase_authorization

import (
	"context"
	"strings"
	"wash_admin/internal/app"
)

func (svc *FirebaseService) Auth(bearer string) (*app.Auth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	idToken := strings.TrimSpace(strings.Replace(bearer, "Bearer", "", 1))

	if idToken == "" {
		return nil, app.ErrAccessDenied
	}

	token, err := svc.auth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, app.ErrAccessDenied
	}

	user, err := svc.auth.GetUser(ctx, token.UID)
	if err != nil {
		return nil, app.ErrAccessDenied
	}

	return &app.Auth{
		UID:          user.UID,
		Disabled:     user.Disabled,
		UserMetadata: (*app.AuthUserMeta)(user.UserMetadata),
	}, nil
}
