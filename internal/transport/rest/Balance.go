package api

import (
	"errors"
	"wash-bonus/internal/def"
	"wash-bonus/openapi/models"
	"wash-bonus/openapi/restapi/operations"
	"wash-bonus/openapi/restapi/operations/balance"

	"github.com/go-openapi/swag"

	"strconv"
	"wash-bonus/internal/app"
	"wash-bonus/internal/dto"

	"github.com/go-openapi/runtime/middleware"
)

func setBalanceHandlers(api *operations.WashBonusAPI, svc *service) {
	api.BalanceGetBalanceHandler = balance.GetBalanceHandlerFunc(svc.GetBalance)
	api.BalanceAddBalanceHandler = balance.AddBalanceHandlerFunc(svc.AddBalance)
	api.BalanceEditBalanceHandler = balance.EditBalanceHandlerFunc(svc.EditBalance)
	api.BalanceDeleteBalanceHandler = balance.DeleteBalanceHandlerFunc(svc.DeleteBalance)
}

func (svc *service) GetBalance(params balance.GetBalanceParams, profile interface{}) middleware.Responder {
	c, err := svc.bonusSvc.GetBalance(params.Body.ID)
	switch {
	default:
		log.PrintErr("GetBalance server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return balance.NewGetBalanceDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("GeBalance client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return balance.NewGetBalanceDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("GetBalance client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return balance.NewGetBalanceDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("GetBalance ok", "id", params.Body.ID)
		return balance.NewGetBalanceOK().WithPayload(dto.BalanceToRest(c))
	}
}

func (svc *service) AddBalance(params balance.AddBalanceParams, profile interface{}) middleware.Responder {
	balanc, _ := strconv.ParseFloat(params.Body.Balance, 8)
	c, err := svc.bonusSvc.AddBalance(params.Body.UserID, balanc)
	switch {
	default:
		log.PrintErr("AddBalance server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return balance.NewAddBalanceDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("AddBalance client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return balance.NewAddBalanceDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("AddBalance ok")
		return balance.NewAddBalanceCreated().WithPayload(dto.BalanceToRest(c))
	}
}

func (svc *service) EditBalance(params balance.EditBalanceParams, profile interface{}) middleware.Responder {
	balanc, _ := strconv.ParseFloat(params.Body.Data.Balance, 8)
	err := svc.bonusSvc.EditBalance(params.Body.ID, balanc)
	switch {
	default:
		log.PrintErr("EditBalance server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return balance.NewEditBalanceDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("EditBalance client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return balance.NewEditBalanceDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("EditBalance client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return balance.NewEditBalanceDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("EditBalance ok")
		return balance.NewEditBalanceOK()
	}
}

func (svc *service) DeleteBalance(params balance.DeleteBalanceParams, profile interface{}) middleware.Responder {
	err := svc.bonusSvc.DeleteBalance(params.Body.ID, params.Body.UserID)
	switch {
	default:
		log.PrintErr("DeleteBalance server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return balance.NewDeleteBalanceDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("DeleteBalance client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return balance.NewDeleteBalanceDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("DeleteBalance client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return balance.NewDeleteBalanceDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("DeleteBalance ok", "id", params.Body.ID)
		return balance.NewDeleteBalanceNoContent()
	}
}
