package rest

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"wash_admin/internal/app"
	"wash_admin/internal/conversions"
	"wash_admin/internal/entity"
	"wash_admin/openapi/restapi/operations"
	"wash_admin/openapi/restapi/operations/wash_servers"
)

func (svc *service) initWashServerHandlers(api *operations.WashAdminAPI) {
	api.WashServersGetHandler = wash_servers.GetHandlerFunc(svc.getWashServer)
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
