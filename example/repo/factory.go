package repo

import "github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database"

type RepoFactory interface {
	NewUser(getter database.DataContextGetter) User
}
