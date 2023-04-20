package factory

import (
	"errors"
	"mong-specification/example/model"
	"mong-specification/example/repo"
)

var (
	ErrNotImplemented   = errors.New("not implemented")
	ErrMissingParameter = errors.New("missing parameter")
)

func NewRepoFactory() repo.RepoFactory {

	return model.NewMongoRepoFactory()

}
