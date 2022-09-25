package firebase_auth

import (
	"context"
	"errors"
	"strings"
	"wash-bonus/internal/app"
)

func (svc *FirebaseService) Auth(bearer string) (interface{}, error) {
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

	return user, nil
}

func (svc *FirebaseService) VerifyClaims(user FirebaseProfile, requiredClaims ...string) error {
	userClaims := user.CustomClaims
	for _, claim := range requiredClaims {
		if _, ok := userClaims[claim]; !ok {
			return errors.New("missing claims")
		}
	}
	return nil
}
