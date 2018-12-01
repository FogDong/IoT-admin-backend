package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionUser holds the name of the user collection
	CollectionUser = "user"
)

// User model
type User struct {
	Id             bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Email          string        `json:"email" binding:"required" bson:"email"`
	Password       string        `json:"password" binding:"required" bson:"password"`
	Phone          string        `json:"phone,omitempty" bson:"phone,omitempty"`
	FullName       string        `json:"fullname,omitempty" bson:"fullname,omitempty"`
	CreatedBy      bson.ObjectId `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	OrganizationId bson.ObjectId `json:"organizationId,omitempty" bson:"organizationId,omitempty"`
}
