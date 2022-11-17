// Code generated by mtgroup-generator.
package api

import (
	"strings"
	"time"
	models2 "wash-bonus/internal/transport/rest/models"

	"github.com/go-openapi/strfmt"
	"wash-bonus/internal/app"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func appListParams(apiLP *models2.ListParams) *app.ListParams {
	var offset int64 = 0
	if apiLP.Offset != nil {
		offset = *apiLP.Offset
	}
	return &app.ListParams{
		Offset:       offset,
		Limit:        apiLP.Limit,
		FilterGroups: appFilterGroups(apiLP.FilterGroups),
		SortBy:       apiLP.SortBy,
		OrderBy:      apiLP.OrderBy,
	}
}

func appFilterGroups(apiFG []*models2.FilterGroup) []*app.FilterGroup {
	appFG := []*app.FilterGroup{}
	for _, fg := range apiFG {
		appFG = append(appFG, &app.FilterGroup{
			Key:         fg.Key,
			LogicFilter: fg.LogicFilter,
			Filters:     appFilters(fg.Filters),
		})
	}
	return appFG
}

func appFilters(apiFP []*models2.Filter) []*app.Filter {
	appF := []*app.Filter{}
	for _, fp := range apiFP {
		appF = append(appF, &app.Filter{
			Value:      fp.Value,
			Operator:   fp.Operator,
			IgnoreCase: fp.IgnoreCase,
		})
	}
	return appF
}

func fromDateTimesArray(dts []*strfmt.DateTime) (dates []*time.Time) {
	for _, date := range dts {
		dates = append(dates, (*time.Time)(date))
	}
	return
}

func toDateTimesArray(dates []*time.Time) (dts []*strfmt.DateTime) {
	for _, date := range dates {
		dts = append(dts, (*strfmt.DateTime)(date))
	}
	return
}

func fromDatesArray(ds []*strfmt.Date) (dates []*time.Time) {
	for _, date := range ds {
		dates = append(dates, (*time.Time)(date))
	}
	return
}

func toDatesArray(dates []*time.Time) (ds []*strfmt.Date) {
	for _, date := range dates {
		ds = append(ds, (*strfmt.Date)(date))
	}
	return
}

func splitCommaSeparatedStr(commaSeparated string) (result []string) {
	for _, item := range strings.Split(commaSeparated, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			result = append(result, item)
		}
	}
	return
}
