package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	router := gin.New()

	router.Use(protectedRoute)

	router.GET("/auth/login")
	router.POST("/auth/register")

	router.Run(":" + port)
}

func protectedRoute(context *gin.Context) {
}
