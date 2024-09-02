package controllers

import "github.com/gin-gonic/gin"

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get all orders from db
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
	}

}
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func DeleteOrders() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// route.GET("/Orders", controller.GetOrders())
// route.POST("/Orders", controller.CreateOrder())
// route.GET("/Orders/:id", controller.GetOrder())
// route.PATCH("/Orders/:id", controller.UpdateOrder())
// route.DELETE("/Orders/:id", controller.DeleteOrders())
