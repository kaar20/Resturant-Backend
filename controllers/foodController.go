package controllers

import (
	"context"
	"math"
	"net/http"
	"strconv"
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
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		recordPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPage < 1 {
			recordPage = 10
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPage
		if c.Query("startIndex") != "" {
			startIndex, err = strconv.Atoi(c.Query("startIndex"))
			if err != nil || startIndex < 0 {
				startIndex = 0
			}
		}

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{
			{"$group", bson.D{
				{"_id", bson.D{{"_id", "null"}}},
				{"total_count", bson.D{{"$sum", 1}}},
				{"data", bson.D{{"$push", "$$ROOT"}}},
			}},
		}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"food_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPage}}}},
			}},
		}

		// Execute the aggregation pipeline (matchStage, groupStage, projectStage)
		cursor, err := foodCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching the foods"})
			return
		}
		defer cancel()

		var foods []bson.M
		if err = cursor.All(ctx, &foods); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while processing the foods"})
			return
		}

		c.JSON(http.StatusOK, foods[0])
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

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)

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
	// return int(num + 0.5)
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precesion int) float64 {
	// return float64(round(num*float64(10^precesion))) / float64(10^precesion)
	output := math.Pow(10, float64(precesion))
	return float64(round(num*output)) / output
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
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food
		// foodId := c.Param("id")
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D

		if food.Name != nil {
			updateObj = append(updateObj, bson.E{"$set", bson.D{{"name", food.Name}}})
		}
		if food.Price != nil {
			updateObj = append(updateObj, bson.E{"$set", bson.E{"price", food.Price}})

		}
		if food.Food_image != nil {
			updateObj = append(updateObj, bson.E{"$set", bson.E{"food_image", food.Food_image}})
		}

		if food.Menu_id != "" {
			err := menuCollection.FindOne(ctx, bson.M{"id": food.Menu_id}).Decode(&menu)
			defer cancel()
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
				return
			}

			updateObj = append(updateObj, bson.E{"$set", bson.E{"menu", food.Price}})

		}
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertError := foodCollection.InsertOne(ctx, food)
		if insertError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occurred while Inserting the food"})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}

func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
