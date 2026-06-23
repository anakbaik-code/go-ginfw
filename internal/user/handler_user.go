package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	service UserService
}

func NewHandleUser(service UserService) *HandlerUser {
	return &HandlerUser{
		service: service,
	}
}

func (h *HandlerUser) Register(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.service.Register(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		Address:   u.Address,
		Role:      u.Role,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
	}

	c.JSON(http.StatusCreated, res)
}

func (h *HandlerUser) ListUser(c *gin.Context) {
	var req ListUserRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, total, err := h.service.List(c.Request.Context(), int(req.Page), int(req.Limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users, "total": total})
}

func (h *HandlerUser) UpdateUserProfile(c *gin.Context) {
	// parsing
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req UpdateUserProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := User{
		ID:      id,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	err = h.service.UpdateProfile(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully updated",
	})

}
func (h *HandlerUser) GetActiveUsers(c *gin.Context) {
	users, err := h.service.GetActiveUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *HandlerUser) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, accessToken, refreshToken, err := h.service.Login(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	finalResponse := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: UserResponse{
			ID:       u.ID,
			Name:     u.Name,
			Email:    u.Email,
			Role:     u.Role,
			IsActive: u.IsActive,
		},
	}

	c.JSON(http.StatusOK, finalResponse)
}

func (h *HandlerUser) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := h.service.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	resUser := UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		Address:   u.Address,
		Role:      u.Role,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	c.JSON(http.StatusOK, resUser)
}
