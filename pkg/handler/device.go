package handler

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"IoT-admin-backend/models"

	"github.com/gin-gonic/gin"
)

// List all devices
func ListDevice(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var devices []models.Device
	err := db.C(models.CollectionDevice).Find(nil).All(&devices)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, devices)
}

// Get a device
func GetDevice(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var device models.Device

	err := db.C(models.CollectionDevice).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&device)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, device)
}

// Create a device
func CreateDevice(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var device models.Device
	err := c.BindJSON(&device)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionDevice).Insert(device)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, device)
}

// Delete device
func DeleteDevice(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionDevice).Remove(query)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, nil)
}

// Update device
func UpdateDevice(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var device models.Device
	err := c.BindJSON(&device)
	if err != nil {
		c.Error(err)

		return
	}

	// 查找原来的文档
	query := bson.M{
		"_id": bson.ObjectIdHex(c.Param("_id")),
	}

	// 更新
	err = db.C(models.CollectionDevice).Update(query, device)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, device)
}


