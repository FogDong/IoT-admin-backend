package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Product model
type Product struct {
	Id             bson.ObjectId `json:"_id" bson:"_id"`
	Name           string        `json:"name" bson:"name"`
	Description    string        `json:"description,omitempty" bson:"description,omitempty"`
	ProductKey     string        `json:"productKey" bson:"productKey"`
	ProductSecret  string        `json:"productSecret" bson:"productSecret"`
	OrganizationId bson.ObjectId `json:"organizationId" bson:"organizationId"`
	CreatedBy      bson.ObjectId `json:"createdBy" bson:"createdBy"`
	CreatedAt      time.Time     `json:"createdAt" bson:"createdAt"`
	Specification  Specification
	Tags           []string `json:"tags" bson:"tags"`
}

type Specification struct {
	Identifier     string        `json:"identifier" bson:"identifier"`
	Name           string        `json:"name" bson:"name"`
	Description    string        `json:"description,omitempty" bson:"description,omitempty"`
	DataType       DataType      `json:"productKey" bson:"productKey"`
	ProductSecret  string        `json:"productSecret" bson:"productSecret"`
	OrganizationId bson.ObjectId `json:"organizationId" bson:"organizationId"`
	CreatedBy      bson.ObjectId `json:"createdBy" bson:"createdBy"`
}

type DataType struct {
	Type  string `json:"type" bson:"type"`
	Specs Specs
}

type Specs struct {
	Min      float32 `json:"min" bson:"min"`
	Max      float32 `json:"max" bson:"max"`
	Unit     string  `json:"unit" bson:"unit"`
	UnitName string  `json:"unitName" bson:"unitName"`
}
