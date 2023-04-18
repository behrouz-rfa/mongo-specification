package mongo

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

func (d *MongoController) GetTransactionFactory() error {
	if !d.initialized {

		err := d.Init()
		if err != nil {
			return err
		}
	}

	return nil
}
func (d *MongoController) Generate() error {

	return nil
}
func (d *MongoController) Init() error {
	return nil
}
func (t *MongoController) GetDataContext() any {
	return t.db
}
