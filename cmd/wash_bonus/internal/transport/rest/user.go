package rest

import (
	"errors"
	"wash_bonus/internal/app"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/entity"
	"wash_bonus/openapi/restapi/operations"
	"wash_bonus/openapi/restapi/operations/user"
)

func (svc *service) initUserHandlers(api *operations.WashBonusAPI) {
	api.UserGetHandler = user.GetHandlerFunc(svc.getProfile)
}

func (svc *service) getProfile(params user.GetParams, auth *app.Auth) user.GetResponder {
	res, err := svc.userSvc.Get(params.HTTPRequest.Context(), auth)

	payload := conversions.UserToRest(res)

	switch {
	case err == nil:
		return user.NewGetOK().WithPayload(&payload)
	case errors.Is(err, entity.ErrNotFound):
		return user.NewGetNotFound()
	default:
		return user.NewGetInternalServerError()
	}
}
