package repo

import (
	"context"
	"mong-specification/example/entity"
	specification "mong-specification/pkg/infrastructure/database/specefication"

	data "mong-specification/pkg/infrastructure/database"
)

// User is the interface that wraps the basic CRUD operations for User.
type User interface {
	FindOneBy(ctx context.Context, spec specification.Set) (entity.User, error)
	FindBy(ctx context.Context, spec specification.Set) ([]entity.User, error)
	Create(ctx context.Context, user entity.User) (string, error)
	Update(ctx context.Context, id string, user entity.User) (entity.User, error)
	Delete(ctx context.Context, id string) error
}

type UserRepoFactory interface {
	NewTender(data.DataContextGetter) User
}
