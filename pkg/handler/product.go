package handler

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"IoT-admin-backend/middleware"
	"IoT-admin-backend/models"

	"github.com/gin-gonic/gin"
)

// List all products
func ListProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var products []models.Product
	err := db.C(models.CollectionProduct).Find(nil).All(&products)
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
		"data":   products,
	})
}

// Get a product
func GetProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var product models.Product

	err := db.C(models.CollectionProduct).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&product)
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
		"data":   product,
	})
}

// List products from name
func ListNameProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var products []models.Product
	err := db.C(models.CollectionProduct).
		Find(bson.M{"name": bson.M{"$regex": `/c.Param("name")/`}}).
		All(&products)
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
		"data":   products,
	})
}

// Create a product
func CreateProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	product.CreatedBy = claims.ID

	err = db.C(models.CollectionProduct).Insert(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionUser).Update(bson.M{"_id": product.CreatedBy},
		bson.M{"$inc": bson.M{"productCount": 1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionCustomer).Update(bson.M{"_id": product.CustomerID},
		bson.M{"$inc": bson.M{"productCount": 1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionOrg).Update(bson.M{"_id": product.OrganizationID},
		bson.M{"$inc": bson.M{"productCount": 1}})
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

// Delete product
func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var product models.Product

	err := db.C(models.CollectionProduct).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionUser).Update(bson.M{"_id": product.CreatedBy},
		bson.M{"$inc": bson.M{"productCount": -1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionCustomer).Update(bson.M{"_id": product.CustomerID},
		bson.M{"$inc": bson.M{"productCount": -1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionOrg).Update(bson.M{"_id": product.OrganizationID},
		bson.M{"$inc": bson.M{"productCount": -1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionProduct).Remove(bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))})
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

// Update product
func UpdateProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var product models.Product
	err := c.BindJSON(&product)
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
	err = db.C(models.CollectionProduct).Update(query, product)
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
		"data":   product,
	})
}

// List all organization products
func ListOrgProducts(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var products []models.Product
	query := bson.M{
		"organizationId": bson.ObjectIdHex(c.Param("_id")),
	}
	err := db.C(models.CollectionProduct).Find(query).All(&products)
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
		"data":   products,
	})
}

// List all customer products
func ListCustomerProducts(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var customer models.Customer
	var products []models.Product
	var product models.Product

	err := db.C(models.CollectionCustomer).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&customer)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	for _, id := range customer.ProductID {
		err := db.C(models.CollectionProduct).
			FindId(id).
			One(&product)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"msg":    err.Error(),
			})
			return
		} else {
			products = append(products, product)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   products,
	})
}
