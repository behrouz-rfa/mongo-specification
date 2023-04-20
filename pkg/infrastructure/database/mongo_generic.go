package database

import (
	"context"
	"encoding/json"
	specification "github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database/specefication"
	"github.com/fatih/structs"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var pluralizer *pluralize.Client

func Init() {
	pluralizer = pluralize.NewClient()
}

func collectionName[T any](s T) string {
	name := structs.Name(s)
	name = strcase.ToLowerCamel(name)
	name = pluralizer.Plural(name)

	return name
}

type FindByParams[Q, T any] struct {
	Spec       specification.Set
	Collection *mongo.Collection
	ToModel    func(Q) *T
}

func FindOneBy[Q, T any](ctx context.Context, params *FindByParams[Q, T]) (*T, error) {
	var results []Q

	params.Spec.WithContext(ctx).Limit(1)
	cursor, err := params.Collection.Aggregate(ctx, params.Spec.Query())

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	if len(results) < 1 {
		return nil, nil
	}

	return params.ToModel(results[0]), nil
}

func FindBy[Q, T any](ctx context.Context, params *FindByParams[Q, T]) ([]*T, error) {
	var results []*T

	params.Spec.WithContext(ctx)

	cursor, err := params.Collection.Aggregate(ctx, params.Spec.Query(), DiskAggregationOption)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var result Q
		err = cursor.Decode(&result)

		if err != nil {
			return nil, err
		}

		results = append(results, params.ToModel(result))
	}

	return results, nil
}

func LogList(ctx context.Context, m *mongo.Collection, spec specification.Set) {

	obj := make([]map[string]interface{}, 0)

	cursor, err := m.Aggregate(ctx, spec.Query(), DiskAggregationOption)

	if err != nil {
		log.Println(err)
	}

	for cursor.Next(ctx) {
		var result map[string]interface{}
		err = cursor.Decode(&result)

		if err != nil {
			log.Println(err)
		}

		obj = append(obj, result)
	}

	res, _ := json.MarshalIndent(obj[0], "", "  ")

	log.Println(string(res))
}

func LogObject(ctx context.Context, m *mongo.Collection, spec specification.Set) {
	obj := make([]map[string]interface{}, 0)

	cursor, err := m.Aggregate(ctx, spec.Query(), DiskAggregationOption)

	if err != nil {
		log.Println(err)
	}

	err = cursor.All(ctx, &obj)

	if err != nil {
		log.Println(err)
	}

	if len(obj) < 1 {
		return
	}

	res, _ := json.MarshalIndent(obj[0], "", "  ")

	log.Println(string(res))
}
