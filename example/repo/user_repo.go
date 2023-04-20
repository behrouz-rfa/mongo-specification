package repo

import (
	"context"
	"mong-specification/example/model"
	"mong-specification/example/specification"
)

// UserRepository is the interface that wraps the basic CRUD operations for User.
type UserRepository interface {
	GetUser(ctx context.Context, spec specification.UserSpecification) (*model.User, error)
	GetUsers(ctx context.Context, spec specification.UserSpecification) ([]*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (string, error)
	UpdateUser(ctx context.Context, id string, user *model.User) error
	DeleteUser(ctx context.Context, id string) error
	NewUserSpecification(ctx context.Context) specification.UserSpecification
}
