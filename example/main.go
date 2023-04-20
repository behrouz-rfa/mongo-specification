package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mong-specification/pkg/database/factory"
	"mong-specification/pkg/infrastructure/database"
)

type User struct {
	database.DocumentBase `bson:"inline"`
	Name                  string
}

func CollectionName() {

}

func main() {
	cfg := factory.MongoConfig{}
	cfg.Name = "test"
	cfg.Host = "localhost"
	cfg.Port = 27017
	cfg.Username = "root"
	cfg.Password = "root"
	cfg.Timeout = 20
	cfg.SSL = "false"
	cfg.Clustered = false
	cfg.Driver = "mongo"
	db := DbController(cfg)

	err := db.Generate()
	if err != nil {
		log.Println(err)
		return
	}

	err = db.Init()
	if err != nil {
		log.Println(err)
		return
	}

	factory, _ := db.GetTransactionFactory()
	t := factory.New()
	ctx := context.Background()
	t.Begin(ctx)
	defer t.Commit(ctx)
	data := t.GetDataContext().(*mongo.Database)
	one, err := data.Collection("user").InsertOne(ctx, bson.M{"user_id": "123", "product": "computer"})
	if err != nil {
		return
	}
	fmt.Println(one)
	return

}

func DbController(config factory.MongoConfig) database.DatabaseController {
	controller := factory.NewDatabaseController(factory.Mongo,
		[]database.DocumentBase{},
		[]database.DocumentBase{},
		config,
	)
	err := controller.Generate()
	if err != nil {
		panic(err)
	}
	err = controller.Init()
	if err != nil {
		panic(err)
	}

	return controller
}
