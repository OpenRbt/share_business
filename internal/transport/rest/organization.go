package rest

import (
	"errors"
	"washBonus/internal/app"
	"washBonus/internal/conversions"
	"washBonus/internal/entity"
	"washBonus/openapi/restapi/operations"
	"washBonus/openapi/restapi/operations/organizations"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initOrganizationsHandlers(api *operations.WashBonusAPI) {
	api.OrganizationsGetOrganizationsHandler = organizations.GetOrganizationsHandlerFunc(svc.getOrganizations)
	api.OrganizationsGetOrganizationByIDHandler = organizations.GetOrganizationByIDHandlerFunc(svc.getOrganizationByID)
	api.OrganizationsCreateOrganizationHandler = organizations.CreateOrganizationHandlerFunc(svc.createOrganization)
	api.OrganizationsUpdateOrganizationHandler = organizations.UpdateOrganizationHandlerFunc(svc.updateOrganization)
	api.OrganizationsDeleteOrganizationHandler = organizations.DeleteOrganizationHandlerFunc(svc.deleteOrganization)

	api.OrganizationsAssignUserToOrganizationHandler = organizations.AssignUserToOrganizationHandlerFunc(svc.assignOrganizationManager)
	api.OrganizationsRemoveUserFromOrganizationHandler = organizations.RemoveUserFromOrganizationHandlerFunc(svc.removeOrganizationManager)

	api.OrganizationsGetSettingsForOrganizationHandler = organizations.GetSettingsForOrganizationHandlerFunc(svc.getSettingsForOrganization)
	api.OrganizationsUpdateSettingForOrganizationHandler = organizations.UpdateSettingForOrganizationHandlerFunc(svc.updateSettingsForOrganization)
}

func (svc *service) getOrganizations(params organizations.GetOrganizationsParams, auth *app.Auth) organizations.GetOrganizationsResponder {
	pagination := conversions.PaginationFromRest(*params.Limit, *params.Offset)

	filter, err := conversions.OrganizationFilterFromRest(pagination, *params.IsManagedByMe, params.Ids)
	if err != nil {
		svc.l.Errorln(err)
		return organizations.NewGetOrganizationsBadRequest()
	}

	res, err := svc.orgCtrl.Get(params.HTTPRequest.Context(), *auth, filter)

	switch {
	case err == nil:
		return organizations.NewGetOrganizationsOK().WithPayload(conversions.OrganizationsToRest(res))
	case errors.Is(err, entity.ErrAccessDenied):
		return organizations.NewGetOrganizationsForbidden()
	default:
		svc.l.Errorln("Get organizations:", err)
		return organizations.NewGetOrganizationsInternalServerError()
	}
}

func (svc *service) getOrganizationByID(params organizations.GetOrganizationByIDParams, auth *app.Auth) organizations.GetOrganizationByIDResponder {
	id, err := uuid.FromString(params.OrganizationID.String())
	if err != nil {
		return organizations.NewGetOrganizationByIDBadRequest()
	}

	res, err := svc.orgCtrl.GetById(params.HTTPRequest.Context(), *auth, id)

	switch {
	case err == nil:
		return organizations.NewGetOrganizationByIDOK().WithPayload(conversions.OrganizationToRest(res))
	case errors.Is(err, entity.ErrAccessDenied):
		return organizations.NewGetOrganizationByIDForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return organizations.NewGetOrganizationByIDNotFound()
	default:
		svc.l.Errorln("Get organization:", err)
		return organizations.NewGetOrganizationByIDInternalServerError()
	}
}

func (svc *service) createOrganization(params organizations.CreateOrganizationParams, auth *app.Auth) organizations.CreateOrganizationResponder {
	orgCreation := conversions.OrganizationCreationFromRest(*params.Body)

	org, err := svc.orgCtrl.Create(params.HTTPRequest.Context(), *auth, orgCreation)

	switch {
	case err == nil:
		return organizations.NewCreateOrganizationOK().WithPayload(conversions.OrganizationToRest(org))
	case errors.Is(err, entity.ErrAccessDenied):
		return organizations.NewCreateOrganizationForbidden()
	default:
		svc.l.Errorln("Create organization:", err)
		return organizations.NewCreateOrganizationInternalServerError()
	}
}

func (svc *service) updateOrganization(params organizations.UpdateOrganizationParams, auth *app.Auth) organizations.UpdateOrganizationResponder {
	id, err := uuid.FromString(params.OrganizationID.String())
	if err != nil {
		return organizations.NewUpdateOrganizationBadRequest()
	}

	orgUpdate := conversions.OrganizationUpdateFromRest(*params.Body)
	updatedOrg, err := svc.orgCtrl.Update(params.HTTPRequest.Context(), *auth, id, orgUpdate)

	switch {
	case err == nil:
		return organizations.NewUpdateOrganizationOK().WithPayload(conversions.OrganizationToRest(updatedOrg))
	case errors.Is(err, entity.ErrAccessDenied):
		return organizations.NewUpdateOrganizationForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return organizations.NewUpdateOrganizationNotFound()
	default:
		svc.l.Errorln("Update organization:", err)
		return organizations.NewUpdateOrganizationInternalServerError()
	}
}

func (svc *service) deleteOrganization(params organizations.DeleteOrganizationParams, auth *app.Auth) organizations.DeleteOrganizationResponder {
	id, err := uuid.FromString(params.OrganizationID.String())
	if err != nil {
		return organizations.NewDeleteOrganizationBadRequest()
	}

	err = svc.orgCtrl.Delete(params.HTTPRequest.Context(), *auth, id)

	switch {
	case err == nil:
		return organizations.NewDeleteOrganizationNoContent()
	case errors.Is(err, entity.ErrAccessDenied):
		return organizations.NewDeleteOrganizationForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return organizations.NewDeleteOrganizationNotFound()
	default:
		svc.l.Errorln("Delete organization:", err)
		return organizations.NewDeleteOrganizationInternalServerError()
	}
}

func (svc *service) assignOrganizationManager(params organizations.AssignUserToOrganizationParams, auth *app.Auth) organizations.AssignUserToOrganizationResponder {
	id, err := uuid.FromString(params.OrganizationID.String())
	if err != nil {
		return organizations.NewAssignUserToOrganizationBadRequest()
	}

	err = svc.orgCtrl.AssignManager(params.HTTPRequest.Context(), *auth, id, params.UserID)

	switch {
	case err == nil:
		return organizations.NewAssignUserToOrganizationNoContent()
	case errors.Is(err, entity.ErrAccessDenied):
		return organizations.NewAssignUserToOrganizationForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return organizations.NewAssignUserToOrganizationNotFound()
	default:
		svc.l.Errorln("Assign user to organization:", err)
		return organizations.NewAssignUserToOrganizationInternalServerError()
	}
}

func (svc *service) removeOrganizationManager(params organizations.RemoveUserFromOrganizationParams, auth *app.Auth) organizations.RemoveUserFromOrganizationResponder {
	id, err := uuid.FromString(params.OrganizationID.String())
	if err != nil {
		return organizations.NewRemoveUserFromOrganizationBadRequest()
	}

	err = svc.orgCtrl.RemoveManager(params.HTTPRequest.Context(), *auth, id, params.UserID)

	switch {
	case err == nil:
		return organizations.NewRemoveUserFromOrganizationNoContent()
	case errors.Is(err, entity.ErrAccessDenied):
		return organizations.NewRemoveUserFromOrganizationForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return organizations.NewRemoveUserFromOrganizationNotFound()
	default:
		svc.l.Errorln("Remove user from organization:", err)
		return organizations.NewRemoveUserFromOrganizationInternalServerError()
	}
}

func (svc *service) getSettingsForOrganization(params organizations.GetSettingsForOrganizationParams, auth *app.Auth) organizations.GetSettingsForOrganizationResponder {
	id, err := uuid.FromString(params.ID.String())
	if err != nil {
		return organizations.NewGetSettingsForOrganizationBadRequest()
	}

	settings, err := svc.orgCtrl.GetSettingsForOrganization(params.HTTPRequest.Context(), *auth, id)

	switch {
	case err == nil:
		payload := conversions.OrganizationSettingsToRest(settings)
		return organizations.NewGetSettingsForOrganizationOK().WithPayload(&payload)
	case errors.Is(err, entity.ErrAccessDenied):
		return organizations.NewGetSettingsForOrganizationForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return organizations.NewGetSettingsForOrganizationNotFound()
	default:
		svc.l.Errorln("Get organization settings:", err)
		return organizations.NewGetSettingsForOrganizationInternalServerError()
	}
}

func (svc *service) updateSettingsForOrganization(params organizations.UpdateSettingForOrganizationParams, auth *app.Auth) organizations.UpdateSettingForOrganizationResponder {
	id, err := uuid.FromString(params.ID.String())
	if err != nil {
		return organizations.NewUpdateSettingForOrganizationBadRequest()
	}

	settings, err := svc.orgCtrl.UpdateSettingsForOrganization(params.HTTPRequest.Context(), *auth, id, conversions.OrganizationSettingsUpdateFromRest(*params.Body))

	switch {
	case err == nil:
		payload := conversions.OrganizationSettingsToRest(settings)
		return organizations.NewUpdateSettingForOrganizationOK().WithPayload(&payload)
	case errors.Is(err, entity.ErrAccessDenied):
		return organizations.NewUpdateSettingForOrganizationForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return organizations.NewUpdateSettingForOrganizationNotFound()
	default:
		svc.l.Errorln("Update organization settings:", err)
		return organizations.NewUpdateSettingForOrganizationInternalServerError()
	}
}
