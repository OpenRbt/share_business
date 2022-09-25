package vo

type FilterGroup struct {
	Key         string
	LogicFilter bool
	Filters     []Filter
}
