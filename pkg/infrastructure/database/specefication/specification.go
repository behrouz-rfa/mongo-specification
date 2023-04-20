package specification

import (
	"context"
	"github.com/behrouz-rfa/mongo-specification/pkg/sort"
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
