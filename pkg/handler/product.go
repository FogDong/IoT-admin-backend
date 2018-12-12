package handler

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"IoT-admin-backend/models"

	"github.com/gin-gonic/gin"
)

// List all products
func ListProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var products []models.Product
	err := db.C(models.CollectionProduct).Find(nil).All(&products)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, products)
}

// Get a product
func GetProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var product models.Product

	err := db.C(models.CollectionProduct).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&product)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, product)
}

// Create a product
func CreateProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionProduct).Insert(product)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, product)
}

// Delete product
func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionProduct).Remove(query)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, nil)
}

// Update product
func UpdateProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.Error(err)

		return
	}

	// 查找原来的文档
	query := bson.M{
		"_id": bson.ObjectIdHex(c.Param("_id")),
	}

	// 更新
	err = db.C(models.CollectionProduct).Update(query, product)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, product)
}

// List all organization products
func ListOrgProducts(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var products []models.Product
	query := bson.M{
		"orgnizationId": bson.ObjectIdHex(c.Param("_id")),
	}
	err := db.C(models.CollectionProduct).Find(query).All(&products)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, products)
}

