package database

import (
	"errors"
	"fmt"
)

type Config struct {
	Host   string
	Port   int
	User   string
	Pass   string
	Name   string
	Driver string
	Schema string
}

func (config Config) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Pass, config.Name)
}

// TODO: Ambiguous definition for db first strategy, func name should be changed.
type DatabaseGenerator interface {
	Generate() error
	Init() error
}
type DatabaseController interface {
	DatabaseGenerator
}

var (
	ErrDbMigerationFailed = errors.New("database migeration failed")
	ErrDbNotFound         = errors.New("database not found")
)
