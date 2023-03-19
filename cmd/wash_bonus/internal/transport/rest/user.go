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
	api.UserGetProfileHandler = user.GetProfileHandlerFunc(svc.getProfile)
	api.UserGetBalanceHandler = user.GetBalanceHandlerFunc(svc.getBalance)
}

func (svc *service) getProfile(params user.GetProfileParams, auth *app.Auth) user.GetProfileResponder {
	res, err := svc.userSvc.GetByID(params.HTTPRequest.Context(), auth.UID)

	payload := conversions.UserToRest(res)

	switch {
	case err == nil:
		return user.NewGetProfileOK().WithPayload(&payload)
	case errors.Is(err, entity.ErrNotFound):
		return user.NewGetProfileNotFound()
	default:
		return user.NewGetProfileInternalServerError()
	}
}

func (svc *service) getBalance(params user.GetBalanceParams, auth *app.Auth) user.GetBalanceResponder {
	res, err := svc.userSvc.GetByID(params.HTTPRequest.Context(), auth.UID)
	switch {
	case err == nil:
		return user.NewGetBalanceOK().WithPayload(&user.GetBalanceOKBody{Balance: res.Balance.IntPart()})
	case errors.Is(err, entity.ErrNotFound):
		return user.NewGetBalanceNotFound()
	default:
		return user.NewGetBalanceInternalServerError()
	}
}
