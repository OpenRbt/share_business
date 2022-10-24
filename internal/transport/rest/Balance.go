package api

import (
	"errors"
	"wash-bonus/internal/def"
	"wash-bonus/internal/transport/rest/restapi/models"
	"wash-bonus/internal/transport/rest/restapi/restapi/operations"
	balance2 "wash-bonus/internal/transport/rest/restapi/restapi/operations/balance"

	"github.com/go-openapi/swag"

	"strconv"
	"wash-bonus/internal/app"
	"wash-bonus/internal/dto"

	"github.com/go-openapi/runtime/middleware"
)

func setBalanceHandlers(api *operations.WashBonusAPI, svc *service) {
	api.BalanceGetBalanceHandler = balance2.GetBalanceHandlerFunc(svc.GetBalance)
	api.BalanceAddBalanceHandler = balance2.AddBalanceHandlerFunc(svc.AddBalance)
	api.BalanceEditBalanceHandler = balance2.EditBalanceHandlerFunc(svc.EditBalance)
	api.BalanceDeleteBalanceHandler = balance2.DeleteBalanceHandlerFunc(svc.DeleteBalance)
}

func (svc *service) GetBalance(params balance2.GetBalanceParams, profile interface{}) middleware.Responder {
	c, err := svc.bonusSvc.GetBalance(params.Body.ID)
	switch {
	default:
		log.PrintErr("GetBalance server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return balance2.NewGetBalanceDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("GeBalance client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return balance2.NewGetBalanceDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("GetBalance client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return balance2.NewGetBalanceDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("GetBalance ok", "id", params.Body.ID)
		return balance2.NewGetBalanceOK().WithPayload(dto.BalanceToRest(c))
	}
}

func (svc *service) AddBalance(params balance2.AddBalanceParams, profile interface{}) middleware.Responder {
	balanc, _ := strconv.ParseFloat(params.Body.Balance, 8)
	c, err := svc.bonusSvc.AddBalance(params.Body.UserID, balanc)
	switch {
	default:
		log.PrintErr("AddBalance server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return balance2.NewAddBalanceDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("AddBalance client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return balance2.NewAddBalanceDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("AddBalance ok")
		return balance2.NewAddBalanceCreated().WithPayload(dto.BalanceToRest(c))
	}
}

func (svc *service) EditBalance(params balance2.EditBalanceParams, profile interface{}) middleware.Responder {
	balanc, _ := strconv.ParseFloat(params.Body.Data.Balance, 8)
	err := svc.bonusSvc.EditBalance(params.Body.ID, balanc)
	switch {
	default:
		log.PrintErr("EditBalance server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return balance2.NewEditBalanceDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("EditBalance client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return balance2.NewEditBalanceDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("EditBalance client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return balance2.NewEditBalanceDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("EditBalance ok")
		return balance2.NewEditBalanceOK()
	}
}

func (svc *service) DeleteBalance(params balance2.DeleteBalanceParams, profile interface{}) middleware.Responder {
	err := svc.bonusSvc.DeleteBalance(params.Body.ID, params.Body.UserID)
	switch {
	default:
		log.PrintErr("DeleteBalance server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return balance2.NewDeleteBalanceDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("DeleteBalance client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return balance2.NewDeleteBalanceDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("DeleteBalance client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return balance2.NewDeleteBalanceDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("DeleteBalance ok", "id", params.Body.ID)
		return balance2.NewDeleteBalanceNoContent()
	}
}
