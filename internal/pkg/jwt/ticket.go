package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *Service) GenerateTicketToken(
	ticketID uint64,
	expiry time.Duration,
) (string, error) {

	claims := TicketClaims{
		TicketID: ticketID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.secret)
}

func (s *Service) ValidateTicketToken(tokenString string) (*TicketClaims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&TicketClaims{},
		func(t *jwt.Token) (any, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			return s.secret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TicketClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid ticket token")
	}

	return claims, nil
}