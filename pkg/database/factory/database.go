package factory

import (
	"errors"
	"mong-specification/pkg/database/mongo"
	"mong-specification/pkg/infrastructure/database"
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
		cfg := param[0].(PgConfig)
		return mongo.NewPgController(mongo.Config{
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

type PgConfig struct {
	mongo.Config
}
