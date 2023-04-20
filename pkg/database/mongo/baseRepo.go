package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mong-specification/pkg/infrastructure/database"
	specification "mong-specification/pkg/infrastructure/database/specefication"
	"mong-specification/pkg/utils"
)

type ModelRepo interface {
	FindOneBy(ctx context.Context, spec specification.Set)
	FindBy(ctx context.Context, spec specification.Set)
	ToModel() interface{}
}

type RepoBase[DbModel database.Document, CoreModel any] struct {
	db *mongo.Database
}

func NewRepo[DbModel database.Document, CoreModel any](getter database.DataContextGetter) *RepoBase[DbModel, CoreModel] {
	db := getter.GetDataContext().(*mongo.Database)
	return &RepoBase[DbModel, CoreModel]{
		db: db,
	}
}

func (r *RepoBase[D, C]) FindBy(ctx context.Context, spec specification.Set) ([]*C, error) {
	var results []*C
	var d D

	spec.WithContext(ctx)
	cursor, err := r.db.Collection(d.CollectionName()).Aggregate(ctx, spec.Query(), database.DiskAggregationOption)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var result D
		err = cursor.Decode(&result)

		if err != nil {
			return nil, err
		}

		results = append(results, r.ToModel(result))
	}

	return results, nil
}

func (r *RepoBase[D, C]) ToModel(a D) *C {
	// TODO implement me
	panic("implement me")
}

func (r *RepoBase[D, C]) FindOneBy(ctx context.Context, spec specification.Set) (*C, error) {
	var results []D
	var d D

	spec.WithContext(ctx).Limit(1)
	cursor, err := r.db.Collection(d.CollectionName()).Aggregate(ctx, spec.Query())

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

	return r.ToModel(results[0]), nil
}

func (r *RepoBase[D, C]) Update(ctx context.Context, id string, entry D) (*C, error) {
	filter := bson.M{"_id": id}
	data := utils.ToMap(entry, utils.MethodUpdate)

	_, err := r.db.Collection(entry.CollectionName()).UpdateOne(ctx, filter, bson.M{"$set": data})
	if err != nil {
		return nil, err
	}

	return r.ToModel(entry), nil
}

func (r *RepoBase[D, C]) Create(ctx context.Context, entry D) (string, error) {
	col := r.db.Collection(entry.CollectionName())
	data := utils.ToMap(entry)
	i, err := col.InsertOne(ctx, data)

	if err != nil {
		// TODO: move the logs to service
		return "", err
	}

	return i.InsertedID.(string), nil
}

func (r *RepoBase[D, C]) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	var d D
	_, err := r.db.Collection(d.CollectionName()).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
