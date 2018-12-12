package api

import (
	"IoT-admin-backend/middleware"
	"IoT-admin-backend/pkg/handler"

	"github.com/gin-gonic/gin"
)

func GetMappingAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1alpha1")
	v1.Use(middleware.JWTAuth())

	v1.GET("/product/:_id/mapping", handler.ListMapping)
	v1.GET("/mapping/:_id", handler.GetMapping)
	v1.POST("/mapping", handler.CreateMapping)
	v1.PUT("/mapping/:_id", handler.UpdateMapping)
	v1.DELETE("/mapping/:_id", handler.DeleteMapping)
}
