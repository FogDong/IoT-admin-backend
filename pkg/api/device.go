package api

import (
	"IoT-admin-backend/middleware"
	"IoT-admin-backend/pkg/handler"

	"github.com/gin-gonic/gin"
)

func GetDeviceAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1alpha1")
	v1.Use(middleware.JWTAuth())

	v1.GET("/device", handler.ListDevice)
	v1.GET("/device/:_id", handler.GetDevice)
	v1.POST("/device", handler.CreateDevice)
	v1.PUT("/device/:_id", handler.UpdateDevice)
	v1.DELETE("/device/:_id", handler.DeleteDevice)

	v1.GET("/customer/:_id/device", handler.ListCustomerDevices)
	v1.GET("/product/:_id/device", handler.ListProductDevices)
	v1.GET("/customer/:_cid/product/:_id/device", handler.ListCustomerProductDevices)
}
