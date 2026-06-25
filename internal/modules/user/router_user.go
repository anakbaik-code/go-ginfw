package user

import (
	"go-fwgin/internal/config"
	"go-fwgin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (h *HandlerUser) RoutesUser(rg *gin.RouterGroup, cfg *config.Config) {
	users := rg.Group("/users")
	{
		users.POST("/register", h.Register)
		users.POST("/login", h.Login)
		users.POST("/refresh_token", h.RefreshToken)

		protected := users.Group("", middleware.AuthJwtMiddleware(cfg))
		{
			protected.GET("", h.ListUser)
			protected.PUT("/:id", h.UpdateUserProfile)
			protected.GET("/:id", h.GetByID)
			protected.GET("/active", h.GetActiveUsers)
		}

	}
}
