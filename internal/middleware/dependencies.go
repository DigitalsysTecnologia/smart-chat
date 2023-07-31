package middleware

type tokenProvider interface {
	ValidateToken(token string) error
}
