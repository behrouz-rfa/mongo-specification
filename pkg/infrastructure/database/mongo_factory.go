package database

import "context"

type MongoTransactionFactory interface {
	New() Transaction
}

type Transaction interface {
	DataContextGetter
	Begin(ctx context.Context) error
	Rollback(ctx context.Context) error
	RollbackUnlessCommitted(ctx context.Context)
	Commit(ctx context.Context) error
}

type DataContextGetter interface {
	// GetDataContext is used when we want to access underlying database for crud
	GetDataContext() any
}
