package admin

import (
	"github.com/gin-gonic/gin"
	"go-fwgin/internal/config"
)

func (h *HandlerDashboardAdmin) RoutesDashboard(r *gin.RouterGroup, cfg *config.Config) {
	dashboard := r.Group("/dashboard/admin")
	// dashboard.Use(
	// 	middleware.AuthJwtMiddleware(cfg),
	// )
	{
		dashboard.GET("", h.GetDashboardAdmin)
	}
}
