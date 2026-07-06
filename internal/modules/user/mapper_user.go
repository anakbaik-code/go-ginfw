package user

import "go-fwgin/internal/database"

// Mapper function untuk single Organizer
func ToOrganizer(row database.ListOrganizersRow) User {
	return User{
		ID:       row.ID,
		Name:     row.Name,
		Email:    row.Email,
		Phone:    row.Phone.String,
		Address:  row.Address.String,
		Role:     row.Role,
		IsActive: row.IsActive,
	}
}

// Mapper function untuk slice
func ToOrganizers(rows []database.ListOrganizersRow) []User {
	users := make([]User, 0, len(rows))
	for _, row := range rows {
		users = append(users, ToOrganizer(row))
	}
	return users
}

func ToUserResponse(user User) UserResponse {
	return UserResponse{
		ID:        uint64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
func ToUserResponses(users []User) []UserResponse {
	responses := make([]UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, ToUserResponse(user))
	}
	return responses
}
