package controllers

import "github.com/gin-gonic/gin"


func GetOrderItems() gin.HandlerFunc{
	return func(c *gin.Context){
        // get all order items from db
    }
}

func CreateOrderItem() gin.HandlerFunc{
	return func(c *gin.Context){

    }
}
func GetOrderItem() gin.HandlerFunc{
	return func(c *gin.Context){

    }
}
func UpdateOrderItem() gin.HandlerFunc{
    return func(c *gin.Context){

    }
}


func DeleteOrderItems() gin.HandlerFunc{
	return func(c *gin.Context){

    }
}


// route.GET("/orderItems", controller.GetOrderItems())
// route.POST("/orderItems", controller.CreateOrderItem())
// route.GET("/orderItems/:id", controller.GetOrderItem())
// route.PATCH("/orderItems/:id", controller.UpdateOrderItem())
// route.DELETE("/orderItems/:id", controller.DeleteOrderItems())