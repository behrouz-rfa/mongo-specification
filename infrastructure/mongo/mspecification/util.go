//nolint:deadcode,unused// created for generated code
package mspecification

import (
	"go.mongodb.org/mongo-driver/bson"
	"mong-specification/infrastructure/mongo/filters"
)

// stringFilterBson converts a StringFilter to a bson.M
// Example:
// stringFilterBson(filters.StringFilter{Eq: "test"}) // returns bson.M{"$eq": "test"}.
func stringFilterBson(filter filters.StringFilter) bson.M {
	bsonFilter := bson.M{}

	if filter.Eq != nil {
		bsonFilter["$regex"] = "^" + *filter.Eq + "$"
	}

	if filter.Neq != nil {
		bsonFilter["$regex"] = "^(?!" + *filter.Neq + "$)"
	}

	if filter.Contains != nil {
		bsonFilter["$regex"] = *filter.Contains
	}

	if filter.StartWith != nil {
		bsonFilter["$regex"] = "^" + *filter.StartWith
	}

	if filter.EndWith != nil {
		bsonFilter["$regex"] = *filter.EndWith + "$"
	}

	if len(bsonFilter) > 0 {
		bsonFilter["$options"] = "i"
	}

	return bsonFilter
}

// intFilterBson converts a IntFilter to a bson.M
// Example:
// intFilterBson(filters.IntFilter{Eq: 1}) // returns bson.M{"$eq": 1}.
func intFilterBson(filter filters.IntFilter) bson.M {
	bsonFilter := bson.M{}
	if filter.Eq != nil {
		bsonFilter["$eq"] = *filter.Eq
	}

	if filter.NEq != nil {
		bsonFilter["$ne"] = *filter.NEq
	}

	if filter.Gt != nil {
		bsonFilter["$gt"] = *filter.Gt
	}

	if filter.Gte != nil {
		bsonFilter["$gte"] = *filter.Gte
	}

	if filter.Lt != nil {
		bsonFilter["$lt"] = *filter.Lt
	}

	if filter.Lte != nil {
		bsonFilter["$lte"] = *filter.Lte
	}

	return bsonFilter
}

// timeRangeBson converts a DateFilter to a bson.M
// Example:
// timeRangeBson(filters.DateFilter{From: "2020-01-01"}) // returns bson.M{"$gt": "2020-01-01"}.
func timeRangeBson(filter filters.TimeRange) bson.M {
	bsonFilter := bson.M{}

	if filter.From != nil {
		bsonFilter["$gt"] = *filter.From
	}

	if filter.To != nil {
		bsonFilter["$lt"] = *filter.To
	}

	return bsonFilter
}

func floatFilterBson(filter filters.FloatFilter) bson.M {
	bsonFilter := bson.M{}
	if filter.Eq != nil {
		bsonFilter["$eq"] = *filter.Eq
	}

	if filter.NEq != nil {
		bsonFilter["$ne"] = *filter.NEq
	}

	if filter.Gt != nil {
		bsonFilter["$gt"] = *filter.Gt
	}

	if filter.Gte != nil {
		bsonFilter["$gte"] = *filter.Gte
	}

	if filter.Lt != nil {
		bsonFilter["$lt"] = *filter.Lt
	}

	if filter.Lte != nil {
		bsonFilter["$lte"] = *filter.Lte
	}

	return bsonFilter
}
