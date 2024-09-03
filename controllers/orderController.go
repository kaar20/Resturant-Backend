package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kaar20/resturant_backend/database"
	"github.com/kaar20/resturant_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ordersCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")
var validate = validator.New()

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := ordersCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error getting orders"})
			return
		}
		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders)
		// get all orders from db
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var table models.Table
		var order models.Order

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		validation := validate.Struct(order)
		if validation != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": validation.Error()})
			return
		}
		if order.Table_id != nil {
			err := tableCollections.FindOne(ctx, bson.M{"order": order.Table_id}).Decode(&table)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
				return
			}

		}
		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		result,err := ordersCollection.InsertOne(ctx,order)
		if err!= nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data into the database"})
            return
        }

		defer cancel()
		c.JSON(http.StatusOK,result)

	}

}
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var orderId = c.Param("id")
		var order models.Order

		var err = ordersCollection.FindOne(ctx, bson.M{"id": orderId}).Decode(&order)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, order)

	}
}
func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)

		var table models.Table
		var order models.Order
		var updateObj primitive.D

		var orderId = c.Param("id")
		
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
			return
		}

		if order.Table_id !=nil{
			err := menuCollection.FindOne(ctx,bson.M{"order":orderId}).Decode(&table)
			defer cancel()
			if err!= nil {
                c.JSON(http.StatusBadRequest,gin.H{"Error":"Menu Was Not Found"})
                return
            }
			updateObj  = append(updateObj, bson.E{"menu",order.Table_id})
		}
		order.Updated_at,_=time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", order.Updated_at})
		upsert := true
		filter := bson.M{"order_id":orderId}
		opt:=options.UpdateOptions{
			Upsert: &upsert,
		}

	result , err :=	ordersCollection.UpdateOne(
			ctx,
			filter,bson.D{
				{"$set",updateObj},
			},
			&opt,
		)
		if err!= nil {
            c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
            return
        }
		defer cancel()
		c.JSON(http.StatusOK,result)


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

func orderItemOrderCreater(order models.Order) string{
	var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
	order.Created_at , _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at , _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()
	ordersCollection.InsertOne(ctx,order)
	defer cancel()

	return order.Order_id


}