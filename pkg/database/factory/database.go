package factory

import (
	"errors"
	"github.com/behrouz-rfa/mongo-specification/pkg/database/mongo/mg"
	"github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database"
)

type DatabaseType string

const (
	Mongo DatabaseType = "Mongo"
	Test  DatabaseType = "Test"
)

var (
	ErrMissingParameter = errors.New("missing parameter")
	ErrNotImplemented   = errors.New("not implemented")
)

func NewDatabaseController(name DatabaseType, definitions, baseEntities []database.DocumentBase, param ...any) database.DatabaseController {
	if name == "" {
		name = Test
	}
	switch name {
	case Test:
		fallthrough

	case Mongo:
		cfg := param[0].(MongoConfig)
		return mg.NewPgController(mg.Config{
			Host:     cfg.Host,
			Port:     cfg.Port,
			Username: cfg.Username,
			Password: cfg.Password,
			Name:     cfg.Name,
			Driver:   cfg.Driver,
			Schema:   cfg.Schema,
		}, definitions, baseEntities)

	}
	panic(ErrNotImplemented)
}

type MongoConfig struct {
	mg.Config
}
