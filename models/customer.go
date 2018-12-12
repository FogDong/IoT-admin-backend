package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionCustomer holds the name of the customer collection
	CollectionCustomer = "customer"
)

// Customer model
type Customer struct {
	Id             bson.ObjectId `json:"_id" bson:"_id"`
	Name           string        `json:"name" bson:"name"`
	Description    string        `json:"description,omitempty" bson:"description,omitempty"`
	OrganizationId bson.ObjectId `json:"organizationId" bson:"organizationId"`
	CreatedBy      bson.ObjectId `json:"createdBy" bson:"createdBy"`
	CreatedAt      time.Time     `json:"createdAt" bson:"createdAt"`
	MemberCount   int           `json:"memberCount,omitempty" bson:"memberCount,omitempty"`
	ProductCount  int           `json:"productCount,omitempty" bson:"productCount,omitempty"`
	DeviceCount   int           `json:"deviceCount,omitempty" bson:"deviceCount,omitempty"`
}
