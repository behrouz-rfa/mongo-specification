package model

import (
	"github.com/google/uuid"
	"mong-specification/example/entity"
	"mong-specification/example/repo"
	monggDb "mong-specification/pkg/database/mongo"
	data "mong-specification/pkg/infrastructure/database"
)

func NewUser(getter data.DataContextGetter) repo.User {
	r := monggDb.NewRepo(getter, data.NewDefaultMapper[User, entity.User]())
	return &r
}

type User struct {
	data.DocumentBase `bson:"inline"`
	Name              string
}

func (d User) CollectionName() string {
	return "users"
}
func (u User) GetID() string {

	return u.ID
}

func (u User) SetID(id string) {
	u.ID = id
}

func (u User) GenerateID() {
	u.ID = uuid.New().String()
}
