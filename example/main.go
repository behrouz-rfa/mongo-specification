package main

import (
	"context"
	"fmt"
	"log"
	"mong-specification/example/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"mong-specification/pkg/database/factory"
	monggDb "mong-specification/pkg/database/mongo"
	data "mong-specification/pkg/infrastructure/database"
)

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

	// Begin transaction
	err = t.Begin(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	getterDb := t.GetDataContext().(*mongo.Database)
	one, err := getterDb.Collection("user").InsertOne(ctx, bson.M{"user_id": "123", "product": "computer"})
	if err != nil {
		// Rollback transaction on error
		t.Rollback(ctx)
		log.Println(err)
		return
	}

	// Commit transaction
	err = t.Commit(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	t2 := factory.New()
	// Begin transaction
	err = t2.Begin(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	repo := monggDb.NewRepo[*model.User, *model.User](t2)
	create, err := repo.Create(ctx, &model.User{
		DocumentBase: data.DocumentBase{},
		Name:         "data",
	})
	if err != nil {
		// Rollback transaction on error
		t2.Rollback(ctx)
		log.Println(err)
		return
	}

	// Commit transaction
	err = t2.Commit(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(one)
	fmt.Println(create)

	return
}

func DbController(config factory.MongoConfig) data.DatabaseController {
	controller := factory.NewDatabaseController(factory.Mongo,
		[]data.DocumentBase{},
		[]data.DocumentBase{},
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
