package errors

import "errors"

// User-related errors
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailInvalid       = errors.New("email is invalid")
	ErrPasswordInvalid    = errors.New("password is invalid")
	ErrUserAlreadyyExists = errors.New("user already exists")
)

// Auth-related errors
var (
	ErrUnauthorized  = errors.New("unauthorized access")
	ErrTokenInvalid  = errors.New("invalid token")
	ErrSigningMethod = errors.New("unexpected signing method")
	ErrExpToken      = errors.New("token has expired")
	ErrInvalidToken  = errors.New("invalid token")
)
