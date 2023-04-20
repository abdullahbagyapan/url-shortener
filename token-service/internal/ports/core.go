package ports

import (
	"github.com/golang-jwt/jwt"
	"url-shortener/token-service/internal/adapters/core"
)

type CorePort interface {
	GenerateNewToken(id string) (string, error)
	ValidateToken(jwt *jwt.Token) bool
	GetUserIdByClaims(claims *core.JWTData) string
	GetJWTAndClaims(token string) (*jwt.Token, *core.JWTData, error)
}
