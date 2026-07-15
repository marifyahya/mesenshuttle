package apperrors

import "errors"

var (
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrTokenGeneration     = errors.New("failed to generate token")
	ErrJWTSecretNotSet     = errors.New("JWT_SECRET is not configured")
	ErrValidation          = errors.New("validation failed")
	ErrInternalServerError = errors.New("internal server error")
	ErrRouteNotFound       = errors.New("route not found")
)
