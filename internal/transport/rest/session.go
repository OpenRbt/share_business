package rest

import (
	"fmt"
	"washbonus/internal/app"
	"washbonus/internal/conversions"
	"washbonus/internal/entities"
	"washbonus/openapi/bonus/models"
	"washbonus/openapi/bonus/restapi/operations"
	"washbonus/openapi/bonus/restapi/operations/sessions"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (svc *service) initSessionHandlers(api *operations.WashBonusAPI) {
	api.SessionsGetSessionByIDHandler = sessions.GetSessionByIDHandlerFunc(svc.getSession)
	api.SessionsChargeBonusesOnSessionHandler = sessions.ChargeBonusesOnSessionHandlerFunc(svc.chargeBonuses)
	api.SessionsAssignUserToSessionHandler = sessions.AssignUserToSessionHandlerFunc(svc.assignUserToSession)
}

func (svc *service) getSession(params sessions.GetSessionByIDParams, auth *app.Auth) sessions.GetSessionByIDResponder {
	op := "Get session:"
	resp := sessions.NewGetSessionByIDDefault(500)

	sessionID, err := uuid.FromString(params.SessionID)
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong session ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	res, err := svc.sessionCtrl.GetSession(params.HTTPRequest.Context(), *auth, sessionID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return sessions.NewGetSessionByIDOK().WithPayload(conversions.SessionToRest(res))
}

func (svc *service) chargeBonuses(params sessions.ChargeBonusesOnSessionParams, auth *app.Auth) sessions.ChargeBonusesOnSessionResponder {
	op := "Charge bonuses:"
	resp := sessions.NewChargeBonusesOnSessionDefault(500)

	sessionID, err := uuid.FromString(params.SessionID)
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong session ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	amount := decimal.NewFromInt(params.Body.Amount)
	err = svc.sessionCtrl.ChargeBonuses(params.HTTPRequest.Context(), *auth, amount, sessionID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return sessions.NewChargeBonusesOnSessionOK().WithPayload(&models.BonusCharge{Amount: params.Body.Amount})
}

func (svc *service) assignUserToSession(params sessions.AssignUserToSessionParams, auth *app.Auth) sessions.AssignUserToSessionResponder {
	op := "Assign user to session:"
	resp := sessions.NewAssignUserToSessionDefault(500)

	sessionID, err := uuid.FromString(params.SessionID)
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong session ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	err = svc.sessionCtrl.AssignUserToSession(params.HTTPRequest.Context(), *auth, sessionID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return sessions.NewAssignUserToSessionNoContent()
}
