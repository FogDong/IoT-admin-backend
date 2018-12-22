package api

import (
	"IoT-admin-backend/middleware"
	"IoT-admin-backend/pkg/handler"

	"github.com/gin-gonic/gin"
)

func GetUserAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1alpha1")
	v1.POST("/login", handler.Login)
	v1.POST("/user", handler.CreateUser)

	v1.Use(middleware.JWTAuth())

	v1.GET("/user", handler.ListUser)
	v1.GET("/user/:_id", handler.GetUser)
	v1.GET("/email/:email", handler.GetEmail)
	v1.PUT("/user/:_id", handler.UpdateUser)
	v1.DELETE("/user/:_id", handler.DeleteUser)

	v1.GET("/organization/:_id/user", handler.ListOrgUsers)
	v1.GET("/customer/:_id/user", handler.ListCustomerUsers)
}
