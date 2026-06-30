package response

import "errors"

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
}
var (
	ErrForbidden = errors.New("forbidden")
	ErrEventNotFound = errors.New("event not found")
)