package rest

import (
	"errors"
	"washBonus/internal/app"
	"washBonus/internal/conversions"
	"washBonus/internal/entity"
	"washBonus/openapi/restapi/operations"
	washServers "washBonus/openapi/restapi/operations/wash_servers"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initWashServerHandlers(api *operations.WashBonusAPI) {
	api.WashServersGetWashServerByIDHandler = washServers.GetWashServerByIDHandlerFunc(svc.getWashServer)
	api.WashServersCreateWashServerHandler = washServers.CreateWashServerHandlerFunc(svc.createWashServer)
	api.WashServersUpdateWashServerHandler = washServers.UpdateWashServerHandlerFunc(svc.updateWashServer)
	api.WashServersDeleteWashServerHandler = washServers.DeleteWashServerHandlerFunc(svc.deleteWashServer)
	api.WashServersGetWashServersHandler = washServers.GetWashServersHandlerFunc(svc.getWashServers)

	api.WashServersAssignServerToGroupHandler = washServers.AssignServerToGroupHandlerFunc(svc.assignServerToGroup)
}

func (svc *service) getWashServer(params washServers.GetWashServerByIDParams, auth *app.Auth) washServers.GetWashServerByIDResponder {
	serverID, err := uuid.FromString(params.ServerID)
	if err != nil {
		return washServers.NewGetWashServerByIDBadRequest()
	}

	res, err := svc.washServerCtrl.GetWashServerById(params.HTTPRequest.Context(), *auth, serverID)

	switch {
	case err == nil:
		return washServers.NewGetWashServerByIDOK().WithPayload(conversions.WashServerToAdminRest(res))
	case errors.Is(err, entity.ErrAccessDenied):
		return washServers.NewGetWashServerByIDForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return washServers.NewGetWashServerByIDNotFound()
	default:
		svc.l.Errorln("Get wash server:", err)
		return washServers.NewGetWashServerByIDInternalServerError()
	}
}

func (svc *service) createWashServer(params washServers.CreateWashServerParams, auth *app.Auth) washServers.CreateWashServerResponder {
	createWashServerFromRest := conversions.WashServerCreationFromRest(*params.Body)

	newServer, err := svc.washServerCtrl.CreateWashServer(params.HTTPRequest.Context(), *auth, createWashServerFromRest)

	switch {
	case err == nil:
		return washServers.NewCreateWashServerOK().WithPayload(conversions.WashServerToAdminRest(newServer))
	case errors.Is(err, entity.ErrNotFound):
		return washServers.NewCreateWashServerBadRequest()
	case errors.Is(err, entity.ErrAccessDenied):
		return washServers.NewCreateWashServerForbidden()
	default:
		svc.l.Errorln("Create wash server:", err)
		return washServers.NewCreateWashServerInternalServerError()
	}
}

func (svc *service) updateWashServer(params washServers.UpdateWashServerParams, auth *app.Auth) washServers.UpdateWashServerResponder {
	serverID, err := uuid.FromString(params.ServerID)
	if err != nil {
		return washServers.NewUpdateWashServerBadRequest()
	}

	updateWashServerFromRest := conversions.WashServerUpdateFromRest(*params.Body)

	updatedServer, err := svc.washServerCtrl.UpdateWashServer(params.HTTPRequest.Context(), *auth, serverID, updateWashServerFromRest)

	switch {
	case err == nil:
		return washServers.NewUpdateWashServerOK().WithPayload(conversions.WashServerToRest(updatedServer))
	case errors.Is(err, entity.ErrNotFound):
		return washServers.NewUpdateWashServerNotFound()
	case errors.Is(err, entity.ErrAccessDenied):
		return washServers.NewUpdateWashServerForbidden()
	case errors.Is(err, entity.ErrBadRequest):
		return washServers.NewUpdateWashServerBadRequest()
	default:
		svc.l.Errorln("Update wash server:", err)
		return washServers.NewUpdateWashServerInternalServerError()
	}
}

func (svc *service) deleteWashServer(params washServers.DeleteWashServerParams, auth *app.Auth) washServers.DeleteWashServerResponder {
	serverID, err := uuid.FromString(params.ServerID)
	if err != nil {
		return washServers.NewDeleteWashServerBadRequest()
	}

	err = svc.washServerCtrl.DeleteWashServer(params.HTTPRequest.Context(), *auth, serverID)

	switch {
	case err == nil:
		return washServers.NewDeleteWashServerNoContent()
	case errors.Is(err, entity.ErrNotFound):
		return washServers.NewDeleteWashServerNotFound()
	case errors.Is(err, entity.ErrAccessDenied):
		return washServers.NewDeleteWashServerForbidden()
	default:
		svc.l.Errorln("Delete wash server:", err)
		return washServers.NewDeleteWashServerInternalServerError()
	}
}

func (svc *service) getWashServers(params washServers.GetWashServersParams, auth *app.Auth) washServers.GetWashServersResponder {
	pagination := conversions.PaginationFromRest(*params.Limit, *params.Offset)

	filter := conversions.WashServerFilterFromRest(pagination, *params.IsManagedByMe, params.OrganizationID, params.GroupID)

	res, err := svc.washServerCtrl.GetWashServers(params.HTTPRequest.Context(), *auth, filter)

	switch {
	case err == nil:
		payload := conversions.WashServerListToRest(res)
		return washServers.NewGetWashServersOK().WithPayload(payload)
	case errors.Is(err, entity.ErrAccessDenied):
		return washServers.NewGetWashServersForbidden()
	default:
		svc.l.Errorln("Get wash servers:", err)
		return washServers.NewGetWashServersInternalServerError()
	}
}

func (svc *service) assignServerToGroup(params washServers.AssignServerToGroupParams, auth *app.Auth) washServers.AssignServerToGroupResponder {
	serverID, err := uuid.FromString(params.ServerID.String())
	if err != nil {
		return washServers.NewAssignServerToGroupBadRequest()
	}

	groupID, err := uuid.FromString(params.GroupID.String())
	if err != nil {
		return washServers.NewAssignServerToGroupBadRequest()
	}

	err = svc.washServerCtrl.AssignToServerGroup(params.HTTPRequest.Context(), *auth, serverID, groupID)

	switch {
	case err == nil:
		return washServers.NewAssignServerToGroupNoContent()
	case errors.Is(err, entity.ErrBadRequest):
		return washServers.NewAssignServerToGroupBadRequest()
	case errors.Is(err, entity.ErrAccessDenied):
		return washServers.NewAssignServerToGroupForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return washServers.NewAssignServerToGroupNotFound()
	default:
		svc.l.Errorln("Assign wash server to server group:", err)
		return washServers.NewAssignServerToGroupInternalServerError()
	}
}
