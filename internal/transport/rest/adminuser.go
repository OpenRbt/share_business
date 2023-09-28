package rest

import (
	"fmt"
	"washbonus/internal/app"
	"washbonus/internal/conversions"
	"washbonus/internal/entities"
	"washbonus/openapi/admin/restapi/operations"
	"washbonus/openapi/admin/restapi/operations/applications"
	"washbonus/openapi/admin/restapi/operations/users"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initAdminUserHandlers(api *operations.WashAdminAPI) {
	api.UsersGetAdminUserByIDHandler = users.GetAdminUserByIDHandlerFunc(svc.getAdminUserByID)
	api.UsersUpdateAdminUserRoleHandler = users.UpdateAdminUserRoleHandlerFunc(svc.updateAdminUserRole)
	api.UsersGetAdminUsersHandler = users.GetAdminUsersHandlerFunc(svc.getAdminUsers)
	api.UsersBlockAdminUserHandler = users.BlockAdminUserHandlerFunc(svc.blockAdminUser)

	api.ApplicationsGetAdminApplicationsHandler = applications.GetAdminApplicationsHandlerFunc(svc.getAdminApplications)
	api.ApplicationsGetAdminApplicationByIDHandler = applications.GetAdminApplicationByIDHandlerFunc(svc.getAdminApplicationByID)
	api.ApplicationsCreateAdminApplicationHandler = applications.CreateAdminApplicationHandlerFunc(svc.createAdminApplication)
	api.ApplicationsReviewAdminApplicationHandler = applications.ReviewAdminApplicationHandlerFunc(svc.reviewAdminApplication)
}

func (svc *service) getAdminUsers(params users.GetAdminUsersParams, auth *app.AdminAuth) users.GetAdminUsersResponder {
	op := "Get admin users:"
	resp := users.NewGetAdminUsersDefault(500)

	filter, err := conversions.AdminUserFilterFromRest(*params.Limit, *params.Offset, params.Role, params.IsBlocked)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	res, err := svc.adminCtrl.Get(params.HTTPRequest.Context(), *auth, filter)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	payload := conversions.AdminUsersToRest(res)
	return users.NewGetAdminUsersOK().WithPayload(payload)
}

func (svc *service) getAdminUserByID(params users.GetAdminUserByIDParams, auth *app.AdminAuth) users.GetAdminUserByIDResponder {
	op := "Get admin user by ID:"
	resp := users.NewGetAdminUserByIDDefault(500)

	res, err := svc.adminCtrl.GetById(params.HTTPRequest.Context(), *auth, params.UserID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	payload := conversions.AdminUserToRest(res)
	return users.NewGetAdminUserByIDOK().WithPayload(&payload)
}

func (svc *service) updateAdminUserRole(params users.UpdateAdminUserRoleParams, auth *app.AdminAuth) users.UpdateAdminUserRoleResponder {
	op := "Update admin user:"
	resp := users.NewUpdateAdminUserRoleDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	updateRole, err := conversions.AdminUserRoleUpdateFromRest(params.UserID, params.Body.Role)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	err = svc.adminCtrl.UpdateRole(ctx, *auth, updateRole)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return users.NewUpdateAdminUserRoleNoContent()
}

func (svc *service) blockAdminUser(params users.BlockAdminUserParams, auth *app.AdminAuth) users.BlockAdminUserResponder {
	op := "Block admin user:"
	resp := users.NewBlockAdminUserDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	err := svc.adminCtrl.Block(ctx, *auth, params.UserID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return users.NewBlockAdminUserNoContent()
}

func (svc *service) getAdminApplications(params applications.GetAdminApplicationsParams, auth *app.AdminAuth) applications.GetAdminApplicationsResponder {
	op := "Get admin applications:"
	resp := applications.NewGetAdminApplicationsDefault(500)

	pagination := conversions.PaginationFromRest(*params.Limit, *params.Offset)
	res, err := svc.adminCtrl.GetApplications(params.HTTPRequest.Context(), *auth, conversions.AdminApplicationFilterFromRest(pagination, params.Status))
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	payload := conversions.AdminApplicationsToRest(res)
	return applications.NewGetAdminApplicationsOK().WithPayload(&applications.GetAdminApplicationsOKBody{Applications: payload})
}

func (svc *service) createAdminApplication(params applications.CreateAdminApplicationParams) applications.CreateAdminApplicationResponder {
	op := "Create admin application:"
	resp := applications.NewCreateAdminApplicationDefault(500)

	res, err := svc.adminCtrl.CreateApplication(params.HTTPRequest.Context(), conversions.AdminApplicationCreationFromRest(*params.Body.Application))
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	payload := conversions.AdminApplicationToRest(res)
	return applications.NewCreateAdminApplicationOK().WithPayload(&payload)
}

func (svc *service) reviewAdminApplication(params applications.ReviewAdminApplicationParams, auth *app.AdminAuth) applications.ReviewAdminApplicationResponder {
	op := "Review admin application:"
	resp := applications.NewReviewAdminApplicationDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	id, err := uuid.FromString(params.ID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong admin application ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	review, err := conversions.AdminApplicationReviewFromRest(*params.Body)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	err = svc.adminCtrl.ReviewApplication(ctx, *auth, id, review)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return applications.NewReviewAdminApplicationNoContent()
}

func (svc *service) getAdminApplicationByID(params applications.GetAdminApplicationByIDParams, auth *app.AdminAuth) applications.GetAdminApplicationByIDResponder {
	const op = "Get admin application by ID:"
	resp := applications.NewGetAdminApplicationByIDDefault(500)

	id, err := uuid.FromString(params.ID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong Application ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	res, err := svc.adminCtrl.GetApplicationByID(params.HTTPRequest.Context(), *auth, id)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	payload := conversions.AdminApplicationToRest(res)
	return applications.NewGetAdminApplicationByIDOK().WithPayload(&payload)
}
