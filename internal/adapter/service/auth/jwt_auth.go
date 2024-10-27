package auth

import (
	"os"

	"github.com/stasdashkevitch/crypto_info/pkg/jwt"
)

type JWTAuth struct {
	jwtService jwt.JWTAuthService
}

func NewJWTAuth() *JWTAuth {
	secretKey := os.Getenv("CRYPTO_INFO_JWT_SECRET_KEY")
	return &JWTAuth{
		jwtService: *jwt.NewJWTAuthService(secretKey),
	}
}

func (s *JWTAuth) GenerateToken(userID string) (string, error) {
	return s.jwtService.GenerateToken(userID)
}

func (s *JWTAuth) ValidateToken(token string) (string, error) {
	return s.jwtService.ValidateToken(token)
}
