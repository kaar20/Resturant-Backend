package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kaar20/resturant_backend/middleware"
	"github.com/kaar20/resturant_backend/routes"
	// "go.mongodb.org/mongo-driver/mongo"
)

// var foodCollection *mongo.Collection =data

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoute(router)
	router.Use(middleware.Authentication())

	// router.FoodRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.TableRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}
