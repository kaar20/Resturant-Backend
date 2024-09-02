package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kaar20/resturant_backend/database"
	"github.com/kaar20/resturant_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// route.GET("/foods", controller.GetFoods())
// route.POST("/foods", controller.CreateFood())
// route.GET("/foods/:id", controller.GetFood())
// route.PATCH("/foods/:id", controller.UpdateFood())
// route.DELETE("/foods/:id", controller.DeleteFood())
var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		defer cancel()
		validation := validate.Struct(food)

		if validation != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": validation.Error()})
			return
		}

		err := menuCollection.FindOne(ctx, bson.M{"id": food.Food_id}).Decode(&menu)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
			return
		}
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, err := foodCollection.InsertOne(ctx, food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occurred while Inserting the food"})
			return

		}

		defer cancel()

		c.JSON(http.StatusCreated, gin.H{"Data": result})

	}

}

func round(num float64) int {
	return int(num + 0.5)
}

func toFixed(num float64, precesion int) float64 {
	return float64(round(num*float64(10^precesion))) / float64(10^precesion)
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("id")
		var food models.Food
		err := foodCollection.FindOne(ctx, bson.M{"id": foodId}).Decode(&food)

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occurred while Fetching the food "})
			return
		}

		c.JSON(http.StatusOK, food)

	}

}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
