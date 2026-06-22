package user

import "time"

type User struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash string
	Phone        string
	Address      string
	Role         string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
