package stex

import (
	"fmt"
)

// APIError define API error when response status is 4xx or 5xx
type APIError struct {
	Success bool   `json:"success"`
	Message string `json:"msg"`
}

// Error return error code and message
func (e APIError) Error() string {
	return fmt.Sprintf("<APIError> success=%t, msg=%s", e.Success, e.Message)
}

// IsAPIError check if e is an API error
func IsAPIError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}
