package user

type CreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Address  string `json:"address,omitempty"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}

type ListUserRequest struct {
	Page  int32 `form:"page" binding:"gte=1"`
	Limit int32 `form:"limit" binding:"gte=1,lte=100"`
}

type UpdateUserProfileRequest struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
