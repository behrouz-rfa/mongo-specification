package repo

import "mong-specification/pkg/infrastructure/database"

type RepoFactory interface {
	NewUser(getter database.DataContextGetter) User
}
