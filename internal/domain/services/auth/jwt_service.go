package auth

import "github.com/stasdashkevitch/crypto_info/pkg/jwt"

type JWTAuthService struct {
	jwtService jwt.JWTAuthService
}

func (s *JWTAuthService) GenerateToken(userID string) (string, error) {
	return s.jwtService.GenerateToken(userID)
}

func (s *JWTAuthService) ValidateToken(token string) (string, error) {
	return s.jwtService.ValidateToken(token)
}
