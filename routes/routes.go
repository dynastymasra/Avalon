package routes

import (
	"avalon/controller"
	"avalon/controller/order"

	"gopkg.in/gin-gonic/gin.v1"
)

// AvalonRouters endpoint
func AvalonRouters(router *gin.RouterGroup) {

	router.GET("/ping", controller.PingController)
	router.HEAD("/ping", controller.PingController)

	router.POST("/orders", order.CreateOrderController)
}
