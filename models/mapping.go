package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Mapping model
type Mapping struct {
	Id          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	ProductId   bson.ObjectId `json:"productId" bson:"productId"`
	Mapping     []string      `json:"mapping" bson:"mapping"`
}
