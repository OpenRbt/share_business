package dto

import (
	"wash-bonus/internal/app/entity/vo"
	models2 "wash-bonus/internal/transport/rest/restapi/models"
)

func ListFilterFromRest(lf *models2.ListParams) vo.ListFilter {
	var offset int64 = 0
	if lf.Offset != nil {
		offset = *lf.Offset
	}
	return vo.ListFilter{
		Offset:       offset,
		Limit:        lf.Limit,
		FilterGroups: filterGroupsFromRest(lf.FilterGroups),
		SortBy:       lf.SortBy,
		OrderBy:      lf.OrderBy,
	}
}

func filterGroupsFromRest(ff []*models2.FilterGroup) []vo.FilterGroup {
	res := make([]vo.FilterGroup, len(ff))

	for i, f := range ff {
		res[i] = filterGroupFromRest(*f)
	}

	return res
}

func filterGroupFromRest(f models2.FilterGroup) vo.FilterGroup {
	return vo.FilterGroup{
		Key:         f.Key,
		LogicFilter: f.LogicFilter,
		Filters:     filtersFromRest(f.Filters),
	}
}
func filtersFromRest(ff []*models2.Filter) []vo.Filter {
	res := make([]vo.Filter, len(ff))

	for i, f := range ff {
		res[i] = vo.Filter{
			Column:   f.Value,
			Operator: f.Operator,
		}
	}

	return res
}
