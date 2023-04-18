package specification

import (
	"context"
	"mong-specification/infrastructure/mongo/sort"
)

type Specification interface {
	Query() interface{}
}

type BaseSpecification struct {
}

type Set interface {
	Specification
	WithContext(ctx context.Context) Set
	FilterByID(id string) Set
	FilterEntry(attributes map[string]interface{}) Set
	Prefetch(preload ...string) Set
	Sort(sortBy string, sortOrder sort.Order) Set
	Limit(limit int) Set
	LessThan(field string, value interface{}) Set
	GreaterThan(field string, value interface{}) Set
}
