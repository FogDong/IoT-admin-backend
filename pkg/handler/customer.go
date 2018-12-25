package handler

import (
	"net/http"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"IoT-admin-backend/middleware"
	"IoT-admin-backend/models"

	"github.com/gin-gonic/gin"
)

// List all customers
func ListCustomer(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var customers []models.Customer
	err := db.C(models.CollectionCustomer).Find(nil).All(&customers)
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
		"data":   customers,
	})
}

// Get a customer
func GetCustomer(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var customer models.Customer

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

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   customer,
	})
}

// List customers from name
func ListNameCustomer(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var customers []models.Customer
	query := "/" + c.Param("name") + "/"
	err := db.C(models.CollectionCustomer).
		Find(bson.M{"name": bson.M{"$regex": query}}).
		All(&customers)
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
		"data":   customers,
	})
}

// Create a customer
func CreateCustomer(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var customer models.Customer
	err := c.BindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	customer.CreatedBy = claims.ID
	customer.CreatedAt = time.Now()

	err = db.C(models.CollectionCustomer).Insert(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionUser).Update(bson.M{"_id": customer.CreatedBy},
		bson.M{"$inc": bson.M{"customerCount": 1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionOrg).Update(bson.M{"_id": customer.OrganizationID},
		bson.M{"$inc": bson.M{"customerCount": 1}})
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

// Delete customer
func DeleteCustomer(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var customer models.Customer

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

	err = db.C(models.CollectionUser).Update(bson.M{"_id": customer.CreatedBy},
		bson.M{"$inc": bson.M{"customerCount": -1}})
	if err != nil {
		if !strings.Contains(err.Error(), `not found`) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"msg":    err.Error(),
			})
			return
		}
	}

	err = db.C(models.CollectionOrg).Update(bson.M{"_id": customer.OrganizationID},
		bson.M{"$inc": bson.M{"customerCount": -1}})
	if err != nil {
		if !strings.Contains(err.Error(), `not found`) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"msg":    err.Error(),
			})
			return
		}
	}

	err = db.C(models.CollectionCustomer).Remove(bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))})

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

// Update customer
func UpdateCustomer(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var customer models.Customer
	err := c.BindJSON(&customer)
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
	err = db.C(models.CollectionCustomer).Update(query, customer)
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
		"data":   customer,
	})
}

// List all organization customers
func ListOrgCustomers(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var customers []models.Customer
	query := bson.M{
		"organizationId": bson.ObjectIdHex(c.Param("_id")),
	}
	err := db.C(models.CollectionCustomer).Find(query).All(&customers)
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
		"data":   customers,
	})
}

// List all product customers
func ListProductCustomers(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var product models.Product
	var customers []models.Customer
	var customer models.Customer

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

	for _, id := range product.CustomerID {
		err := db.C(models.CollectionCustomer).
			FindId(id).
			One(&customer)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"msg":    err.Error(),
			})
			return
		} else {
			customers = append(customers, customer)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   customers,
	})
}
