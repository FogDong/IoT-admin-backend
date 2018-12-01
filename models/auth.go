package models

import "gopkg.in/mgo.v2/bson"

// Auth model
type Auth struct {
	Id       bson.ObjectId `json:"_id" bson:"_id"`
	Type     int           `json:"type" bson:"type"`
	Role     int           `json:"role" bson:"role"`
	Identity string        `json:"identity" bson:"identity"`
	User     string        `json:"user" bson:"user"`
}
