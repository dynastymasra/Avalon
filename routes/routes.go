package routes

import (
	"avalon/controller"
	"avalon/controller/order"
	"avalon/middleware"

	"gopkg.in/gin-gonic/gin.v1"
)

// AvalonRouters endpoint
func AvalonRouters(router *gin.RouterGroup) {

	router.GET("/ping", controller.PingController)
	router.GET("/orders", middleware.RequestType(), order.GetAllOrderController)

	router.POST("/orders", middleware.RequestType(), order.CreateOrderController)

	router.HEAD("/ping", controller.PingController)
}
