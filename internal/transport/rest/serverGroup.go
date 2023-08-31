package rest

import (
	"errors"
	"washBonus/internal/app"
	"washBonus/internal/conversions"
	"washBonus/internal/entity"
	"washBonus/openapi/restapi/operations"
	serverGroups "washBonus/openapi/restapi/operations/server_groups"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initServerGroupHandlers(api *operations.WashBonusAPI) {
	api.ServerGroupsGetServerGroupsHandler = serverGroups.GetServerGroupsHandlerFunc(svc.getServerGroups)
	api.ServerGroupsGetServerGroupByIDHandler = serverGroups.GetServerGroupByIDHandlerFunc(svc.getServerGroupByID)
	api.ServerGroupsCreateServerGroupHandler = serverGroups.CreateServerGroupHandlerFunc(svc.createServerGroup)
	api.ServerGroupsUpdateServerGroupHandler = serverGroups.UpdateServerGroupHandlerFunc(svc.updateServerGroup)
	api.ServerGroupsDeleteServerGroupHandler = serverGroups.DeleteServerGroupHandlerFunc(svc.deleteServerGroup)
}

func (svc *service) getServerGroups(params serverGroups.GetServerGroupsParams, auth *app.Auth) serverGroups.GetServerGroupsResponder {
	pagination := conversions.PaginationFromRest(*params.Limit, *params.Offset)
	filter := conversions.ServerGroupFilterFromRest(pagination, *params.IsManagedByMe, params.OrganizationID)

	res, err := svc.serverGroupCtrl.Get(params.HTTPRequest.Context(), *auth, filter)

	switch {
	case err == nil:
		return serverGroups.NewGetServerGroupsOK().WithPayload(conversions.ServerGroupsToRest(res))
	case errors.Is(err, entity.ErrAccessDenied):
		return serverGroups.NewGetServerGroupsForbidden()
	default:
		svc.l.Errorln("Get server groups:", err)
		return serverGroups.NewGetServerGroupsInternalServerError()
	}
}

func (svc *service) getServerGroupByID(params serverGroups.GetServerGroupByIDParams, auth *app.Auth) serverGroups.GetServerGroupByIDResponder {
	groupID, err := uuid.FromString(params.GroupID.String())
	if err != nil {
		return serverGroups.NewGetServerGroupByIDBadRequest()
	}

	res, err := svc.serverGroupCtrl.GetById(params.HTTPRequest.Context(), *auth, groupID)

	switch {
	case err == nil:
		return serverGroups.NewGetServerGroupByIDOK().WithPayload(conversions.ServerGroupToRest(res))
	case errors.Is(err, entity.ErrAccessDenied):
		return serverGroups.NewGetServerGroupByIDForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return serverGroups.NewGetServerGroupByIDNotFound()
	default:
		svc.l.Errorln("Get server group:", err)
		return serverGroups.NewGetServerGroupByIDInternalServerError()
	}
}

func (svc *service) createServerGroup(params serverGroups.CreateServerGroupParams, auth *app.Auth) serverGroups.CreateServerGroupResponder {
	groupCreation := conversions.ServerGroupCreationFromRest(*params.Body)

	group, err := svc.serverGroupCtrl.Create(params.HTTPRequest.Context(), *auth, groupCreation)

	switch {
	case err == nil:
		return serverGroups.NewCreateServerGroupOK().WithPayload(conversions.ServerGroupToRest(group))
	case errors.Is(err, entity.ErrAccessDenied):
		return serverGroups.NewCreateServerGroupForbidden()
	default:
		svc.l.Errorln("Create server group:", err)
		return serverGroups.NewCreateServerGroupInternalServerError()
	}
}

func (svc *service) updateServerGroup(params serverGroups.UpdateServerGroupParams, auth *app.Auth) serverGroups.UpdateServerGroupResponder {
	groupID, err := uuid.FromString(params.GroupID.String())
	if err != nil {
		return serverGroups.NewUpdateServerGroupBadRequest()
	}

	groupUpdate := conversions.ServerGroupUpdateFromRest(*params.Body)
	updatedGroup, err := svc.serverGroupCtrl.Update(params.HTTPRequest.Context(), *auth, groupID, groupUpdate)

	switch {
	case err == nil:
		return serverGroups.NewUpdateServerGroupOK().WithPayload(conversions.ServerGroupToRest(updatedGroup))
	case errors.Is(err, entity.ErrAccessDenied):
		return serverGroups.NewUpdateServerGroupForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return serverGroups.NewUpdateServerGroupNotFound()
	default:
		svc.l.Errorln("Update server group:", err)
		return serverGroups.NewUpdateServerGroupInternalServerError()
	}
}

func (svc *service) deleteServerGroup(params serverGroups.DeleteServerGroupParams, auth *app.Auth) serverGroups.DeleteServerGroupResponder {
	groupID, err := uuid.FromString(params.GroupID.String())
	if err != nil {
		return serverGroups.NewDeleteServerGroupBadRequest()
	}

	err = svc.serverGroupCtrl.Delete(params.HTTPRequest.Context(), *auth, groupID)

	switch {
	case err == nil:
		return serverGroups.NewDeleteServerGroupNoContent()
	case errors.Is(err, entity.ErrAccessDenied):
		return serverGroups.NewDeleteServerGroupForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return serverGroups.NewDeleteServerGroupNotFound()
	default:
		svc.l.Errorln("Delete server group:", err)
		return serverGroups.NewDeleteServerGroupInternalServerError()
	}
}
