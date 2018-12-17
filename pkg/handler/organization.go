package handler

import (
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   orgs,
	})
}

// Get a organization
func GetOrg(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var org models.Organization

	err := db.C(models.CollectionOrg).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&org)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   org,
	})
}

// Create a organization
func CreateOrg(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var org models.Organization
	err := c.BindJSON(&org)
	org.CreatedAt = time.Now()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionOrg).Insert(org)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	err = db.C(models.CollectionUser).Update(bson.M{"_id": org.CreatedBy},
		bson.M{"$inc": bson.M{"orgCount": 1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
	})
}

// Delete organization
func DeleteOrg(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var org models.Organization

	err := db.C(models.CollectionOrg).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&org)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionUser).Update(bson.M{"_id": org.CreatedBy},
		bson.M{"$inc": bson.M{"orgCount": -1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionOrg).Remove(bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
	})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   org,
	})
}
