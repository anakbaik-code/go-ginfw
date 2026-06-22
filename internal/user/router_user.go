package user

import "github.com/gin-gonic/gin"

func (h *HandlerUser) RoutesUser(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("", h.Register)
		users.GET("", h.ListUser)
		users.PUT("/:id", h.UpdateUserProfile)
	}
}
