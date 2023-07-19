package rest

import (
	"errors"
	"washBonus/internal/app"
	"washBonus/internal/conversions"
	"washBonus/internal/entity"
	"washBonus/openapi/restapi/operations"
	"washBonus/openapi/restapi/operations/users"
)

func (svc *service) initUserHandlers(api *operations.WashBonusAPI) {
	api.UsersGetUserByIDHandler = users.GetUserByIDHandlerFunc(svc.getUserByID)
	api.UsersGetCurrentUserHandler = users.GetCurrentUserHandlerFunc(svc.getCurrentUser)
	api.UsersUpdateUserHandler = users.UpdateUserHandlerFunc(svc.updateUser)
}

func (svc *service) getUserByID(params users.GetUserByIDParams, auth *app.Auth) users.GetUserByIDResponder {
	res, err := svc.userCtrl.Get(params.HTTPRequest.Context(), auth.UID, params.ID)

	payload := conversions.UserToRest(res)

	switch {
	case err == nil:
		return users.NewGetUserByIDOK().WithPayload(&payload)
	case errors.Is(err, entity.ErrNotFound):
		return users.NewGetUserByIDNotFound()
	case errors.Is(err, entity.ErrAccessDenied):
		return users.NewGetUserByIDForbidden()
	default:
		return users.NewGetUserByIDInternalServerError()
	}
}

func (svc *service) getCurrentUser(params users.GetCurrentUserParams, auth *app.Auth) users.GetCurrentUserResponder {
	res, err := svc.userCtrl.Get(params.HTTPRequest.Context(), auth.UID, auth.UID)

	payload := conversions.UserToRest(res)

	switch {
	case err == nil:
		return users.NewGetCurrentUserOK().WithPayload(&payload)
	case errors.Is(err, entity.ErrNotFound):
		return users.NewGetCurrentUserNotFound()
	case errors.Is(err, entity.ErrAccessDenied):
		return users.NewGetCurrentUserForbidden()
	default:
		return users.NewGetCurrentUserInternalServerError()
	}
}

func (svc *service) updateUser(params users.UpdateUserParams, auth *app.Auth) users.UpdateUserResponder {
	err := svc.userCtrl.UpdateUserRole(params.HTTPRequest.Context(), auth.UID, entity.UpdateUser{ID: params.ID, Role: entity.Role(params.Update.Role)})

	switch {
	case err == nil:
		return users.NewUpdateUserNoContent()
	case errors.Is(err, entity.ErrNotFound):
		return users.NewUpdateUserNotFound()
	case errors.Is(err, entity.ErrAccessDenied):
		return users.NewUpdateUserForbidden()
	default:
		return users.NewUpdateUserInternalServerError()
	}
}
