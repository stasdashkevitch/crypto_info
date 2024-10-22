package auth

type Auth interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (string, error)
}
