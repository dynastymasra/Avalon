package config

import "errors"

var (
	// ErrorNotFound type
	ErrorNotFound = errors.New("Endpoint or data your requested not found")
	// ErrDatabase databse
	ErrDatabase = errors.New("Database have error")
	// ErrNotSupportedHeader header
	ErrNotSupportedHeader = errors.New("Required request Header is missing")
	// ErrorRecordNotFound type
	ErrorRecordNotFound = errors.New("Record or data your requested not found")
	// ErrorNotValidUUID error
	ErrorNotValidUUID = errors.New("Id order is invalid")
	// Error not Acknowledged index elastic
	ErrorNotAcknowledgedIndex = errors.New("Create index not acknowledged")
)
