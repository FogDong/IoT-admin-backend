package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Device model
type Device struct {
	Id           bson.ObjectId `json:"_id" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	Status       int           `json:"status" bson:"status"`
	Description  string        `json:"description,omitempty" bson:"description,omitempty"`
	DeviceKey    string        `json:"deviceKey" bson:"deviceKey"`
	DeviceSecret string        `json:"deviceSecret" bson:"deviceSecret"`
	ProductId    bson.ObjectId `json:"productId" bson:"productId"`
	CustomerId   bson.ObjectId `json:"customerId" bson:"customerId"`
	MappingId    bson.ObjectId `json:"mappingId" bson:"mappingId"`
	CreatedBy    bson.ObjectId `json:"createdBy" bson:"createdBy"`
	CreatedAt    time.Time     `json:"createdAt" bson:"createdAt"`
	ActivatedAt  time.Time     `json:"activatedAt" bson:"activatedAt"`
	Shadow       Shadow
	Tags         Tags
}

type Shadow struct {
	State     []string `json:"state" bson:"state"`
	Metadata  []string `json:"metadata" bson:"metadata"`
	TimeStamp int      `json:"timeStamp" bson:"timeStamp"`
	Version   string   `json:"version" bson:"version"`
}

type Tags struct {
}
