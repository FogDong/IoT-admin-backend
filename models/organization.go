package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionOrg holds the name of the organization collection
	CollectionOrg = "organization"
)

// Organization model
type Organization struct {
	Id            bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string        `json:"name" binding:"required" bson:"name"`
	Description   string        `json:"description,omitempty" bson:"description,omitempty"`
	Address       string        `json:"address,omitempty" bson:"address,omitempty"`
	Contact       string        `json:"contact,omitempty" bson:"contact,omitempty"`
	CreatedBy     bson.ObjectId `json:"createdBy" bson:"createdBy"`
	CreatedAt     time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	MemberCount   int           `json:"memberCount,omitempty" bson:"memberCount,omitempty"`
	CustomerCount int           `json:"customerCount,omitempty" bson:"customerCount,omitempty"`
	ProductCount  int           `json:"productCount,omitempty" bson:"productCount,omitempty"`
	DeviceCount   int           `json:"deviceCount,omitempty" bson:"deviceCount,omitempty"`
}
