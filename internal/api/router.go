// internal/api/router.go
package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Simple in-memory store for API keys
var apiKeys = map[string]string{
	"BAPE-SECRET-KEY": "default-user",
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API key required"})
			return
		}

		user, exists := apiKeys[apiKey]
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			return
		}

		c.Set("user_id", user)
		c.Next()
	}
}

func RateLimitMiddleware() gin.HandlerFunc {
	// Allow 1 request per second with a burst of 5.
	limiter := rate.NewLimiter(rate.Every(1*time.Second), 5)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			return
		}
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.Use(AuthMiddleware())
	v1.Use(RateLimitMiddleware())
	{
		v1.POST("/workflow", StartWorkflowHandler)
	}

	return r
}
