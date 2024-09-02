package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/kaar20/resturant_backend/controllers"
)

func TableRoutes(route *gin.Engine) {
	route.GET("/tables", controller.GetTables())
	route.POST("/tables", controller.CreateTable())
	route.GET("/tables/:id", controller.GetTable())
	route.PATCH("/tables/:id", controller.UpdateTable())
	route.DELETE("/tables/:id", controller.Deletetables())
}
