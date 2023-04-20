package specification

import (
	"mong-specification/example/filters"
	"mong-specification/example/sort"
)

type UserSpecification interface {
	Set
	By(filter filters.UserBy) UserSpecification
	Filter(filter *filters.UserFilter) UserSpecification
	SortBy(sort *sort.UserSort) UserSpecification
}
