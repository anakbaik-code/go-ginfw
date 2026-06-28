package user

import "time"

type CreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type UserResponse struct {
	ID        uint64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone,omitempty"`
	Address   string    `json:"address,omitempty"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListUserRequest struct {
	Page  int32 `form:"page" binding:"gte=1"`
	Limit int32 `form:"limit" binding:"gte=1"`
}

type UpdateUserProfileRequest struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gt=3,lte=50" `
}
type LoginResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         UserResponse `json:"user"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}
