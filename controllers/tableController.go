package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kaar20/resturant_backend/database"
	"go.mongodb.org/mongo-driver/mongo"
)


var tableCollections *mongo.Collection= database.OpenCollection(database.Client,"table")
func GetTables() gin.HandlerFunc{
	return func(c *gin.Context){
        // get all tables from db
    }
}

func CreateTable() gin.HandlerFunc{
    return func(c *gin.Context){
    }
}
func GetTable() gin.HandlerFunc{
    return func(c *gin.Context){

    }
}
func UpdateTable() gin.HandlerFunc{
    return func(c *gin.Context){

    }
}
func Deletetables() gin.HandlerFunc{
    return func(c *gin.Context){

    }
}

// route.GET("/tables", controller.GetTables())
// route.POST("/tables", controller.CreateTable())
// route.GET("/tables/:id", controller.GetTable())
// route.PATCH("/tables/:id", controller.UpdateTable())
// route.DELETE("/tables/:id", controller.Deletetables())