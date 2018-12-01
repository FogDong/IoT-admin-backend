package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Organization model
type Organization struct {
	Id          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	Address     string        `json:"address,omitempty" bson:"address,omitempty"`
	Contact     string        `json:"contact,omitempty" bson:"contact,omitempty"`
	CreatedBy   bson.ObjectId `json:"createdBy" bson:"createdBy"`
	CreatedAt   time.Time     `json:"createdAt" bson:"createdAt"`
}
