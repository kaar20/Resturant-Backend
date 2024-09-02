package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kaar20/resturant_backend/controllers"
)

func FoodRoutes(route *gin.Engine) {
	route.GET("/foods", controller.GetFoods())
	route.POST("/foods", controller.CreateFood())
	route.GET("/foods/:id", controller.GetFood())
	route.PATCH("/foods/:id", controller.UpdateFood())
	route.DELETE("/foods/:id", controller.DeleteFood())
}
