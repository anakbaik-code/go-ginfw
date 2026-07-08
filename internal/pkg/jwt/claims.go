package jwt

import "github.com/golang-jwt/jwt/v5"

type AccessClaims struct {
	UserID uint64 `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type TicketClaims struct {
	TicketID uint64 `json:"ticket_id"`
	jwt.RegisteredClaims
}