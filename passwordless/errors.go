package passwordless

import (
	"fmt"
	"strings"
)

// APIErrorResponse represents an error response from the API.
//
// It contains fields for the error type, title, status, and error code.
//
// Errors Docs: https://docs.passwordless.dev/guide/errors.html#errors
type APIErrorResponse struct {
	Type      string `json:"type"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	ErrorCode string `json:"errorCode"`
	Detail    string `json:"detail"`
}

// APIErrorResponse is compatible with the built-in error interface.
func (e APIErrorResponse) Error() string {
	var errorFields []string

	if e.Type != "" {
		errorFields = append(errorFields, fmt.Sprintf("Type: %s", e.Type))
	}
	if e.Title != "" {
		errorFields = append(errorFields, fmt.Sprintf("Title: %s", e.Title))
	}
	if e.Status != 0 {
		errorFields = append(errorFields, fmt.Sprintf("Status: %d", e.Status))
	}
	if e.ErrorCode != "" {
		errorFields = append(errorFields, fmt.Sprintf("ErrorCode: %s", e.ErrorCode))
	}
	if e.Detail != "" {
		errorFields = append(errorFields, fmt.Sprintf("Detail: %s", e.Detail))
	}

	return strings.Join(errorFields, ", ")
}
