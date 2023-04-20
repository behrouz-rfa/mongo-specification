package factory

import (
	"errors"
	"github.com/behrouz-rfa/mong-specification/example/model"
	"github.com/behrouz-rfa/mong-specification/example/repo"
)

var (
	ErrNotImplemented   = errors.New("not implemented")
	ErrMissingParameter = errors.New("missing parameter")
)

func NewRepoFactory() repo.RepoFactory {

	return model.NewMongoRepoFactory()

}
