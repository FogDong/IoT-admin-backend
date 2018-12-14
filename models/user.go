package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionUser holds the name of the user collection
	CollectionUser = "user"
)

// User model
type User struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string        `json:"email" binding:"required" bson:"email"`
	// 0: admin 1: organization 2: customer
	Type int `json:"type" bson:"type"`
	// 0: admin 1: user
	Role           int           `json:"role" bson:"role"`
	Password       string        `json:"password" binding:"required" bson:"password"`
	Phone          string        `json:"phone,omitempty" bson:"phone,omitempty"`
	FullName       string        `json:"fullname,omitempty" bson:"fullname,omitempty"`
	CreatedBy      bson.ObjectId `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	OrganizationID bson.ObjectId `json:"organizationId,omitempty" bson:"organizationId,omitempty"`
	CustomerID     bson.ObjectId `json:"customerId,omitempty" bson:"customerId,omitempty"`
	OrgCount       int           `json:"orgCount,omitempty" bson:"orgCount,omitempty"`
	CustomerCount  int           `json:"customerCount,omitempty" bson:"customerCount,omitempty"`
	ProductCount   int           `json:"productCount,omitempty" bson:"productCount,omitempty"`
	DeviceCount    int           `json:"deviceCount,omitempty" bson:"deviceCount,omitempty"`
}

// Login param
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
