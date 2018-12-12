package handler

import (
	"fmt"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"IoT-admin-backend/models"

	"github.com/gin-gonic/gin"
)

// List all organizations
func ListOrg(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var orgs []models.Organization
	err := db.C(models.CollectionOrg).Find(nil).All(&orgs)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, orgs)
}

// Get a organization
func GetOrg(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var org models.Organization

	err := db.C(models.CollectionOrg).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&org)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, org)
}

// Create a organization
func CreateOrg(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var org models.Organization
	err := c.BindJSON(&org)
	org.CreatedAt = time.Now()
	fmt.Print(org)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionOrg).Insert(org)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, org)
}

// Delete organization
func DeleteOrg(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionOrg).Remove(query)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, nil)
}

// Update organization
func UpdateOrg(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var org models.Organization
	err := c.BindJSON(&org)
	if err != nil {
		c.Error(err)

		return
	}

	query := bson.M{
		"_id": bson.ObjectIdHex(c.Param("_id")),
	}

	err = db.C(models.CollectionOrg).Update(query, org)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, org)
}
