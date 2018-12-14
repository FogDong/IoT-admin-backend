package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionProduct holds the name of the product collection
	CollectionProduct = "product"
)

// Product model
type Product struct {
	ID             bson.ObjectId   `json:"_id" bson:"_id"`
	Name           string          `json:"name" bson:"name"`
	Description    string          `json:"description,omitempty" bson:"description,omitempty"`
	ProductKey     string          `json:"productKey,omitempty" bson:"productKey,omitempty"`
	ProductSecret  string          `json:"productSecret,omitempty" bson:"productSecret,omitempty"`
	OrganizationID bson.ObjectId   `json:"organizationId" bson:"organizationId"`
	CustomerID     []bson.ObjectId `json:"customerId,omitempty" bson:"customerId,omitempty"`
	CreatedBy      bson.ObjectId   `json:"createdBy" bson:"createdBy"`
	CreatedAt      time.Time       `json:"createdAt" bson:"createdAt"`
	Specification  Specification
	Tags           []string `json:"tags,omitempty" bson:"tags,omitempty"`
}

type Specification struct {
	Identifier  string `json:"identifier,omitempty" bson:"identifier,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	DataType    DataType
}

type DataType struct {
	Type  string `json:"type,omitempty" bson:"type,omitempty"`
	Specs Specs
}

type Specs struct {
	Min      float32 `json:"min,omitempty" bson:"min,omitempty"`
	Max      float32 `json:"max,omitempty" bson:"max,omitempty"`
	Unit     string  `json:"unit,omitempty" bson:"unit,omitempty"`
	UnitName string  `json:"unitName,omitempty" bson:"unitName,omitempty"`
}
