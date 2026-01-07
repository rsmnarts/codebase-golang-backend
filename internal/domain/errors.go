package domain

import "errors"

var (
	// ErrNotFound is returned when a resource is not found
	ErrNotFound = errors.New("resource not found")
	
	// ErrInvalidInput is returned when input validation fails
	ErrInvalidInput = errors.New("invalid input")
	
	// ErrInternalServer is returned for internal server errors
	ErrInternalServer = errors.New("internal server error")
)
