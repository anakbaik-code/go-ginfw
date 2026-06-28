package category

import (
	"go-fwgin/internal/config"
	"go-fwgin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (h *HandlerCategory) RoutesUser(rg *gin.RouterGroup, cfg *config.Config) {
	categories := rg.Group("/categories")
	{
		categories.GET("", h.ListCategory)
		categories.GET("/:id", h.GetByID)
		protected := categories.Group("", middleware.AuthJwtMiddleware(cfg))
		{
			protected.POST("", h.Create)
			protected.PUT("/:id", h.Update)
			protected.DELETE("/:id", h.Delete)
		}

	}
}
