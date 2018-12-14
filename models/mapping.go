package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionMapping holds the name of the mapping collection
	CollectionMapping = "mapping"
)

// Mapping model
type Mapping struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	ProductID   bson.ObjectId `json:"productId" bson:"productId"`
	Mapping     []string      `json:"mapping,omitempty" bson:"mapping,omitempty"`
}
