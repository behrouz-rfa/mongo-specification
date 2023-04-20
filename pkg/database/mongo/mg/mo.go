package mg

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

type Config struct {
	Host                   string
	Port                   int
	Username               string
	Password               string
	Name                   string
	Driver                 string
	Schema                 string
	IgnorePermissionDenied bool
	Timeout                int
	SSL                    string
	Clustered              bool
}

const Public = "public"

func (config *Config) DSN() string {
	authSegment := ""

	if config.Username != "" && config.Password != "" {
		authSegment = config.Username + ":" + config.Password + "@"
	}

	prefix := "mongodb://"

	if config.Clustered {
		prefix = "mongodb+srv://"

		return prefix + authSegment + config.Host
	}

	return prefix +
		authSegment +
		config.Host +
		":" +
		strconv.Itoa(config.Port)
}

func (config *Config) RawDSN() string {
	if config.Schema == "" {
		config.Schema = Public
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password)
}

func (config *Config) GetTableSchema() string {
	if config.Schema == "" {
		config.Schema = Public
	}

	return config.Schema + "."
}

type MongoDatabase struct {
	Database *mongo.Database
	Client   *mongo.Client
	DBConfig Config
}

func NewMoDatabase(DBConfig Config) *MongoDatabase {
	md := MongoDatabase{}
	md.DBConfig = DBConfig
	return &md
}

func (md *MongoDatabase) GetMongo() *mongo.Database {
	return md.Database
}

func (md *MongoDatabase) open() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(md.DBConfig.Timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(md.DBConfig.DSN()))

	db := client.Database(md.DBConfig.Name)
	if err != nil {
		return err
	}
	md.Database = db
	md.Client = client
	return nil
}

func (md *MongoDatabase) Open() error {
	return md.open()
}

func (md *MongoDatabase) OpenRaw() error {
	//TODO
	return nil
}
