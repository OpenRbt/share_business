package rest

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"wash_bonus/internal/app"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/entity"
	"wash_bonus/openapi/models"
	"wash_bonus/openapi/restapi/operations"
	"wash_bonus/openapi/restapi/operations/session"
)

func (svc *service) initSessionHandlers(api *operations.WashBonusAPI) {
	api.SessionGetSessionHandler = session.GetSessionHandlerFunc(svc.getSession)
	api.SessionPostSessionHandler = session.PostSessionHandlerFunc(svc.chargeBonuses)
}

func (svc *service) getSession(params session.GetSessionParams, auth *app.Auth) session.GetSessionResponder {
	var payload *models.Session

	sessionUID, err := uuid.FromString(params.UID)
	if err != nil {
		return session.NewGetSessionInternalServerError()
	}
	res, err := svc.sessionSvc.GetUserSession(params.HTTPRequest.Context(), auth, sessionUID)

	if err == nil {
		payload = conversions.SessionToRest(res)
	}

	switch {
	case err == nil:
		return session.NewGetSessionOK().WithPayload(payload)
	case errors.Is(err, entity.ErrNotFound):
		return session.NewGetSessionNotFound()
	default:
		return session.NewGetSessionInternalServerError()
	}
}

func (svc *service) chargeBonuses(params session.PostSessionParams, auth *app.Auth) session.PostSessionResponder {
	sessionID, err := uuid.FromString(params.UID)
	if err != nil {
		return session.NewPostSessionInternalServerError()
	}

	amountD := decimal.NewFromInt(params.Body.Amount)

	err = svc.sessionSvc.ChargeBonuses(params.HTTPRequest.Context(), auth, sessionID, amountD)

	switch {

	case err == nil:
		return session.NewPostSessionNoContent()
	case errors.Is(err, entity.ErrNotFound):
		return session.NewPostSessionNotFound()
	case errors.Is(err, entity.ErrForbidden):
		return session.NewPostSessionForbidden()
	default:
		return session.NewPostSessionInternalServerError()
	}

}
