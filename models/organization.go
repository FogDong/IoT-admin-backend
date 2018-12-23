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
	ID            bson.ObjectId   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string          `json:"name" binding:"required" bson:"name"`
	Description   string          `json:"description,omitempty" bson:"description,omitempty"`
	Address       string          `json:"address,omitempty" bson:"address,omitempty"`
	Contact       string          `json:"contact,omitempty" bson:"contact,omitempty"`
	Phone         string          `json:"phone,omitempty" bson:"phone,omitempty"`
	CreatedBy     bson.ObjectId   `json:"createdBy" bson:"createdBy"`
	CreatedName   string          `json:"createdName" bson:"createdName"`
	ProductID     []bson.ObjectId `json:"productId,omitempty" bson:"productId,omitempty"`
	CreatedAt     time.Time       `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	MemberCount   int             `json:"memberCount,omitempty" bson:"memberCount,omitempty"`
	CustomerCount int             `json:"customerCount,omitempty" bson:"customerCount,omitempty"`
	ProductCount  int             `json:"productCount,omitempty" bson:"productCount,omitempty"`
	DeviceCount   int             `json:"deviceCount,omitempty" bson:"deviceCount,omitempty"`
}
