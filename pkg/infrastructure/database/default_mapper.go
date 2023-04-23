package database

import "github.com/behrouz-rfa/nilmapper"

type defaultMapper[T, J Document] struct{}

func NewDefaultMapper[T, J Document]() Mapper[T, J] {
	return defaultMapper[T, J]{}
}

func (defaultMapper[T, J]) MapToEntity(model J) (out T) {
	nilmapper.Copy(model, &out)
	return
}

func (defaultMapper[T, J]) MapToModel(entity T) (out J) {
	nilmapper.Copy(entity, &out)
	return
}

func (d defaultMapper[T, J]) MapToEntities(models []J) (out []T) {
	out = []T{}
	for k := range models {
		out = append(out, d.MapToEntity(models[k]))
	}
	return out
}

func (d defaultMapper[T, J]) MapToModels(identities []T) (out []J) {
	out = []J{}

	for k := range identities {
		out = append(out, d.MapToModel(identities[k]))
	}
	return
}
