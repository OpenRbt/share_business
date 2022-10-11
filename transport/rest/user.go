package api

import (
	"errors"
	"fmt"
	"wash-bonus/internal/def"
	"wash-bonus/internal/dto"
	"wash-bonus/internal/firebase_auth"
	"wash-bonus/transport/rest/restapi/restapi/operations"

	"github.com/go-openapi/swag"

	"wash-bonus/transport/rest/restapi/models"

	"wash-bonus/transport/rest/restapi/restapi/operations/user"

	"wash-bonus/internal/app"

	"github.com/go-openapi/runtime/middleware"
)

func setUserHandlers(api *operations.WashBonusAPI, svc *service) {
	api.UserGetUserHandler = user.GetUserHandlerFunc(svc.GetUser)
	api.UserGetCurrentUserHandler = user.GetCurrentUserHandlerFunc(svc.GetCurrentUser)
	api.UserAddUserHandler = user.AddUserHandlerFunc(svc.AddUser)
	api.UserEditUserHandler = user.EditUserHandlerFunc(svc.EditUser)
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(svc.DeleteUser)
	api.UserListUserHandler = user.ListUserHandlerFunc(svc.ListUser)
}

func (svc *service) GetCurrentUser(params user.GetCurrentUserParams, profile interface{}) middleware.Responder {
	prof := profile.(*firebase_auth.FirebaseProfile)

	u, err := svc.userSvc.GetByIdentityID(dto.ToAppIdentityProfile(*prof))

	switch {
	default:
		log.PrintErr("Get server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return user.NewGetUserDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("Get client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return user.NewGetUserDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("Get client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return user.NewGetUserDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("Get ok")
		resp := dto.UserToRest(*u)
		return user.NewGetUserOK().WithPayload(&resp)
	}
}

func (svc *service) GetUser(params user.GetUserParams, profile interface{}) middleware.Responder {
	prof := profile.(*firebase_auth.FirebaseProfile)

	u, err := svc.userSvc.Get(dto.ToAppIdentityProfile(*prof), params.ID)

	switch {
	default:
		log.PrintErr("Get server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return user.NewGetUserDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("Get client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return user.NewGetUserDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("Get client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return user.NewGetUserDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("Get ok", "id", params.ID)
		resp := dto.UserToRest(*u)
		return user.NewGetUserOK().WithPayload(&resp)
	}
}
func (svc *service) AddUser(params user.AddUserParams, profile interface{}) middleware.Responder {
	fmt.Println("\nAddUser start: ", profile)
	prof := profile.(*firebase_auth.FirebaseProfile)
	fmt.Println("\nAddUser prof: ", prof.UID)

	err := svc.userSvc.Add(dto.ToAppIdentityProfile(*prof), dto.UserFromRestAdd(*params.Body))
	fmt.Println("\nAddUser com: ", err)

	switch {
	default:
		log.PrintErr("Add server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return user.NewAddUserDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("Add client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return user.NewAddUserDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("Add ok")
		return user.NewAddUserOK()
	}
}
func (svc *service) EditUser(params user.EditUserParams, profile interface{}) middleware.Responder {
	prof := profile.(*firebase_auth.FirebaseProfile)

	err := svc.userSvc.Edit(dto.ToAppIdentityProfile(*prof), params.ID, dto.UserFromRestUpdate(*params.Body))

	switch {
	default:
		log.PrintErr("Edit server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return user.NewEditUserDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("Edit client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return user.NewEditUserDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("Edit client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return user.NewEditUserDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("Edit ok")
		return user.NewEditUserOK()
	}
}
func (svc *service) DeleteUser(params user.DeleteUserParams, profile interface{}) middleware.Responder {
	prof := profile.(*firebase_auth.FirebaseProfile)

	err := svc.userSvc.Delete(dto.ToAppIdentityProfile(*prof), params.ID)

	switch {
	default:
		log.PrintErr("Delete server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return user.NewDeleteUserDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("Delete client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return user.NewDeleteUserDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("Delete client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return user.NewDeleteUserDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("Delete ok", "id", params.ID)
		return user.NewDeleteUserNoContent()
	}
}
func (svc *service) ListUser(params user.ListUserParams, profile interface{}) middleware.Responder {
	prof := profile.(*firebase_auth.FirebaseProfile)

	c, warnings, err := svc.userSvc.List(dto.ToAppIdentityProfile(*prof), dto.ListFilterFromRest(params.Body))

	switch {
	default:
		log.PrintErr("List server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return user.NewListUserDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("List client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return user.NewListUserDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("List ok")
		return user.NewListUserOK().WithPayload(&user.ListUserOKBody{
			Items:    dto.UsersToRest(c),
			Warnings: warnings,
		})
	}
}
