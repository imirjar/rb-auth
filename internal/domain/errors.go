package domain

// Errors
var (
	ErrInvalidCredentials = NewAuthError("invalid credentials")
	ErrTokenExpired       = NewAuthError("token expired")
	ErrInvalidToken       = NewAuthError("invalid token")
	ErrInvalidClaims      = NewAuthError("invalid claims")
)

// AuthError - кастомная ошибка аутентификации
type AuthError struct {
	message string
}

func NewAuthError(message string) *AuthError {
	return &AuthError{message: message}
}

func (e *AuthError) Error() string {
	return e.message
}
