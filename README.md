# Mongo Specification
This is a Go-based implementation of MongoDB specifications. It provides a way to interact with MongoDB using a high-level interface, allowing for easy database access and manipulation.

#Installation
To use this package, you need to have Go installed on your machine. You can then install the package using the following command:

```
go get -u github.com/behrouz-rfa/mongo-specification
```

# Usage
To use the package, you need to import it in your Go code:
```
import "mong-specification/pkg/database/mongo"
```

You can then create a new MongoDB repository using the following code:

```repo := mongo.NewRepo[*User, User](t2)

```

Here, User is the struct that represents the MongoDB document, and *User is the interface that defines the CRUD methods for the repository.

You can then call the CRUD methods on the repository to interact with the database.


# Example
Here's an example of how to use the package:
```
package main

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "log"
    "mong-specification/pkg/database/factory"
    monggDb "mong-specification/pkg/database/mongo"
    data "mong-specification/pkg/infrastructure/database"
)

type User struct {
    data.DocumentBase `bson:"inline"`
    Name              string
}

func (d User) CollectionName() string {
    return "users"
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

    getterDb := t.GetDataContext().(*mongo.Database)
    one, err := getterDb.Collection("user").InsertOne(ctx, bson.M{"user_id": "123", "product": "computer"})
    if err != nil {
        return
    }

    t.Commit(ctx)

    t2 := factory.New()
    t2.Begin(ctx)
    fmt.Println(one)
    repo := monggDb.NewRepo[*User, User](t2)
    create, err := repo.Create(ctx, &User{
        DocumentBase: data.DocumentBase{},
        Name:         "data",
    })
    if err != nil {
        t2.Rollback(ctx)
        return
    }
    fmt.Println(create)
    t2.Commit(ctx)
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


```
In this example, we create a MongoDB repository for the User struct, and then use the repository to create a new user document in the database.

# License
This package is licensed under the MIT license. See
