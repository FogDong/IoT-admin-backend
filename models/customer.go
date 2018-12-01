package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Customer model
type Customer struct {
	Id             bson.ObjectId `json:"_id" bson:"_id"`
	Name           string        `json:"name" bson:"name"`
	Description    string        `json:"description,omitempty" bson:"description,omitempty"`
	OrganizationId bson.ObjectId `json:"organizationId" bson:"organizationId"`
	CreatedBy      bson.ObjectId `json:"createdBy" bson:"createdBy"`
	CreatedAt      time.Time     `json:"createdAt" bson:"createdAt"`
}
