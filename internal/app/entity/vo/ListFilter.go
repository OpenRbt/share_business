package vo

type ListFilter struct {
	Offset int64
	Limit  int64

	FilterGroups []FilterGroup

	SortBy  string
	OrderBy string
}
