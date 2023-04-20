package database

type Mapper[T, J Document] interface {
	MapToEntity(model J) T
	MapToModel(T) J

	MapToEntities(model []J) []T
	MapToModels(identity []T) []J
}
