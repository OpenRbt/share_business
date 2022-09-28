package api

import (
	"errors"
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/def"

	"github.com/go-openapi/swag"

	"wash-bonus/internal/api/restapi/models"
	balance "wash-bonus/internal/api/restapi/restapi/operations/bonus_balance"
	"wash-bonus/internal/app"

	"strconv"

	"github.com/go-openapi/runtime/middleware"
)

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
		return balance.NewGetBonusBalanceOK().WithPayload(apiBonusBalance(c))
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
		return balance.NewAddBonusBalanceCreated().WithPayload(apiBonusBalance(c))
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

func apiBonusBalance(a *entity.BonusBalance) *models.Balance {
	if a == nil {
		return nil
	}
	return &models.Balance{
		ID:      a.ID,
		UserID:  a.UserId,
		Balance: strconv.FormatFloat(a.Balance, 'f', 6, 64),
	}
}
