package user

import "github.com/gin-gonic/gin"

func (h *HandlerUser) RoutesUser(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("/register", h.Register)
		users.POST("/login", h.Login)
		users.GET("", h.ListUser)
		users.PUT("/:id", h.UpdateUserProfile)
		users.GET("/:id", h.GetByID)
		users.GET("/active", h.GetActiveUsers)
	}
}
