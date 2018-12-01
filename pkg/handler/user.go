package handler

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"IoT-admin-backend/models"

	"github.com/gin-gonic/gin"
)

// List all users
func ListUser(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var users []models.User
	err := db.C(models.CollectionUser).Find(nil).All(&users)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, users)
}

// Get a user
func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var user models.User

	err := db.C(models.CollectionUser).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&user)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, user)
}

// Create a user
func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var user models.User
	err := c.BindJSON(&user)
	fmt.Print(user)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionUser).Insert(user)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, user)
}

// Delete user
func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionUser).Remove(query)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, nil)
}

// Update user
func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.Error(err)

		return
	}
	fmt.Print(user)

	// 查找原来的文档
	query := bson.M{
		"_id": bson.ObjectIdHex(c.Param("_id")),
	}

	// 更新
	err = db.C(models.CollectionUser).Update(query, user)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, user)
}
