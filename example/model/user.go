package model

import (
	"github.com/behrouz-rfa/mongo-specification/example/entity"
	"github.com/behrouz-rfa/mongo-specification/example/repo"
	monggDb "github.com/behrouz-rfa/mongo-specification/pkg/database/mongo"
	data "github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database"
	"github.com/google/uuid"
)

func NewUser(getter data.DataContextGetter) repo.User {
	r := monggDb.NewRepo(getter, data.NewDefaultMapper[User, entity.User]())
	return &r
}

type User struct {
	data.DocumentBase `bson:"inline"`
	Name              string
	Age               int
	Lastname          *string
}

func (d User) CollectionName() string {
	return "user"
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
