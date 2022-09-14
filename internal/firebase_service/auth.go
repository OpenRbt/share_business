package firebase_service

import (
	"context"
	"net/http"
	"strings"
	"wash-bonus/internal/app"
)

func (svc *Service) AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			authToken := r.Header.Get("Authorization")

			idToken := strings.TrimSpace(strings.Replace(authToken, "Bearer", "", 1))

			if idToken == "" {
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}

			_, err := svc.auth.VerifyIDToken(context.Background(), idToken)

			if err != nil {
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}

			handler.ServeHTTP(rw, r)
		})
}

func (svc *Service) GetFirebaseProfile(bearer string) (interface{}, error) {
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
