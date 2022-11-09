package app

type ListParams struct {
	Offset       int64
	Limit        int64
	FilterGroups []*FilterGroup
	SortBy       string
	OrderBy      string
}

type FilterGroup struct {
	Key         string
	LogicFilter bool
	Filters     []*Filter
}

type Filter struct {
	Value      string
	Operator   string
	IgnoreCase bool
}
