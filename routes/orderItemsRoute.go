package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/kaar20/resturant_backend/controllers"
)

func OrderItemsRoutes(route *gin.Engine) {
	route.GET("/orderItems", controller.GetOrderItems())
	route.POST("/orderItems", controller.CreateOrderItem())
	route.GET("/orderItems/:id", controller.GetOrderItem())
	route.PATCH("/orderItems/:id", controller.UpdateOrderItem())
	route.DELETE("/orderItems/:id", controller.DeleteOrderItems())
}
