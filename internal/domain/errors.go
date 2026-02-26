package domain

import "errors"

var (
	ErrInvalidServiceName = errors.New("invalid service name")
	ErrInvalidPrice       = errors.New("price must be greater than zero")
	ErrInvalidDateRange   = errors.New("end date cannot be before start date")
)
