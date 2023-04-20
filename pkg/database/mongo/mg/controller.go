package mg

import "mong-specification/pkg/infrastructure/database"

type MongoController struct {
	db                *MongoDatabase
	entityDefinitions []database.DocumentBase
	initialized       bool
	baseEntities      []database.DocumentBase
	config            Config
}

func NewPgController(dbConfig Config, entityDefinitions, baseEntities []database.DocumentBase) database.DatabaseController {
	return &MongoController{entityDefinitions: entityDefinitions, baseEntities: baseEntities, config: dbConfig}
}

func (d *MongoController) GetTransactionFactory() (database.MongoTransactionFactory, error) {
	if !d.initialized {
		err := d.Init()
		if err != nil {
			return nil, err
		}
	}

	return NewTransactionFactory(d.db), nil
}

func (d *MongoController) Generate() error {
	md := NewMoDatabase(d.config)
	d.db = md
	return md.Open()
}
func (d *MongoController) Init() error {
	NewTransactionFactory(d.db)
	d.initialized = true
	return nil
}
