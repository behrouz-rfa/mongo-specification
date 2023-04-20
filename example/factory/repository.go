package factory

import (
	"errors"
	"github.com/behrouz-rfa/mongo-specification/example/model"
	"github.com/behrouz-rfa/mongo-specification/example/repo"
)

var (
	ErrNotImplemented   = errors.New("not implemented")
	ErrMissingParameter = errors.New("missing parameter")
)

func NewRepoFactory() repo.RepoFactory {

	return model.NewMongoRepoFactory()

}
