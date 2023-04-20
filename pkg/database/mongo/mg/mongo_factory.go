package mg

import (
	"context"
	"github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionFactory struct {
	db *MongoDatabase
}
type Transaction struct {
	database *mongo.Database
	client   *mongo.Client
	session  mongo.Session
	commited bool
}

func (m *Transaction) Begin(_ context.Context) error {

	session, err := m.client.StartSession()
	if err != nil {
		return err
	}
	m.session = session
	return nil
}

func (m *Transaction) Rollback(ctx context.Context) error {
	if m.session == nil {
		return nil
	}
	err := m.session.AbortTransaction(ctx)

	return err
}

func (m *Transaction) RollbackUnlessCommitted(ctx context.Context) {
	if !m.commited {
		m.Rollback(ctx)
	}
}

func (m *Transaction) Commit(ctx context.Context) error {
	err := m.session.CommitTransaction(ctx)
	m.commited = true

	return err
}

func (m *Transaction) GetDataContext() any {
	return m.database
}

func NewTransactionFactory(db *MongoDatabase) database.MongoTransactionFactory {
	return TransactionFactory{db: db}
}

func (t TransactionFactory) New() database.Transaction {
	return &Transaction{database: t.db.Database, client: t.db.Client}
}
