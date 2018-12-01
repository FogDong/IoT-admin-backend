package main

import (
	"IoT-admin-backend/db"
	"IoT-admin-backend/middleware"
	"IoT-admin-backend/pkg/api"

	"github.com/gin-gonic/gin"
)

const Port = "9002"

func main() {
	db.Connect()
	router := gin.Default()
	router.Group("/api/v1/")
	router.Use(middleware.Connect)
	api.RunHTTPServer(router)
	// router.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "hello world",
	// 	})
	// })

	router.Run(":" + Port)
}
