package handler

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"IoT-admin-backend/middleware"
	"IoT-admin-backend/models"

	"github.com/gin-gonic/gin"
)

// List all devices
func ListDevice(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var devices []models.Device
	err := db.C(models.CollectionDevice).Find(nil).All(&devices)
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
		"data":   devices,
	})
}

// Get a device
func GetDevice(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var device models.Device

	err := db.C(models.CollectionDevice).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&device)
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
		"data":   device,
	})
}

// Create a device
func CreateDevice(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var device models.Device
	err := c.BindJSON(&device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	device.CreatedBy = claims.ID

	err = db.C(models.CollectionDevice).Insert(device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionUser).Update(bson.M{"_id": device.CreatedBy},
		bson.M{"$inc": bson.M{"deviceCount": 1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionCustomer).Update(bson.M{"_id": device.CustomerID},
		bson.M{"$inc": bson.M{"deviceCount": 1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionOrg).Update(bson.M{"_id": device.OrganizationID},
		bson.M{"$inc": bson.M{"deviceCount": 1}})
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

// Delete device
func DeleteDevice(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var device models.Device

	err := db.C(models.CollectionDevice).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionUser).Update(bson.M{"_id": device.CreatedBy},
		bson.M{"$inc": bson.M{"deviceCount": -1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionCustomer).Update(bson.M{"_id": device.CustomerID},
		bson.M{"$inc": bson.M{"deviceCount": -1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionOrg).Update(bson.M{"_id": device.OrganizationID},
		bson.M{"$inc": bson.M{"deviceCount": -1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionDevice).Remove(bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))})
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

// Update device
func UpdateDevice(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var device models.Device
	err := c.BindJSON(&device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	// 查找原来的文档
	query := bson.M{
		"_id": bson.ObjectIdHex(c.Param("_id")),
	}

	// 更新
	err = db.C(models.CollectionDevice).Update(query, device)
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
		"data":   device,
	})
}

// List all customer devices
func ListCustomerDevices(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var devices []models.Device
	query := bson.M{
		"customerId": bson.ObjectIdHex(c.Param("_id")),
	}
	err := db.C(models.CollectionDevice).Find(query).All(&devices)
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
		"data":   devices,
	})
}

// List all product devices
func ListProductDevices(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var devices []models.Device
	query := bson.M{
		"productId": bson.ObjectIdHex(c.Param("_id")),
	}
	err := db.C(models.CollectionDevice).Find(query).All(&devices)
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
		"data":   devices,
	})
}

// List all product devices
func ListCustomerProductDevices(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var devices []models.Device
	query := bson.M{
		"customerId": bson.ObjectIdHex(c.Request.Header.Get("cid")),
		"productId":  bson.ObjectIdHex(c.Request.Header.Get("pid")),
	}
	err := db.C(models.CollectionDevice).Find(query).All(&devices)
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
		"data":   devices,
	})
}
