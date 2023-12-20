package firebase

import (
	"context"
	"errors"
	"strings"
	"washbonus/internal/app"
	"washbonus/internal/conversions"
	"washbonus/internal/entities"
	rabbitEntities "washbonus/internal/infrastructure/rabbit/entities"

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

	user, err := svc.userSvc.GetById(ctx, token.UID)
	if errors.Is(err, entities.ErrNotFound) {
		fbUser, err := svc.auth.GetUser(ctx, token.UID)
		if err != nil {
			return nil, ErrUnauthorized
		}

		userCreation := entities.UserCreation{
			ID:    fbUser.UID,
			Email: fbUser.Email,
			Name:  fbUser.DisplayName,
		}

		user, err = svc.userSvc.Create(ctx, userCreation)
		if err != nil {
			return nil, ErrUnauthorized
		}
	} else if err != nil {
		return nil, ErrUnauthorized
	}

	return &app.Auth{
		User: user,
	}, nil
}

func (svc *FirebaseService) AdminAuth(bearer string) (*app.AdminAuth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	jwtToken := strings.TrimSpace(strings.Replace(bearer, "Bearer", "", 1))
	if jwtToken == "" {
		return nil, ErrUnauthorized
	}

	svc.adminCache.RLock()
	cachedAuth, found := svc.adminCache.Cache[jwtToken]
	svc.adminCache.RUnlock()
	if found {
		return cachedAuth, nil
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

	if app.IsAdminHasNoAccess(user) {
		return nil, ErrUnauthorized
	}

	if user.Email == "" || user.Name != fbUser.DisplayName {
		err := svc.adminSvc.Update(ctx, entities.AdminUserUpdate{
			ID:    fbUser.UID,
			Email: &fbUser.Email,
			Name:  &fbUser.DisplayName,
		})

		if err != nil {
			return nil, err
		}

		rabbitUser := conversions.AdminUserToRabbit(user)
		err = svc.rabbitSvc.SendMessage(rabbitUser, rabbitEntities.AdminsExchange, rabbitEntities.WashBonusRoutingKey, rabbitEntities.AdminUserMessageType)
		if err != nil {
			return nil, err
		}
	}

	authData := &app.AdminAuth{
		User:         user,
		Disabled:     fbUser.Disabled,
		UserMetadata: (*app.AuthUserMeta)(fbUser.UserMetadata),
	}

	svc.adminCache.Lock()
	svc.adminCache.Cache[jwtToken] = authData
	svc.adminCache.Unlock()

	return authData, nil
}
