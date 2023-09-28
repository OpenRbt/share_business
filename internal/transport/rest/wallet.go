package rest

import (
	"fmt"
	"washbonus/internal/app"
	"washbonus/internal/conversions"
	"washbonus/internal/entities"
	"washbonus/openapi/bonus/restapi/operations"
	"washbonus/openapi/bonus/restapi/operations/wallets"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) initWalletHandlers(api *operations.WashBonusAPI) {
	api.WalletsGetWalletsHandler = wallets.GetWalletsHandlerFunc(svc.getWallets)
	api.WalletsGetWalletByOrganizationIDHandler = wallets.GetWalletByOrganizationIDHandlerFunc(svc.getWalletByOrganizationID)
}

func (svc *service) getWallets(params wallets.GetWalletsParams, auth *app.Auth) wallets.GetWalletsResponder {
	op := "Get wallets:"
	resp := wallets.NewGetWalletsDefault(500)

	pagination := conversions.PaginationFromRest(*params.Limit, *params.Offset)
	res, err := svc.walletCtrl.Get(params.HTTPRequest.Context(), *auth, pagination)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return wallets.NewGetWalletsOK().WithPayload(conversions.WalletsToRest(res))
}

func (svc *service) getWalletByOrganizationID(params wallets.GetWalletByOrganizationIDParams, auth *app.Auth) wallets.GetWalletByOrganizationIDResponder {
	op := "Get wallet by organization ID:"
	resp := wallets.NewGetWalletByOrganizationIDDefault(500)

	organizationID, err := uuid.FromString(params.ID.String())
	if err != nil {
		setAPIError(svc.l, op, fmt.Errorf("Wrong organization ID: %w", entities.ErrBadRequest), resp)
		return resp
	}

	res, err := svc.walletCtrl.GetByOrganizationId(params.HTTPRequest.Context(), *auth, organizationID)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return wallets.NewGetWalletByOrganizationIDOK().WithPayload(conversions.WalletToRest(res))
}
