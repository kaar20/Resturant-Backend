package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/kaar20/resturant_backend/controllers"
)

func MenuRoutes(route *gin.Engine) {
	route.GET("/menus", controller.GetMenus())
	route.POST("/menus", controller.CreateMenus())
	route.GET("/menus/:id", controller.GetMenu())
	route.PATCH("/menus/:id", controller.UpdateMenu())
	route.DELETE("/menus/:id", controller.DeleteMenu())
}
