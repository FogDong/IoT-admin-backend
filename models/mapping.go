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
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	ProductID   bson.ObjectId `json:"productId" bson:"productId"`
	ProductName string        `json:"productName" bson:"productName"`
	Mapping     []string      `json:"mapping,omitempty" bson:"mapping,omitempty"`
}
