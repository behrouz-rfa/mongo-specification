package database

import (
	"github.com/behrouz-rfa/mongo-specification/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// DiskAggregationOption tells mongo to store aggregation results on disk
// instead of in memory.
// use this option when the aggregation result is large.
// see: https://godoc.org/go.mongodb.org/mongo-driver/mongo/options#AggregateOptions
var DiskAggregationOption = &options.AggregateOptions{AllowDiskUse: utils.ToValue(true)}

type Document interface {
	GetID() string
	SetID(id string)
	GenerateID()
	SetCreatedAt()
	SetUpdatedAt()
	CollectionName() string
}

type DocumentBase struct {
	ID        string    `json:"_id" bson:"_id"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func (d DocumentBase) GetID() string {
	return d.ID
}

func (d DocumentBase) SetID(id string) {
	d.ID = id
}

func (d DocumentBase) GenerateID() {
	d.ID = utils.GenerateUUID()
}

func (d DocumentBase) SetCreatedAt() {
	d.CreatedAt = time.Now()
}

func (d DocumentBase) SetUpdatedAt() {
	d.UpdatedAt = time.Now()
}

func (d DocumentBase) CollectionName() string {
	return collectionName(d)
}
