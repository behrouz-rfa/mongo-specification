package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
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

func (u User) SetCreatedAt() {
	//TODO implement me
	panic("implement me")
}

func (u User) SetUpdatedAt() {
	//TODO implement me
	panic("implement me")
}

func (u User) CollectionName() string {
	//TODO implement me
	panic("implement me")
}
