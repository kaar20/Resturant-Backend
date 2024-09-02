package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPassword(password string) string {
	return "Hashing your password"
}

func VerifyPassword(password string, providedPassword string) (string, error) {
	 return "Password verified", nil
}

// route.GET("/users", controller.GetUsers())
// 	route.GET("/users/:id", controller.GetUser())
// 	route.POST("/users/signup", controller.signup())
// 	route.POST("/users/login", controller.login())
