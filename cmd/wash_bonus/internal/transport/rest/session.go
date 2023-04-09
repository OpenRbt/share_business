package rest

import (
	"errors"
	"wash_bonus/internal/app"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/entity"
	"wash_bonus/openapi/models"
	"wash_bonus/openapi/restapi/operations"
	"wash_bonus/openapi/restapi/operations/session"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (svc *service) initSessionHandlers(api *operations.WashBonusAPI) {
	api.SessionGetSessionHandler = session.GetSessionHandlerFunc(svc.getSession)
	api.SessionPostSessionHandler = session.PostSessionHandlerFunc(svc.chargeBonuses)
	api.SessionAssignUserToSessionHandler = session.AssignUserToSessionHandlerFunc(svc.assignUserToSession)
}

func (svc *service) getSession(params session.GetSessionParams, auth *app.Auth) session.GetSessionResponder {
	var payload *models.Session

	sessionUID, err := uuid.FromString(params.UID)
	if err != nil {
		return session.NewGetSessionInternalServerError()
	}

	res, err := svc.sessionUseCase.Get(params.HTTPRequest.Context(), sessionUID, auth.UID)

	switch {
	case err == nil:
		payload = conversions.SessionToRest(res)
		return session.NewGetSessionOK().WithPayload(payload)
	case errors.Is(err, entity.ErrForbidden):
		return session.NewGetSessionForbidden()
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

	err = svc.sessionUseCase.ChargeBonuses(params.HTTPRequest.Context(), sessionID, auth.UID, amountD)

	switch {

	case err == nil:
		return session.NewPostSessionOK().WithPayload(&models.BonusCharge{Amount: params.Body.Amount})
	case errors.Is(err, entity.ErrNotFound):
		return session.NewPostSessionNotFound()
	case errors.Is(err, entity.ErrForbidden):
		return session.NewPostSessionForbidden()
	default:
		return session.NewPostSessionInternalServerError()
	}

}

func (svc *service) assignUserToSession(params session.AssignUserToSessionParams, auth *app.Auth) session.AssignUserToSessionResponder {
	sessionID, err := uuid.FromString(params.SessionID)
	if err != nil {
		return session.NewAssignUserToSessionInternalServerError()
	}

	err = svc.sessionUseCase.AssignUser(params.HTTPRequest.Context(), sessionID, auth.UID)

	switch {

	case err == nil:
		return session.NewAssignUserToSessionNoContent()
	case errors.Is(err, entity.ErrNotFound):
		return session.NewAssignUserToSessionNotFound()
	case errors.Is(err, entity.ErrForbidden):
		return session.NewAssignUserToSessionForbidden()
	default:
		return session.NewAssignUserToSessionInternalServerError()
	}

}
