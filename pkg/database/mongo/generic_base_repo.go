package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mong-specification/pkg/infrastructure/database"
	specification "mong-specification/pkg/infrastructure/database/specefication"
	"mong-specification/pkg/utils"
)

type GenericMongoRepo[T database.Document] struct {
	db *mongo.Database
}

func NewGenericMongoRepo[T database.Document](getter database.DataContextGetter) GenericMongoRepo[T] {
	db := getter.GetDataContext().(*mongo.Database)
	return GenericMongoRepo[T]{
		db: db,
	}
}

func (r GenericMongoRepo[T]) FindBy(ctx context.Context, spec specification.Set) ([]T, error) {
	var results []T
	var d T

	spec.WithContext(ctx)
	cursor, err := r.db.Collection(d.CollectionName()).Aggregate(ctx, spec.Query(), database.DiskAggregationOption)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var result T
		err = cursor.Decode(&result)

		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (r GenericMongoRepo[T]) ToModel(a T) T {
	// TODO implement me
	panic("implement me")
}

func (r GenericMongoRepo[T]) FindOneBy(ctx context.Context, spec specification.Set) (T, error) {
	var results []T
	var d T

	spec.WithContext(ctx).Limit(1)
	cursor, err := r.db.Collection(d.CollectionName()).Aggregate(ctx, spec.Query())

	if err != nil {
		return d, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return d, err
	}

	if len(results) < 1 {
		return d, nil
	}

	return results[0], nil
}

func (r GenericMongoRepo[T]) Update(ctx context.Context, id string, entry T) (T, error) {
	filter := bson.M{"_id": id}
	data := utils.ToMap(entry, utils.MethodUpdate)
	var d T
	_, err := r.db.Collection(entry.CollectionName()).UpdateOne(ctx, filter, bson.M{"$set": data})
	if err != nil {
		return d, err
	}

	return entry, nil
}

func (r GenericMongoRepo[T]) Create(ctx context.Context, entry T) (string, error) {
	col := r.db.Collection(entry.CollectionName())
	data := utils.ToMap(entry)
	i, err := col.InsertOne(ctx, data)

	if err != nil {
		// TODO: move the logs to service
		return "", err
	}

	return i.InsertedID.(string), nil
}

func (r GenericMongoRepo[T]) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	var d T
	_, err := r.db.Collection(d.CollectionName()).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
