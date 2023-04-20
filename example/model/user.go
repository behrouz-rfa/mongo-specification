package model

import data "mong-specification/pkg/infrastructure/database"

type User struct {
	data.DocumentBase `bson:"inline"`
	Name              string
}

func (d User) CollectionName() string {
	return "users"
}
