package rest

import (
	"errors"
	"log"
	"wash_admin/internal/app"

	"wash_admin/openapi/restapi/operations"
	"wash_admin/openapi/restapi/operations/wash_servers"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initWashServerHandlers(api *operations.WashAdminAPI) {
	api.WashServersGetWashServerHandler = wash_servers.GetWashServerHandlerFunc(svc.getWashServer)
	api.WashServersAddHandler = wash_servers.AddHandlerFunc(svc.addWashServer)
	api.WashServersUpdateHandler = wash_servers.UpdateHandlerFunc(svc.updateWashServer)
	api.WashServersDeleteHandler = wash_servers.DeleteHandlerFunc(svc.deleteWashServer)
}

func (svc *service) getWashServer(params wash_servers.GetWashServerParams, auth *app.Auth) wash_servers.GetWashServerResponder {
	id, err := uuid.FromString(params.ID)

	if err != nil {
		return wash_servers.NewGetWashServerBadRequest()
	}

	res, err := svc.washServers.GetWashServer(params.HTTPRequest.Context(), auth, id)

	switch {
	case err == nil:
		return wash_servers.NewGetWashServerOK().WithPayload(WashServerToRest(res))
	case errors.Is(err, app.ErrNotFound):
		return wash_servers.NewGetWashServerNotFound()
	default:
		return wash_servers.NewGetWashServerInternalServerError()
	}
}

func (svc *service) addWashServer(params wash_servers.AddParams, auth *app.Auth) wash_servers.AddResponder {
	registerWashServerFromRest := RegisterWashServerFromRest(*params.Body)

	newServer, err := svc.washServers.RegisterWashServer(params.HTTPRequest.Context(), auth, registerWashServerFromRest)

	if err != nil {
		log.Println(err)
	}

	switch {
	case err == nil:
		return wash_servers.NewAddOK().WithPayload(WashServerToRest(newServer))
	case errors.Is(err, app.ErrNotFound):
		return wash_servers.NewAddBadRequest()
	default:
		return wash_servers.NewAddInternalServerError()
	}
}

func (svc *service) updateWashServer(params wash_servers.UpdateParams, auth *app.Auth) wash_servers.UpdateResponder {
	updateWashServerFromRest, err := UpdateWashServerFromRest(*params.Body)

	if err != nil {
		return wash_servers.NewUpdateBadRequest()
	}

	err = svc.washServers.UpdateWashServer(params.HTTPRequest.Context(), auth, updateWashServerFromRest)

	switch {
	case err == nil:
		return wash_servers.NewUpdateNoContent()
	case errors.Is(err, app.ErrNotFound):
		return wash_servers.NewUpdateNotFound()
	default:
		return wash_servers.NewUpdateInternalServerError()
	}
}

func (svc *service) deleteWashServer(params wash_servers.DeleteParams, auth *app.Auth) wash_servers.DeleteResponder {
	deleteWashServerFromRest, err := DeleteWashServerFromRest(*params.Body)

	if err != nil {
		return wash_servers.NewDeleteBadRequest()
	}

	err = svc.washServers.DeleteWashServer(params.HTTPRequest.Context(), auth, deleteWashServerFromRest)

	switch {
	case err == nil:
		return wash_servers.NewDeleteNoContent()
	case errors.Is(err, app.ErrNotFound):
		return wash_servers.NewDeleteNotFound()
	default:
		return wash_servers.NewDeleteInternalServerError()
	}
}

func (svc *service) getWashServerList(params wash_servers.ListParams, auth *app.Auth) wash_servers.ListResponder {
	res, err := svc.washServers.GetWashServerList(params.HTTPRequest.Context(), auth, PaginationFromRest(*params.Body))

	payload := WashServerListToRest(res)

	switch {
	case err == nil:
		return wash_servers.NewListOK().WithPayload(payload)
	case errors.Is(err, app.ErrNotFound):
		return wash_servers.NewListNotFound()
	default:
		return wash_servers.NewListInternalServerError()
	}
}
