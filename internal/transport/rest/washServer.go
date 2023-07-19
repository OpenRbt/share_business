package rest

import (
	"errors"
	"washBonus/internal/app"
	"washBonus/internal/conversions"
	"washBonus/internal/entity"
	"washBonus/openapi/restapi/operations"
	"washBonus/openapi/restapi/operations/wash_servers"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initWashServerHandlers(api *operations.WashBonusAPI) {
	api.WashServersGetWashServerByIDHandler = wash_servers.GetWashServerByIDHandlerFunc(svc.getWashServer)
	api.WashServersCreateWashServerHandler = wash_servers.CreateWashServerHandlerFunc(svc.createWashServer)
	api.WashServersUpdateWashServerHandler = wash_servers.UpdateWashServerHandlerFunc(svc.updateWashServer)
	api.WashServersDeleteHandler = wash_servers.DeleteHandlerFunc(svc.deleteWashServer)
	api.WashServersListHandler = wash_servers.ListHandlerFunc(svc.getWashServerList)
}

func (svc *service) getWashServer(params wash_servers.GetWashServerByIDParams, auth *app.Auth) wash_servers.GetWashServerByIDResponder {
	id, err := uuid.FromString(params.ID)

	if err != nil {
		return wash_servers.NewGetWashServerByIDBadRequest()
	}

	res, err := svc.washServerCtrl.GetWashServerById(params.HTTPRequest.Context(), auth.UID, id)

	switch {
	case err == nil:
		return wash_servers.NewGetWashServerByIDOK().WithPayload(conversions.WashServerToAdminRest(res))
	case errors.Is(err, entity.ErrAccessDenied):
		return wash_servers.NewGetWashServerByIDForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return wash_servers.NewGetWashServerByIDNotFound()
	default:
		return wash_servers.NewGetWashServerByIDInternalServerError()
	}
}

func (svc *service) createWashServer(params wash_servers.CreateWashServerParams, auth *app.Auth) wash_servers.CreateWashServerResponder {
	createWashServerFromRest := conversions.CreateWashServerFromRest(*params.Body)

	newServer, err := svc.washServerCtrl.CreateWashServer(params.HTTPRequest.Context(), auth.UID, createWashServerFromRest)

	switch {
	case err == nil:
		return wash_servers.NewCreateWashServerOK().WithPayload(conversions.WashServerToAdminRest(newServer))
	case errors.Is(err, entity.ErrNotFound):
		return wash_servers.NewCreateWashServerBadRequest()
	case errors.Is(err, entity.ErrAccessDenied):
		return wash_servers.NewCreateWashServerForbidden()
	default:
		return wash_servers.NewCreateWashServerInternalServerError()
	}
}

func (svc *service) updateWashServer(params wash_servers.UpdateWashServerParams, auth *app.Auth) wash_servers.UpdateWashServerResponder {
	updateWashServerFromRest := conversions.UpdateWashServerFromRest(*params.Body)

	updatedServer, err := svc.washServerCtrl.UpdateWashServer(params.HTTPRequest.Context(), auth.UID, uuid.FromStringOrNil(params.ID), updateWashServerFromRest)

	switch {
	case err == nil:
		return wash_servers.NewUpdateWashServerOK().WithPayload(conversions.WashServerToRest(updatedServer))
	case errors.Is(err, entity.ErrNotFound):
		return wash_servers.NewUpdateWashServerNotFound()
	case errors.Is(err, entity.ErrAccessDenied):
		return wash_servers.NewUpdateWashServerForbidden()
	case errors.Is(err, entity.ErrBadRequest):
		return wash_servers.NewUpdateWashServerBadRequest()
	default:
		return wash_servers.NewUpdateWashServerInternalServerError()
	}
}

func (svc *service) deleteWashServer(params wash_servers.DeleteParams, auth *app.Auth) wash_servers.DeleteResponder {
	err := svc.washServerCtrl.DeleteWashServer(params.HTTPRequest.Context(), auth.UID, uuid.FromStringOrNil(params.ID))

	switch {
	case err == nil:
		return wash_servers.NewDeleteNoContent()
	case errors.Is(err, entity.ErrNotFound):
		return wash_servers.NewDeleteNotFound()
	case errors.Is(err, entity.ErrAccessDenied):
		return wash_servers.NewDeleteForbidden()
	default:
		return wash_servers.NewDeleteInternalServerError()
	}
}

func (svc *service) getWashServerList(params wash_servers.ListParams, auth *app.Auth) wash_servers.ListResponder {
	pagination := conversions.PaginationFromRest(*params.Body)

	res, err := svc.washServerCtrl.GetWashServers(params.HTTPRequest.Context(), auth.UID, pagination)
	payload := conversions.WashServerListToRest(res)

	switch {
	case err == nil:
		return wash_servers.NewListOK().WithPayload(payload)
	case errors.Is(err, entity.ErrNotFound):
		return wash_servers.NewListNotFound()
	case errors.Is(err, entity.ErrAccessDenied):
		return wash_servers.NewListForbidden()
	default:
		return wash_servers.NewListInternalServerError()
	}
}
