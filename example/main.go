package main

import (
	"context"
	"fmt"
	"log"
	"mong-specification/example/entity"
	"mong-specification/example/model"
	"mong-specification/pkg/mspecification"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	exFactory "mong-specification/example/factory"
	"mong-specification/pkg/database/factory"
	monggDb "mong-specification/pkg/database/mongo"
	data "mong-specification/pkg/infrastructure/database"
)

func main() {
	cfg := loadConfig()

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

	SimpleExampleAdMongo(db, err)

	SampleForGenericMongoRepo(db, err)

	//example by repo
	SampleExampleWithSpecification(db, err)

}

func SampleExampleWithSpecification(db data.DatabaseController, err error) {
	ctx := context.Background()
	factory, _ := db.GetTransactionFactory()
	t := factory.New()
	t.Begin(ctx)
	repofactory := exFactory.NewRepoFactory()
	userRepo := repofactory.NewUser(t)
	id, err := userRepo.Create(ctx, entity.User{
		Name: "User2",
	})
	if err != nil {
		t.Rollback(ctx)
		return
	}
	t.Commit(ctx)
	fmt.Println("id:", id)

	//example for Specification
	spec := mspecification.NewBaseSpecification()
	spec.FilterByID(id)
	by, err := userRepo.FindOneBy(ctx, spec)
	if err != nil {
		return
	}
	fmt.Println(by.Name)
}

func SampleForGenericMongoRepo(db data.DatabaseController, err error) {
	ctx := context.Background()
	factory, _ := db.GetTransactionFactory()
	t := factory.New()
	// Begin transaction
	err = t.Begin(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	repo := monggDb.NewGenericMongoRepo[*model.User](t)
	create, err := repo.Create(ctx, &model.User{
		DocumentBase: data.DocumentBase{},
		Name:         "data",
	})
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

	fmt.Println(create)

	return
}

func SimpleExampleAdMongo(db data.DatabaseController, err error) {
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
	_, err = getterDb.Collection("user").InsertOne(ctx, bson.M{"user_id": "123", "product": "computer"})
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
	return
}

func loadConfig() factory.MongoConfig {
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
	return cfg
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
