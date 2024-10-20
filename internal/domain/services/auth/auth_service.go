package auth

type AuthService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (string, error)
}
