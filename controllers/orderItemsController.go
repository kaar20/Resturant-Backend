package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetOrderItems() gin.HandlerFunc{
	return func(c *gin.Context){
        // get all order items from db
    }
}

func getOrderItemByOrder() gin.HandlerFunc{
	return func(c *gin.Context){

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

func ItemByOrder(id string)(orderItems []primitive.M, err error){
return 
}



// route.GET("/orderItems", controller.GetOrderItems())
// route.POST("/orderItems", controller.CreateOrderItem())
// route.GET("/orderItems/:id", controller.GetOrderItem())
// route.PATCH("/orderItems/:id", controller.UpdateOrderItem())
// route.DELETE("/orderItems/:id", controller.DeleteOrderItems())