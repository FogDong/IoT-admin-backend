package handler

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"IoT-admin-backend/models"

	"github.com/gin-gonic/gin"
)

// List all customers
func ListCustomer(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var customers []models.Customer
	err := db.C(models.CollectionCustomer).Find(nil).All(&customers)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, customers)
}

// Get a customer
func GetCustomer(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var customer models.Customer

	err := db.C(models.CollectionCustomer).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&customer)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, customer)
}

// Create a customer 
func CreateCustomer(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var customer models.Customer
	err := c.BindJSON(&customer)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionCustomer).Insert(customer)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, customer)
}

// Delete customer
func DeleteCustomer(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionCustomer).Remove(query)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, nil)
}

// Update customer
func UpdateCustomer(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var customer models.Customer
	err := c.BindJSON(&customer)
	if err != nil {
		c.Error(err)

		return
	}

	// 查找原来的文档
	query := bson.M{
		"_id": bson.ObjectIdHex(c.Param("_id")),
	}

	// 更新
	err = db.C(models.CollectionCustomer).Update(query, UpdateCustomer)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, customer)
}

// List all organization customers
func ListOrgCustomers(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var customers []models.Customer
	query := bson.M{
		"orgnizationId": bson.ObjectIdHex(c.Param("_id")),
	}
	err := db.C(models.CollectionCustomer).Find(query).All(&customers)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, customers)
}

