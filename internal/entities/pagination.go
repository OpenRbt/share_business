package entities

import "math"

type (
	Pagination struct {
		Limit  int64
		Offset int64
	}

	Page[T any] struct {
		Items      []T
		Page       int64
		PageSize   int64
		TotalPages int64
		TotalItems int64
	}

	Filter struct {
		page     int64
		pageSize int64
	}
)

func NewPage[T any](items []T, filter Filter, totalItems int64) Page[T] {
	return Page[T]{
		Items:      items,
		TotalPages: int64(math.Ceil((float64(totalItems) / float64(filter.pageSize)))),
		Page:       filter.page,
		PageSize:   filter.pageSize,
		TotalItems: totalItems,
	}
}

func NewFilter(page int64, pageSize int64) Filter {
	filter := Filter{
		page:     1,
		pageSize: 10,
	}

	if page > 1 {
		filter.page = page
	}
	if pageSize >= 1 && pageSize <= 100 {
		filter.pageSize = pageSize
	} else if pageSize > 100 {
		filter.pageSize = 100
	}

	return filter
}

func (f Filter) Page() int64 {
	return f.page
}

func (f Filter) PageSize() int64 {
	return f.pageSize
}
