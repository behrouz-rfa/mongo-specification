package mongo

import (
	"context"
	"github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database"
	specification "github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database/specefication"
	"github.com/behrouz-rfa/mongo-specification/pkg/mspecification"
)

type GenericBaseMongoRepo[T database.Document, J database.Document] struct {
	GenericMongoRepo[T]
	mapper database.Mapper[T, J]
}

func NewRepo[T database.Document, J database.Document](getter database.DataContextGetter, mapper database.Mapper[T, J]) GenericBaseMongoRepo[T, J] {

	repo := NewGenericMongoRepo[T](getter)
	return GenericBaseMongoRepo[T, J]{
		GenericMongoRepo: repo, mapper: mapper,
	}
}

func (g GenericBaseMongoRepo[T, J]) FindOneBy(ctx context.Context, spec specification.Set) (J, error) {
	if spec == nil {
		spec = g.NewSpecification(ctx)
	}
	var t J
	by, err := g.GenericMongoRepo.FindOneBy(ctx, spec)
	if err != nil {
		return t, err
	}
	return g.mapper.MapToModel(by), nil
}

func (g GenericBaseMongoRepo[T, J]) FindBy(ctx context.Context, spec specification.Set) ([]J, error) {
	if spec == nil {
		spec = g.NewSpecification(ctx)
	}
	var t []J
	by, err := g.GenericMongoRepo.FindBy(ctx, spec)
	if err != nil {
		return t, err
	}
	return g.mapper.MapToModels(by), nil
}

func (g GenericBaseMongoRepo[T, J]) Create(ctx context.Context, j J) (string, error) {
	create, err := g.GenericMongoRepo.Create(ctx, g.mapper.MapToEntity(j))
	if err != nil {
		return "", err
	}
	return create, nil
}

func (g GenericBaseMongoRepo[T, J]) Update(ctx context.Context, id string, j J) error {

	err := g.GenericMongoRepo.Update(ctx, id, g.mapper.MapToEntity(j))
	if err != nil {
		return err
	}
	return nil
}

func (g GenericBaseMongoRepo[T, J]) Delete(ctx context.Context, id string) error {

	return g.GenericMongoRepo.Delete(ctx, id)

}

func (g GenericBaseMongoRepo[T, J]) NewSpecification(ctx context.Context) specification.Set {
	spec := new(mspecification.BaseSpecification)
	spec.WithContext(ctx)

	return spec
}
