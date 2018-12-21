package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionDevice holds the name of the device collection
	CollectionDevice = "device"
)

// Device model
type Device struct {
	ID               bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name             string        `json:"name" bson:"name"`
	Status           int           `json:"status" bson:"status"`
	Description      string        `json:"description,omitempty" bson:"description,omitempty"`
	DeviceKey        string        `json:"deviceKey,omitempty" bson:"deviceKey,omitempty"`
	DeviceSecret     string        `json:"deviceSecret,omitempty" bson:"deviceSecret,omitempty"`
	ProductID        bson.ObjectId `json:"productId" bson:"productId"`
	CustomerID       bson.ObjectId `json:"customerId" bson:"customerId"`
	OrganizationID   bson.ObjectId `json:"organizationId" bson:"organizationId"`
	MappingID        bson.ObjectId `json:"mappingId" bson:"mappingId"`
	CreatedBy        bson.ObjectId `json:"createdBy" bson:"createdBy"`
	OrganizationName string        `json:"organizationName" bson:"organizationName"`
	ProductName      string        `json:"productName" bson:"productName"`
	CreatedName      string        `json:"createdName" bson:"createdName"`
	CustomerName     string        `json:"customerName" bson:"customerName"`
	MappingName      string        `json:"mappingName" bson:"mappingName"`
	CreatedAt        time.Time     `json:"createdAt" bson:"createdAt"`
	ActivatedAt      time.Time     `json:"activatedAt" bson:"activatedAt"`
	Shadow           Shadow
	Tags             Tags
}

type Shadow struct {
	State     []string `json:"state,omitempty" bson:"state,omitempty"`
	Metadata  []string `json:"metadata,omitempty" bson:"metadata,omitempty"`
	TimeStamp int      `json:"timeStamp,omitempty" bson:"timeStamp,omitempty"`
	Version   string   `json:"version,omitempty" bson:"version,omitempty"`
}

type Tags struct {
}
