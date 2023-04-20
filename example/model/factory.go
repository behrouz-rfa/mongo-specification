package model

import (
	"mong-specification/example/repo"
	data "mong-specification/pkg/infrastructure/database"
)

type MongoRepoFactory struct {
}

func NewMongoRepoFactory() MongoRepoFactory {
	return MongoRepoFactory{}
}

func (g MongoRepoFactory) NewUser(getter data.DataContextGetter) repo.User {
	return NewUser(getter)
}
