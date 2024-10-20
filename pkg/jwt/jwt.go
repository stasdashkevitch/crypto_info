package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthService struct {
	secretKey string
}

func NewJWTAuthService(secretKey string) *JWTAuthService {
	return &JWTAuthService{secretKey: secretKey}
}

func (s *JWTAuthService) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *JWTAuthService) ValidateToken(token string) (string, error) {
	claims := &jwt.MapClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil || !t.Valid {
		return "", errors.New("invalid token")
	}

	return (*claims)["user_id"].(string), nil
}
