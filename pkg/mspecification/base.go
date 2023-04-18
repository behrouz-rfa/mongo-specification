package mspecification

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	specification "mong-specification/pkg/infrastructure/database/specefication"
	"mong-specification/pkg/sort"

	"mong-specification/pkg/utils"
)

type ExtendedSpecifications interface {
	specification.Set
	PreQueryHook()
	NeedsPreload(field string) (bool, []bson.M)
}

type BsonO interface {
	bson.M | bson.A | bson.D
}

// BaseSpecification is the base specification for all mongo specifications.
type BaseSpecification struct {
	aggregationPipeline []bson.M
	preloadIndex        map[string]int
	filters             []bson.M
	postFilters         []bson.M

	limit  int
	sort   bson.D
	parent ExtendedSpecifications
}

// NewBaseSpecification creates a new base specification.
func NewBaseSpecification() *BaseSpecification {
	spec := new(BaseSpecification)
	spec.parent = spec

	return spec
}

// Query returns the query to be used in the database
// Order of operations:
// 1. PreQueryHook
// 2. Filters
// 3. Preload
// 4. PostFilters
// 5. Sort
// 6. Limit.
func (b *BaseSpecification) Query() interface{} {
	// Call pre hooks before anything else
	b.parent.PreQueryHook()

	// Add filters to the query pipeline
	query := make([]bson.M, 0)

	if len(b.filters) > 0 {
		query = append(query, bson.M{"$match": bson.M{"$and": b.filters}})
	}

	// Add the aggregation pipeline to the query
	query = append(query, b.aggregationPipeline...)

	// Add post filters to the query pipeline
	if len(b.postFilters) > 0 {
		query = append(query, bson.M{"$match": bson.M{"$and": b.postFilters}})
	}

	// Add sort to the query pipeline
	// TODO: make sure sort is compatible with the aggregation pipeline

	if b.limit != 1 && b.sort != nil {
		query = append(query, bson.M{"$sort": b.sort})
	}

	return query
}

// WithContext adds a context to the specification.
func (b *BaseSpecification) WithContext(ctx context.Context) specification.Set {
	b.Prefetch(utils.LoadersFromContext(ctx)...)
	return b.parent
}

// FilterByID adds a filter to the specification by ID (primary key).
func (b *BaseSpecification) FilterByID(id string) specification.Set {
	b.filters = append(b.filters, bson.M{"_id": id})
	return b.parent
}

// Filter adds filters to the specification.
func (b *BaseSpecification) FilterEntry(attributes map[string]interface{}) specification.Set {
	for key, value := range attributes {
		if needsPreload, preloader := b.parent.NeedsPreload(key); needsPreload {
			b.aggregationPipeline = append(b.aggregationPipeline, preloader...)
			b.postFilters = append(b.postFilters, bson.M{key: value})
		} else {
			b.filters = append(b.filters, bson.M{key: value})
		}
	}

	return b.parent
}

// Prefetch adds preloads to the specification.
func (b *BaseSpecification) Prefetch(preload ...string) specification.Set {
	for _, preload := range preload {
		if needsPreload, preloader := b.parent.NeedsPreload(preload); needsPreload {
			b.aggregationPipeline = append(b.aggregationPipeline, preloader...)
		}
	}

	return b.parent
}

var sortOrderMap = map[sort.Order]int{
	sort.OrderAsc:  1,
	sort.OrderDesc: -1,
}

// Sort adds a sort to the specification.
func (b *BaseSpecification) Sort(sortBy string, sortOrder sort.Order) specification.Set {
	if needsPreload, preload := b.parent.NeedsPreload(sortBy); needsPreload {
		b.aggregationPipeline = append(b.aggregationPipeline, preload...)
	}

	// warning: all models must have createdAt filed
	if sortBy == "" {
		sortBy = "createdAt"
	}

	fieldName := sortBy

	b.sort = append(b.sort, bson.E{Key: fieldName, Value: sortOrderMap[sortOrder]})

	return b.parent
}

// Limit adds a limit to the specification.
func (b *BaseSpecification) Limit(limit int) specification.Set {
	b.limit = limit
	return b.parent
}

func (b *BaseSpecification) LessThan(field string, value interface{}) specification.Set {
	b.filters = append(b.filters, bson.M{field: bson.M{"$lt": value}})
	return b.parent
}

func (b *BaseSpecification) GreaterThan(field string, value interface{}) specification.Set {
	b.filters = append(b.filters, bson.M{field: bson.M{"$gt": value}})
	return b.parent
}

// Direct functionalities for BaseSpecification

// PreQueryHook is called before the query is executed
// It's either empty or overridden by the child.
func (b *BaseSpecification) PreQueryHook() {}

// indexField marks a field as been indexed for preloading.
func (b *BaseSpecification) indexField(field string) {
	if b.preloadIndex == nil {
		b.preloadIndex = make(map[string]int)
	}

	if _, ok := b.preloadIndex[field]; !ok {
		b.preloadIndex[field] = len(b.aggregationPipeline)
	}
}

// NeedsPreload requires overriding in the child.
func (b *BaseSpecification) NeedsPreload(field string) (bool, []bson.M) {
	if _, ok := b.preloadIndex[field]; ok {
		return false, nil
	}

	defer b.indexField(field)

	return true, nil
}

func (b *BaseSpecification) GetPreloadIndex(field string) int {
	return b.preloadIndex[field]
}

func (b *BaseSpecification) ReplacePreload(field string, preload []bson.M) specification.Set {
	index := utils.SliceIndex(b.aggregationPipeline, func(el bson.M) bool {
		if e, ok := el["$lookup"]; ok {
			if e.(bson.M)["from"] == field {
				return true
			}
		}

		return false
	})

	if index == -1 {
		b.aggregationPipeline = append(b.aggregationPipeline, preload...)
	} else {
		preload = append(preload, b.aggregationPipeline[index+1:]...)
		b.aggregationPipeline = append(b.aggregationPipeline[:index], preload...)
	}

	b.indexField(field)

	return b.parent
}
