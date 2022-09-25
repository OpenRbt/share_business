package dto

import (
	"wash-bonus/internal/api/restapi/models"
	"wash-bonus/internal/app/entity/vo"
)

func ListFilterFromRest(lf *models.ListParams) vo.ListFilter {
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

func filterGroupsFromRest(ff []*models.FilterGroup) []vo.FilterGroup {
	res := make([]vo.FilterGroup, len(ff))

	for i, f := range ff {
		res[i] = filterGroupFromRest(*f)
	}

	return res
}

func filterGroupFromRest(f models.FilterGroup) vo.FilterGroup {
	return vo.FilterGroup{
		Key:         f.Key,
		LogicFilter: f.LogicFilter,
		Filters:     filtersFromRest(f.Filters),
	}
}
func filtersFromRest(ff []*models.Filter) []vo.Filter {
	res := make([]vo.Filter, len(ff))

	for i, f := range ff {
		res[i] = vo.Filter{
			Column:   f.Value,
			Operator: f.Operator,
		}
	}

	return res
}
