package api

import (
	"errors"
	"github.com/go-openapi/swag"
	"wash-bonus/internal/def"
	"wash-bonus/internal/dto"
	"wash-bonus/internal/firebase_auth"
	"wash-bonus/internal/transport/rest/restapi/models"
	"wash-bonus/internal/transport/rest/restapi/restapi/operations"
	user2 "wash-bonus/internal/transport/rest/restapi/restapi/operations/user"

	"wash-bonus/internal/app"

	"github.com/go-openapi/runtime/middleware"
)

func setUserHandlers(api *operations.WashBonusAPI, svc *service) {
	api.UserGetUserHandler = user2.GetUserHandlerFunc(svc.GetUser)
	api.UserGetCurrentUserHandler = user2.GetCurrentUserHandlerFunc(svc.GetCurrentUser)
	api.UserAddUserHandler = user2.AddUserHandlerFunc(svc.AddUser)
	api.UserEditUserHandler = user2.EditUserHandlerFunc(svc.EditUser)
	api.UserDeleteUserHandler = user2.DeleteUserHandlerFunc(svc.DeleteUser)
	api.UserListUserHandler = user2.ListUserHandlerFunc(svc.ListUser)
}

func (svc *service) GetCurrentUser(params user2.GetCurrentUserParams, profile interface{}) middleware.Responder {
	prof := profile.(*firebase_auth.FirebaseProfile)

	u, err := svc.userSvc.GetByIdentityID(dto.ToAppIdentityProfile(*prof))

	switch {
	default:
		log.PrintErr("Get server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return user2.NewGetUserDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("Get client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return user2.NewGetUserDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("Get client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return user2.NewGetUserDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("Get ok")
		resp := dto.UserToRest(*u)
		return user2.NewGetUserOK().WithPayload(&resp)
	}
}

func (svc *service) GetUser(params user2.GetUserParams, profile interface{}) middleware.Responder {
	prof := profile.(*firebase_auth.FirebaseProfile)

	u, err := svc.userSvc.Get(dto.ToAppIdentityProfile(*prof), params.ID)

	switch {
	default:
		log.PrintErr("Get server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return user2.NewGetUserDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("Get client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return user2.NewGetUserDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("Get client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return user2.NewGetUserDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("Get ok", "id", params.ID)
		resp := dto.UserToRest(*u)
		return user2.NewGetUserOK().WithPayload(&resp)
	}
}
func (svc *service) AddUser(params user2.AddUserParams, profile interface{}) middleware.Responder {
	prof := profile.(*firebase_auth.FirebaseProfile)

	err := svc.userSvc.Add(dto.ToAppIdentityProfile(*prof), dto.UserFromRestAdd(*params.Body))

	switch {
	default:
		log.PrintErr("Add server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return user2.NewAddUserDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("Add client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return user2.NewAddUserDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("Add ok")
		return user2.NewAddUserOK()
	}
}
func (svc *service) EditUser(params user2.EditUserParams, profile interface{}) middleware.Responder {
	prof := profile.(*firebase_auth.FirebaseProfile)

	err := svc.userSvc.Edit(dto.ToAppIdentityProfile(*prof), params.ID, dto.UserFromRestUpdate(*params.Body))

	switch {
	default:
		log.PrintErr("Edit server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return user2.NewEditUserDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("Edit client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return user2.NewEditUserDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("Edit client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return user2.NewEditUserDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("Edit ok")
		return user2.NewEditUserOK()
	}
}
func (svc *service) DeleteUser(params user2.DeleteUserParams, profile interface{}) middleware.Responder {
	prof := profile.(*firebase_auth.FirebaseProfile)

	err := svc.userSvc.Delete(dto.ToAppIdentityProfile(*prof), params.ID)

	switch {
	default:
		log.PrintErr("Delete server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return user2.NewDeleteUserDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("Delete client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return user2.NewDeleteUserDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("Delete client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return user2.NewDeleteUserDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("Delete ok", "id", params.ID)
		return user2.NewDeleteUserNoContent()
	}
}
func (svc *service) ListUser(params user2.ListUserParams, profile interface{}) middleware.Responder {
	prof := profile.(*firebase_auth.FirebaseProfile)

	c, warnings, err := svc.userSvc.List(dto.ToAppIdentityProfile(*prof), dto.ListFilterFromRest(params.Body))

	switch {
	default:
		log.PrintErr("List server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return user2.NewListUserDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("List client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return user2.NewListUserDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("List ok")
		return user2.NewListUserOK().WithPayload(&user2.ListUserOKBody{
			Items:    dto.UsersToRest(c),
			Warnings: warnings,
		})
	}
}
