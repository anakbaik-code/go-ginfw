package middleware

import (
	"go-fwgin/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiKeyMidlleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey != cfg.ApiKey {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Security guard: Your API key is incorrect or not registered!",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
