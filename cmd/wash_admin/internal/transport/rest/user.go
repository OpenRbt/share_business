package rest

import (
	"errors"
	"wash_admin/internal/app"
	"wash_admin/openapi/restapi/operations"
	"wash_admin/openapi/restapi/operations/users"
)

func (svc *service) initUserHandlers(api *operations.WashAdminAPI) {
	api.UsersGetUserHandler = users.GetUserHandlerFunc(svc.getUser)
	api.UsersUpdateUserHandler = users.UpdateUserHandlerFunc(svc.updateUser)
}

func (svc *service) getUser(params users.GetUserParams, auth *app.Auth) users.GetUserResponder {
	return nil
}
func (svc *service) updateUser(params users.UpdateUserParams, auth *app.Auth) users.UpdateUserResponder {
	err := svc.washServers.UpdateUserRole(params.HTTPRequest.Context(), auth, app.UpdateUser{ID: params.ID, Role: app.Role(params.Update.Role)})

	switch {
	case err == nil:
		return users.NewUpdateUserNoContent()
	case errors.Is(err, app.ErrAccessDenied):
		return users.NewUpdateUserForbidden()
	default:
		return users.NewUpdateUserInternalServerError()
	}
}
