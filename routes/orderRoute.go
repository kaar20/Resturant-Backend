package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/kaar20/resturant_backend/controllers"
)

func OrderRoutes(route *gin.Engine) {
	route.GET("/Orders", controller.GetOrders())
	route.POST("/Orders", controller.CreateOrder())
	route.GET("/Orders/:id", controller.GetOrder())
	route.PATCH("/Orders/:id", controller.UpdateOrder())
	route.DELETE("/Orders/:id", controller.DeleteOrders())
}
