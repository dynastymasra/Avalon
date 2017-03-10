package config

import "errors"

var (
	// ErrorNotFound type
	ErrorNotFound = errors.New("Endpoint or data your requested not found")
	// ErrDatabase databse
	ErrDatabase = errors.New("Database have error")
	// ErrNotSupportedHeader header
	ErrNotSupportedHeader = errors.New("Required request Header is missing")
)
