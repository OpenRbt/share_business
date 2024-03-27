package rest

import (
	"washbonus/internal/app"
	"washbonus/internal/conversions"
	"washbonus/openapi/admin/restapi/operations"
	"washbonus/openapi/admin/restapi/operations/reports"
)

func (svc *service) initBonusReportHandlers(api *operations.WashAdminAPI) {
	api.ReportsGetBonusReportsHandler = reports.GetBonusReportsHandlerFunc(svc.getBonusReports)
}

func (svc *service) getBonusReports(params reports.GetBonusReportsParams, auth *app.AdminAuth) reports.GetBonusReportsResponder {
	op := "Get server groups:"
	resp := reports.NewGetBonusReportsDefault(500)

	filter, err := conversions.BonusReportFilterFromRest(params)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	list, err := svc.bonusReportCtrl.List(params.HTTPRequest.Context(), *auth, filter)
	if err != nil {
		setAPIError(svc.l, op, err, resp)
		return resp
	}

	return reports.NewGetBonusReportsOK().WithPayload(conversions.BonusReportPageToRest(list))
}
