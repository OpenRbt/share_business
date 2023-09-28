package rest

import (
	"fmt"
	"washbonus/internal/app"
	"washbonus/internal/conversions"
	"washbonus/internal/entities"
	"washbonus/openapi/admin/restapi/operations"
	serverGroups "washbonus/openapi/admin/restapi/operations/server_groups"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initServerGroupHandlers(api *operations.WashAdminAPI) {
	api.ServerGroupsGetServerGroupsHandler = serverGroups.GetServerGroupsHandlerFunc(svc.getServerGroups)
	api.ServerGroupsGetServerGroupByIDHandler = serverGroups.GetServerGroupByIDHandlerFunc(svc.getServerGroupByID)
	api.ServerGroupsCreateServerGroupHandler = serverGroups.CreateServerGroupHandlerFunc(svc.createServerGroup)
	api.ServerGroupsUpdateServerGroupHandler = serverGroups.UpdateServerGroupHandlerFunc(svc.updateServerGroup)
	api.ServerGroupsDeleteServerGroupHandler = serverGroups.DeleteServerGroupHandlerFunc(svc.deleteServerGroup)
}

func (svc *service) getServerGroups(params serverGroups.GetServerGroupsParams, auth *app.AdminAuth) serverGroups.GetServerGroupsResponder {
	op := "Get server groups:"
	resp := serverGroups.NewGetServerGroupsDefault(500)

	pagination := conversions.PaginationFromRest(*params.Limit, *params.Offset)
	filter := conversions.ServerGroupFilterFromRest(pagination, params.OrganizationID)

	res, err := svc.serverGroupCtrl.Get(params.HTTPRequest.Context(), *auth, filter)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return serverGroups.NewGetServerGroupsOK().WithPayload(conversions.ServerGroupsToRest(res))
}

func (svc *service) getServerGroupByID(params serverGroups.GetServerGroupByIDParams, auth *app.AdminAuth) serverGroups.GetServerGroupByIDResponder {
	op := "Get server group by ID:"
	resp := serverGroups.NewGetServerGroupByIDDefault(500)

	groupID, err := uuid.FromString(params.GroupID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong server group ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	res, err := svc.serverGroupCtrl.GetById(params.HTTPRequest.Context(), *auth, groupID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return serverGroups.NewGetServerGroupByIDOK().WithPayload(conversions.ServerGroupToRest(res))
}

func (svc *service) createServerGroup(params serverGroups.CreateServerGroupParams, auth *app.AdminAuth) serverGroups.CreateServerGroupResponder {
	op := "Create server group:"
	resp := serverGroups.NewCreateServerGroupDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	groupCreation := conversions.ServerGroupCreationFromRest(*params.Body)
	group, err := svc.serverGroupCtrl.Create(ctx, *auth, groupCreation)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return serverGroups.NewCreateServerGroupOK().WithPayload(conversions.ServerGroupToRest(group))
}

func (svc *service) updateServerGroup(params serverGroups.UpdateServerGroupParams, auth *app.AdminAuth) serverGroups.UpdateServerGroupResponder {
	op := "Update server group:"
	resp := serverGroups.NewUpdateServerGroupDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	groupID, err := uuid.FromString(params.GroupID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong server group ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	groupUpdate := conversions.ServerGroupUpdateFromRest(*params.Body)
	updatedGroup, err := svc.serverGroupCtrl.Update(ctx, *auth, groupID, groupUpdate)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return serverGroups.NewUpdateServerGroupOK().WithPayload(conversions.ServerGroupToRest(updatedGroup))
}

func (svc *service) deleteServerGroup(params serverGroups.DeleteServerGroupParams, auth *app.AdminAuth) serverGroups.DeleteServerGroupResponder {
	op := "Delete server group:"
	resp := serverGroups.NewDeleteServerGroupDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	groupID, err := uuid.FromString(params.GroupID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong server group ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	err = svc.serverGroupCtrl.Delete(ctx, *auth, groupID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return serverGroups.NewDeleteServerGroupNoContent()
}
