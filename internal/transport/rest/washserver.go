package rest

import (
	"fmt"
	"washbonus/internal/app"
	"washbonus/internal/conversions"
	"washbonus/internal/entities"
	"washbonus/openapi/admin/restapi/operations"
	washServers "washbonus/openapi/admin/restapi/operations/wash_servers"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initWashServerHandlers(api *operations.WashAdminAPI) {
	api.WashServersGetWashServerByIDHandler = washServers.GetWashServerByIDHandlerFunc(svc.getWashServer)
	api.WashServersCreateWashServerHandler = washServers.CreateWashServerHandlerFunc(svc.createWashServer)
	api.WashServersUpdateWashServerHandler = washServers.UpdateWashServerHandlerFunc(svc.updateWashServer)
	api.WashServersDeleteWashServerHandler = washServers.DeleteWashServerHandlerFunc(svc.deleteWashServer)
	api.WashServersGetWashServersHandler = washServers.GetWashServersHandlerFunc(svc.getWashServers)

	api.WashServersAssignServerToGroupHandler = washServers.AssignServerToGroupHandlerFunc(svc.assignServerToGroup)
}

func (svc *service) getWashServer(params washServers.GetWashServerByIDParams, auth *app.AdminAuth) washServers.GetWashServerByIDResponder {
	op := "Get wash server by ID:"
	resp := washServers.NewGetWashServerByIDDefault(500)

	serverID, err := uuid.FromString(params.ServerID)
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong wash server ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	res, err := svc.washServerCtrl.GetWashServerById(params.HTTPRequest.Context(), *auth, serverID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return washServers.NewGetWashServerByIDOK().WithPayload(conversions.WashServerToAdminRest(res))
}

func (svc *service) createWashServer(params washServers.CreateWashServerParams, auth *app.AdminAuth) washServers.CreateWashServerResponder {
	op := "Create wash server:"
	resp := washServers.NewCreateWashServerDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	createWashServerFromRest := conversions.WashServerCreationFromRest(*params.Body)
	newServer, err := svc.washServerCtrl.CreateWashServer(ctx, *auth, createWashServerFromRest)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return washServers.NewCreateWashServerOK().WithPayload(conversions.WashServerToAdminRest(newServer))
}

func (svc *service) updateWashServer(params washServers.UpdateWashServerParams, auth *app.AdminAuth) washServers.UpdateWashServerResponder {
	op := "Update wash server:"
	resp := washServers.NewUpdateWashServerDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	serverID, err := uuid.FromString(params.ServerID)
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong wash server ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	updateWashServerFromRest := conversions.WashServerUpdateFromRest(*params.Body)
	updatedServer, err := svc.washServerCtrl.UpdateWashServer(ctx, *auth, serverID, updateWashServerFromRest)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return washServers.NewUpdateWashServerOK().WithPayload(conversions.WashServerToRest(updatedServer))
}

func (svc *service) deleteWashServer(params washServers.DeleteWashServerParams, auth *app.AdminAuth) washServers.DeleteWashServerResponder {
	op := "Delete wash server:"
	resp := washServers.NewDeleteWashServerDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	serverID, err := uuid.FromString(params.ServerID)
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong wash server ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	err = svc.washServerCtrl.DeleteWashServer(ctx, *auth, serverID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return washServers.NewDeleteWashServerNoContent()
}

func (svc *service) getWashServers(params washServers.GetWashServersParams, auth *app.AdminAuth) washServers.GetWashServersResponder {
	op := "Get wash servers:"
	resp := washServers.NewGetWashServersDefault(500)

	pagination := conversions.PaginationFromRest(*params.Limit, *params.Offset)
	filter := conversions.WashServerFilterFromRest(pagination, params.OrganizationID, params.GroupID)
	res, err := svc.washServerCtrl.GetWashServers(params.HTTPRequest.Context(), *auth, filter)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return washServers.NewGetWashServersOK().WithPayload(conversions.WashServerListToRest(res))
}

func (svc *service) assignServerToGroup(params washServers.AssignServerToGroupParams, auth *app.AdminAuth) washServers.AssignServerToGroupResponder {
	op := "Assign wash server to group:"
	resp := washServers.NewAssignServerToGroupDefault(500)
	ctx := createCtxWithUserID(params.HTTPRequest, auth)

	serverID, err := uuid.FromString(params.ServerID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong wash server ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	groupID, err := uuid.FromString(params.GroupID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong server group ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	err = svc.washServerCtrl.AssignToServerGroup(ctx, *auth, serverID, groupID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return washServers.NewAssignServerToGroupNoContent()
}
