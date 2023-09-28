package rest

import (
	"fmt"
	"washbonus/internal/app"
	conv "washbonus/internal/conversions"
	"washbonus/internal/entities"
	"washbonus/openapi/admin/restapi/operations"
	"washbonus/openapi/admin/restapi/operations/organizations"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initOrganizationsHandlers(api *operations.WashAdminAPI) {
	api.OrganizationsGetOrganizationsHandler = organizations.GetOrganizationsHandlerFunc(svc.getOrganizations)
	api.OrganizationsGetOrganizationByIDHandler = organizations.GetOrganizationByIDHandlerFunc(svc.getOrganizationByID)
	api.OrganizationsCreateOrganizationHandler = organizations.CreateOrganizationHandlerFunc(svc.createOrganization)
	api.OrganizationsUpdateOrganizationHandler = organizations.UpdateOrganizationHandlerFunc(svc.updateOrganization)
	api.OrganizationsDeleteOrganizationHandler = organizations.DeleteOrganizationHandlerFunc(svc.deleteOrganization)

	api.OrganizationsAssignUserToOrganizationHandler = organizations.AssignUserToOrganizationHandlerFunc(svc.assignOrganizationManager)
	api.OrganizationsRemoveUserFromOrganizationHandler = organizations.RemoveUserFromOrganizationHandlerFunc(svc.removeOrganizationManager)
}

func (svc *service) getOrganizations(params organizations.GetOrganizationsParams, auth *app.AdminAuth) organizations.GetOrganizationsResponder {
	op := "Get organizations:"
	resp := organizations.NewGetOrganizationsDefault(500)

	pagination := conv.PaginationFromRest(*params.Limit, *params.Offset)
	filter, err := conv.OrganizationFilterFromRest(pagination, params.Ids)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	res, err := svc.orgCtrl.Get(params.HTTPRequest.Context(), *auth, filter)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return organizations.NewGetOrganizationsOK().WithPayload(conv.OrganizationsToRest(res))
}

func (svc *service) getOrganizationByID(params organizations.GetOrganizationByIDParams, auth *app.AdminAuth) organizations.GetOrganizationByIDResponder {
	op := "Get organization by ID:"
	resp := organizations.NewGetOrganizationByIDDefault(500)

	id, err := uuid.FromString(params.OrganizationID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong organization ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	res, err := svc.orgCtrl.GetById(params.HTTPRequest.Context(), *auth, id)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return organizations.NewGetOrganizationByIDOK().WithPayload(conv.OrganizationToRest(res))
}

func (svc *service) createOrganization(params organizations.CreateOrganizationParams, auth *app.AdminAuth) organizations.CreateOrganizationResponder {
	op := "Create organization:"
	resp := organizations.NewCreateOrganizationDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	orgCreation := conv.OrganizationCreationFromRest(*params.Body)
	org, err := svc.orgCtrl.Create(ctx, *auth, orgCreation)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return organizations.NewCreateOrganizationOK().WithPayload(conv.OrganizationToRest(org))
}

func (svc *service) updateOrganization(params organizations.UpdateOrganizationParams, auth *app.AdminAuth) organizations.UpdateOrganizationResponder {
	op := "Update organization:"
	resp := organizations.NewUpdateOrganizationDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	id, err := uuid.FromString(params.OrganizationID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong organization ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	orgUpdate := conv.OrganizationUpdateFromRest(*params.Body)
	updatedOrg, err := svc.orgCtrl.Update(ctx, *auth, id, orgUpdate)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return organizations.NewUpdateOrganizationOK().WithPayload(conv.OrganizationToRest(updatedOrg))
}

func (svc *service) deleteOrganization(params organizations.DeleteOrganizationParams, auth *app.AdminAuth) organizations.DeleteOrganizationResponder {
	op := "Delete organization:"
	resp := organizations.NewDeleteOrganizationDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	id, err := uuid.FromString(params.OrganizationID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong organization ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	err = svc.orgCtrl.Delete(ctx, *auth, id)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return organizations.NewDeleteOrganizationNoContent()
}

func (svc *service) assignOrganizationManager(params organizations.AssignUserToOrganizationParams, auth *app.AdminAuth) organizations.AssignUserToOrganizationResponder {
	op := "Assign organization manager:"
	resp := organizations.NewAssignUserToOrganizationDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	id, err := uuid.FromString(params.OrganizationID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong organization ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	err = svc.orgCtrl.AssignManager(ctx, *auth, id, params.UserID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return organizations.NewAssignUserToOrganizationNoContent()
}

func (svc *service) removeOrganizationManager(params organizations.RemoveUserFromOrganizationParams, auth *app.AdminAuth) organizations.RemoveUserFromOrganizationResponder {
	op := "Remove organization manager:"
	resp := organizations.NewRemoveUserFromOrganizationDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	id, err := uuid.FromString(params.OrganizationID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong organization ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	err = svc.orgCtrl.RemoveManager(ctx, *auth, id, params.UserID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return organizations.NewRemoveUserFromOrganizationNoContent()
}
