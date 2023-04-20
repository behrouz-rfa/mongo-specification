package repo

import "github.com/behrouz-rfa/mong-specification/pkg/infrastructure/database"

type RepoFactory interface {
	NewUser(getter database.DataContextGetter) User
}
