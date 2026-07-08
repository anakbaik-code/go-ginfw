package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *Service) GenerateRefreshToken(
	userID uint64,
	expiry time.Duration,
) (string, error) {

	claims := jwt.RegisteredClaims{
		Subject:   string(rune(userID)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.secret)
}

func (s *Service) ValidateRefreshToken(tokenString string) (*jwt.RegisteredClaims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return s.secret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	return claims, nil
}