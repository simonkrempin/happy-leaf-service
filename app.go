package main

import (
	"github.com/gin-gonic/gin"
	"happy-leaf-service/controller"
	"happy-leaf-service/middleware"
	"os"
)

func main() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	router := gin.New()

	privateRoutes := router.Group("/", middleware.AuthHandler())
	{
		privateRoutes.GET("/plants")
		privateRoutes.POST("/analyze-plant")
		privateRoutes.POST("/add-plant")
		privateRoutes.DELETE("/logout")
		privateRoutes.PUT("/account")
	}

	router.GET("/auth/login", controller.LoginHandler)
	router.POST("/auth/register", controller.RegisterHandler)

	router.Run(":" + port)
}
