package model

import (
	"github.com/behrouz-rfa/mongo-specification/example/repo"
	data "github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database"
)

type MongoRepoFactory struct {
}

func NewMongoRepoFactory() MongoRepoFactory {
	return MongoRepoFactory{}
}

func (g MongoRepoFactory) NewUser(getter data.DataContextGetter) repo.User {
	return NewUser(getter)
}
