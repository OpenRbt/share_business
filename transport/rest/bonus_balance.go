package api

import (
	"errors"
	"wash-bonus/internal/def"

	"github.com/go-openapi/swag"

	"wash-bonus/internal/app"
	"wash-bonus/internal/dto"
	"wash-bonus/transport/rest/restapi/models"
	"wash-bonus/transport/rest/restapi/restapi/operations"
	"wash-bonus/transport/rest/restapi/restapi/operations/bonus_balance"
	balance "wash-bonus/transport/rest/restapi/restapi/operations/bonus_balance"

	"strconv"

	"github.com/go-openapi/runtime/middleware"
)

func setBonusBalanceHandlers(api *operations.WashBonusAPI, svc *service) {
	api.BonusBalanceGetBonusBalanceHandler = bonus_balance.GetBonusBalanceHandlerFunc(svc.GetBonusBalance)
	api.BonusBalanceAddBonusBalanceHandler = bonus_balance.AddBonusBalanceHandlerFunc(svc.AddBonusBalance)
	api.BonusBalanceEditBonusBalanceHandler = bonus_balance.EditBonusBalanceHandlerFunc(svc.EditBonusBalance)
	api.BonusBalanceDeleteBonusBalanceHandler = bonus_balance.DeleteBonusBalanceHandlerFunc(svc.DeleteBonusBalance)
}

func (svc *service) GetBonusBalance(params balance.GetBonusBalanceParams, profile interface{}) middleware.Responder {
	c, err := svc.bonusSvc.GetBonusBalance(params.Body.ID)
	switch {
	default:
		log.PrintErr("GetBonusBalance server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return balance.NewGetBonusBalanceDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("GeBonusBalance client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return balance.NewGetBonusBalanceDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("GetBonusBalance client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return balance.NewGetBonusBalanceDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("GetBonusBalance ok", "id", params.Body.ID)
		return balance.NewGetBonusBalanceOK().WithPayload(dto.ApiBonusBalance(c))
	}
}

func (svc *service) AddBonusBalance(params balance.AddBonusBalanceParams, profile interface{}) middleware.Responder {
	balanc, _ := strconv.ParseFloat(params.Body.Balance, 8)
	c, err := svc.bonusSvc.AddBonusBalance(params.Body.UserID, balanc)
	switch {
	default:
		log.PrintErr("AddBonusBalance server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return balance.NewAddBonusBalanceDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("AddBonusBalance client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return balance.NewAddBonusBalanceDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("AddBonusBalance ok")
		return balance.NewAddBonusBalanceCreated().WithPayload(dto.ApiBonusBalance(c))
	}
}

func (svc *service) EditBonusBalance(params balance.EditBonusBalanceParams, profile interface{}) middleware.Responder {
	balanc, _ := strconv.ParseFloat(params.Body.Data.Balance, 8)
	err := svc.bonusSvc.EditBonusBalance(params.Body.ID, balanc)
	switch {
	default:
		log.PrintErr("EditBonusBalance server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return balance.NewEditBonusBalanceDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("EditBonusBalance client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return balance.NewEditBonusBalanceDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("EditBonusBalance client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return balance.NewEditBonusBalanceDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("EditBonusBalance ok")
		return balance.NewEditBonusBalanceOK()
	}
}

func (svc *service) DeleteBonusBalance(params balance.DeleteBonusBalanceParams, profile interface{}) middleware.Responder {
	err := svc.bonusSvc.DeleteBonusBalance(params.Body.ID, params.Body.UserID)
	switch {
	default:
		log.PrintErr("DeleteBonusBalance server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return balance.NewDeleteBonusBalanceDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("DeleteBonusBalance client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return balance.NewDeleteBonusBalanceDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("DeleteBonusBalance client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return balance.NewDeleteBonusBalanceDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("DeleteBonusBalance ok", "id", params.Body.ID)
		return balance.NewDeleteBonusBalanceNoContent()
	}
}
