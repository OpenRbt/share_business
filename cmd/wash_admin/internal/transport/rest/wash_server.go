package rest

import (
	"errors"
	"wash_admin/internal/app"
	"wash_admin/internal/conversions"
	"wash_admin/internal/entity"
	"wash_admin/openapi/restapi/operations"
	"wash_admin/openapi/restapi/operations/wash_servers"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initWashServerHandlers(api *operations.WashAdminAPI) {
	api.WashServersGetHandler = wash_servers.GetHandlerFunc(svc.getWashServer)
	api.WashServersAddHandler = wash_servers.AddHandlerFunc(svc.addWashServer)
	api.WashServersUpdateHandler = wash_servers.UpdateHandlerFunc(svc.updateWashServer)
	api.WashServersDeleteHandler = wash_servers.DeleteHandlerFunc(svc.deleteWashServer)
}

func (svc *service) getWashServer(params wash_servers.GetParams, auth *app.Auth) wash_servers.GetResponder {
	if params.Body.ID == nil {
		return wash_servers.NewGetBadRequest()
	}

	id, err := uuid.FromString(*params.Body.ID)

	if err != nil {
		return wash_servers.NewGetBadRequest()
	}

	res, err := svc.washServers.GetWashServer(params.HTTPRequest.Context(), auth, id)

	payload := conversions.WashServerToRest(res)

	switch {
	case err == nil:
		return wash_servers.NewGetOK().WithPayload(&payload)
	case errors.Is(err, entity.ErrNotFound):
		return wash_servers.NewGetNotFound()
	default:
		return wash_servers.NewGetInternalServerError()
	}
}

func (svc *service) addWashServer(params wash_servers.AddParams, auth *app.Auth) wash_servers.AddResponder {
	addWashServerFromRest := conversions.AddWashServerFromRest(*params.Body)

	err := svc.washServers.AddWashServer(params.HTTPRequest.Context(), auth, addWashServerFromRest)

	switch {
	case err == nil:
		return wash_servers.NewAddNoContent()
	case errors.Is(err, entity.ErrNotFound):
		return wash_servers.NewAddBadRequest()
	default:
		return wash_servers.NewAddInternalServerError()
	}
}

func (svc *service) updateWashServer(params wash_servers.UpdateParams, auth *app.Auth) wash_servers.UpdateResponder {
	updateWashServerFromRest, err := conversions.UpdateWashServerFromRest(*params.Body)

	if err != nil {
		return wash_servers.NewUpdateBadRequest()
	}

	err = svc.washServers.UpdateWashServer(params.HTTPRequest.Context(), auth, updateWashServerFromRest)

	switch {
	case err == nil:
		return wash_servers.NewUpdateNoContent()
	case errors.Is(err, entity.ErrNotFound):
		return wash_servers.NewUpdateNotFound()
	default:
		return wash_servers.NewUpdateInternalServerError()
	}
}

func (svc *service) deleteWashServer(params wash_servers.DeleteParams, auth *app.Auth) wash_servers.DeleteResponder {
	deleteWashServerFromRest, err := conversions.DeleteWashServerFromRest(*params.Body)

	if err != nil {
		return wash_servers.NewDeleteBadRequest()
	}

	err = svc.washServers.DeleteWashServer(params.HTTPRequest.Context(), auth, deleteWashServerFromRest)

	switch {
	case err == nil:
		return wash_servers.NewDeleteNoContent()
	case errors.Is(err, entity.ErrNotFound):
		return wash_servers.NewDeleteNotFound()
	default:
		return wash_servers.NewDeleteInternalServerError()
	}
}

func (svc *service) getWashServerList(params wash_servers.ListParams, auth *app.Auth) wash_servers.ListResponder {
	paginationFromRest := conversions.PaginationFromRest(*params.Body)

	res, err := svc.washServers.GetWashServerList(params.HTTPRequest.Context(), auth, paginationFromRest)

	payload := conversions.WashServerListToRest(res)

	switch {
	case err == nil:
		return wash_servers.NewListOK().WithPayload(payload)
	case errors.Is(err, entity.ErrNotFound):
		return wash_servers.NewListNotFound()
	default:
		return wash_servers.NewListInternalServerError()
	}
}