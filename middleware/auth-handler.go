package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, authenticated := checkAuth(c)
		if !authenticated {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

func checkAuth(c *gin.Context) (interface{}, bool) {
	user := map[string]interface{}{
		"id":   1,
		"name": "John Doe",
	}

	return user, true
}
