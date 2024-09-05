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

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
var validtor = validator.New()

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get all menus from db
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data from the database"})
			// return
		}
		var allMenus []bson.M

		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
			// c.JSON(/)

		}
		c.JSON(http.StatusOK, &allMenus)

	}

}

func CreateMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validation := validate.Struct(menu)

		if validation != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": validation.Error()})
			return
		}

		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result, insertError := menuCollection.InsertOne(ctx, menu)
		if insertError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data into the database"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data:": result})
		// c.JSON(http.StatusCreated, gin.H{"Data": result})

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := c.Param("id")
		var menu models.Menu
		var err = menuCollection.FindOne(ctx, bson.M{"id": menuId}).Decode(&menu)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		}
		c.JSON(http.StatusOK, menu)

		defer cancel()

	}
}

func inTimestamp(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)

}
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		menuId := c.Param("id")
		filter := bson.M{"id": menuId}

		var updateObj primitive.D

		if menu.Start_Date != nil && menu.End_Date != nil {
			if inTimestamp(*menu.Start_Date, *menu.End_Date, time.Now()) {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": "Kindly Re-type the time"})
				defer cancel()
				return
			}

			updateObj = append(updateObj, bson.E{"start_date", menu.Start_Date})
			updateObj = append(updateObj, bson.E{"end_date", menu.End_Date})

			if menu.Name != "" {
				updateObj = append(updateObj, bson.E{"name", menu.Name})
			}
			if menu.Category != "" {
				updateObj = append(updateObj, bson.E{"category", menu.Category})
			}
			menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})

			upsert := true
			opt := options.UpdateOptions{
				Upsert: &upsert,
			}

			result, err := menuCollection.UpdateOne(
				ctx, filter, bson.E{"$set", updateObj},
				&opt,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating the menu"})
				defer cancel()
				return
			}
			defer cancel()

			c.JSON(http.StatusOK, result)

		}

	}
}
func DeleteMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// route.GET("/menus", controller.GetMenus())
// 	route.POST("/menus", controller.CreateMenus())
// 	route.GET("/menus/:id", controller.GetMenu())
// 	route.PATCH("/menus/:id", controller.UpdateMenu())
// 	route.DELETE("/menus/:id", controller.DeleteMenu())
