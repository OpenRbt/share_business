package rest

import (
	"errors"
	"washBonus/internal/app"
	"washBonus/internal/conversions"
	"washBonus/internal/entity"
	"washBonus/openapi/models"
	"washBonus/openapi/restapi/operations"
	"washBonus/openapi/restapi/operations/sessions"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (svc *service) initSessionHandlers(api *operations.WashBonusAPI) {
	api.SessionsGetSessionByIDHandler = sessions.GetSessionByIDHandlerFunc(svc.getSession)
	api.SessionsChargeBonusesOnSessionHandler = sessions.ChargeBonusesOnSessionHandlerFunc(svc.chargeBonuses)
	api.SessionsAssignUserToSessionHandler = sessions.AssignUserToSessionHandlerFunc(svc.assignUserToSession)
}

func (svc *service) getSession(params sessions.GetSessionByIDParams, auth *app.Auth) sessions.GetSessionByIDResponder {
	var payload *models.Session

	sessionID, err := uuid.FromString(params.ID)
	if err != nil {
		return sessions.NewGetSessionByIDInternalServerError()
	}

	res, err := svc.sessionCtrl.GetSession(params.HTTPRequest.Context(), sessionID, auth.UID)

	switch {
	case err == nil:
		payload = conversions.SessionToRest(res)
		return sessions.NewGetSessionByIDOK().WithPayload(payload)
	case errors.Is(err, entity.ErrForbidden):
		return sessions.NewGetSessionByIDForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return sessions.NewGetSessionByIDNotFound()
	default:
		return sessions.NewGetSessionByIDInternalServerError()
	}
}

func (svc *service) chargeBonuses(params sessions.ChargeBonusesOnSessionParams, auth *app.Auth) sessions.ChargeBonusesOnSessionResponder {
	sessionID, err := uuid.FromString(params.ID)
	if err != nil {
		return sessions.NewChargeBonusesOnSessionInternalServerError()
	}

	amount := decimal.NewFromInt(params.Body.Amount)
	err = svc.sessionCtrl.ChargeBonuses(params.HTTPRequest.Context(), amount, sessionID, auth.UID)

	switch {

	case err == nil:
		return sessions.NewChargeBonusesOnSessionOK().WithPayload(&models.BonusCharge{Amount: params.Body.Amount})
	case errors.Is(err, entity.ErrNotFound):
		return sessions.NewChargeBonusesOnSessionNotFound()
	case errors.Is(err, entity.ErrForbidden):
		return sessions.NewChargeBonusesOnSessionForbidden()
	default:
		return sessions.NewChargeBonusesOnSessionInternalServerError()
	}

}

func (svc *service) assignUserToSession(params sessions.AssignUserToSessionParams, auth *app.Auth) sessions.AssignUserToSessionResponder {
	sessionID, err := uuid.FromString(params.ID)
	if err != nil {
		return sessions.NewAssignUserToSessionInternalServerError()
	}

	err = svc.sessionCtrl.AssignUserToSession(params.HTTPRequest.Context(), sessionID, auth.UID)

	switch {

	case err == nil:
		return sessions.NewAssignUserToSessionNoContent()
	case errors.Is(err, entity.ErrNotFound):
		return sessions.NewAssignUserToSessionNotFound()
	case errors.Is(err, entity.ErrForbidden):
		return sessions.NewAssignUserToSessionForbidden()
	default:
		return sessions.NewAssignUserToSessionInternalServerError()
	}

}
