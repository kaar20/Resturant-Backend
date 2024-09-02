package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kaar20/resturant_backend/controllers"
)

func UserRoute(route *gin.Engine) {
	route.GET("/users", controller.GetUsers())
	route.GET("/users/:id", controller.GetUser())
	route.POST("/users/signup", controller.Signup())
	route.POST("/users/login", controller.Login())
	// route.PUT("/users/:id", controller.UpdateUser())
	// route.DELETE("/user/:id", controller.DeleteUser())
}
