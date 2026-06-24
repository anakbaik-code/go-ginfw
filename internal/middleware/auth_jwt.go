package middleware

import (
	"go-fwgin/internal/config"
	"go-fwgin/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJwtMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// check format use Bearer<Token>
		prefix := strings.HasPrefix(authHeader, "Bearer")

		if authHeader == "" || !prefix {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "JWT Security: Token missing or incorrect format! Must be 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		// trims teks "bearer" to get token string
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// validate token
		claims, err := utils.ValidateToken(tokenString, cfg.JwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"errors": "JWT Security Guard: Your token is invalid or has expired!" + err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserId)
		c.Set("role", claims.Role)
		c.Next()
	}
}
