package middleware

import (
	"go-fwgin/internal/pkg/jwt"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJwtMiddleware(jwtService *jwt.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized,gin.H{
				"error":"Authorization header is required",
			})
			c.Abort()
			return 
		}

		// check format use Bearer<Token>
		if !strings.HasPrefix(authHeader,"Bearer") {
			c.JSON(http.StatusUnauthorized,gin.H{
				"error" : "Authorization Header Must Use Bearer <Token>",
			})
			c.Abort()
			return 
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims,err := jwtService.ValidateAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized,gin.H{
				"error":"Invalid or expired token",
			})
			c.Abort()
			return 
		}
		c.Set("user_id",claims.UserID)
		c.Set("role",claims.Role)
		c.Next()
	}
}
