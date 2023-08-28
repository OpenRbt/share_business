package rest

import (
	"errors"
	"washBonus/internal/app"
	"washBonus/internal/conversions"
	"washBonus/internal/entity"
	"washBonus/openapi/restapi/operations"
	wallets "washBonus/openapi/restapi/operations/wallets"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initWalletHandlers(api *operations.WashBonusAPI) {
	api.WalletsGetWalletsHandler = wallets.GetWalletsHandlerFunc(svc.getWallets)
	api.WalletsGetWalletByOrganizationIDHandler = wallets.GetWalletByOrganizationIDHandlerFunc(svc.getWalletByOrganizationID)
}

func (svc *service) getWallets(params wallets.GetWalletsParams, auth *app.Auth) wallets.GetWalletsResponder {
	pagination := conversions.PaginationFromRest(*params.Body)

	res, err := svc.walletCtrl.Get(params.HTTPRequest.Context(), auth.User, pagination)

	switch {
	case err == nil:
		return wallets.NewGetWalletsOK().WithPayload(conversions.WalletsToRest(res))
	case errors.Is(err, entity.ErrNotFound):
		return wallets.NewGetWalletsNotFound()
	case errors.Is(err, entity.ErrAccessDenied):
		return wallets.NewGetWalletsForbidden()
	default:
		svc.l.Errorln("Get wallets:", err)
		return wallets.NewGetWalletsInternalServerError()
	}
}

func (svc *service) getWalletByOrganizationID(params wallets.GetWalletByOrganizationIDParams, auth *app.Auth) wallets.GetWalletByOrganizationIDResponder {
	organizationID, err := uuid.FromString(params.ID.String())
	if err != nil {
		return wallets.NewGetWalletByOrganizationIDBadRequest()
	}

	res, err := svc.walletCtrl.GetByOrganizationId(params.HTTPRequest.Context(), auth.User, organizationID)

	switch {
	case err == nil:
		return wallets.NewGetWalletByOrganizationIDOK().WithPayload(conversions.WalletToRest(res))
	case errors.Is(err, entity.ErrAccessDenied):
		return wallets.NewGetWalletByOrganizationIDForbidden()
	case errors.Is(err, entity.ErrNotFound):
		return wallets.NewGetWalletByOrganizationIDNotFound()
	default:
		svc.l.Errorln("Get wallet by organization id:", err)
		return wallets.NewGetWalletByOrganizationIDInternalServerError()
	}
}
