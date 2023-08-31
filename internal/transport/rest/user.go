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
	api.UsersUpdateUserRoleHandler = users.UpdateUserRoleHandlerFunc(svc.updateUserRole)
	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(svc.getUsers)
}

func (svc *service) getUsers(params users.GetUsersParams, auth *app.Auth) users.GetUsersResponder {
	pagination := conversions.PaginationFromRest(*params.Limit, *params.Offset)
	res, err := svc.userCtrl.Get(params.HTTPRequest.Context(), *auth, pagination)

	switch {
	case err == nil:
		payload := conversions.UsersToRest(res)
		return users.NewGetUsersOK().WithPayload(payload)
	case errors.Is(err, entity.ErrAccessDenied):
		return users.NewGetUsersForbidden()
	default:
		svc.l.Errorln("Get users:", err)
		return users.NewGetUsersInternalServerError()
	}
}

func (svc *service) getUserByID(params users.GetUserByIDParams, auth *app.Auth) users.GetUserByIDResponder {
	res, err := svc.userCtrl.GetById(params.HTTPRequest.Context(), *auth, params.UserID)

	switch {
	case err == nil:
		payload := conversions.UserToRest(res)
		return users.NewGetUserByIDOK().WithPayload(&payload)
	case errors.Is(err, entity.ErrNotFound):
		return users.NewGetUserByIDNotFound()
	case errors.Is(err, entity.ErrAccessDenied):
		return users.NewGetUserByIDForbidden()
	default:
		svc.l.Errorln("Get user by id:", err)
		return users.NewGetUserByIDInternalServerError()
	}
}

func (svc *service) getCurrentUser(params users.GetCurrentUserParams, auth *app.Auth) users.GetCurrentUserResponder {
	res, err := svc.userCtrl.GetById(params.HTTPRequest.Context(), *auth, auth.User.ID)

	switch {
	case err == nil:
		payload := conversions.UserToRest(res)
		return users.NewGetCurrentUserOK().WithPayload(&payload)
	case errors.Is(err, entity.ErrNotFound):
		return users.NewGetCurrentUserNotFound()
	case errors.Is(err, entity.ErrAccessDenied):
		return users.NewGetCurrentUserForbidden()
	default:
		svc.l.Errorln("Get current user:", err)
		return users.NewGetCurrentUserInternalServerError()
	}
}

func (svc *service) updateUserRole(params users.UpdateUserRoleParams, auth *app.Auth) users.UpdateUserRoleResponder {
	err := svc.userCtrl.UpdateUserRole(params.HTTPRequest.Context(), *auth, entity.UserUpdateRole{ID: params.UserID, Role: entity.Role(params.Update.Role)})

	switch {
	case err == nil:
		return users.NewUpdateUserRoleNoContent()
	case errors.Is(err, entity.ErrNotFound):
		return users.NewUpdateUserRoleNotFound()
	case errors.Is(err, entity.ErrAccessDenied):
		return users.NewUpdateUserRoleForbidden()
	default:
		svc.l.Errorln("Update user role:", err)
		return users.NewUpdateUserRoleInternalServerError()
	}
}
