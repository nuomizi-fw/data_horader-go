package middleware

import (
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nuomizi-fw/data_horader-go/core"
)

type JWTMiddleware struct {
	config       core.DataHoraderConfig
	logger       core.DataHoraderLogger
	data_horader core.DataHoraderServer
}

func NewJWTMiddleware(
	config core.DataHoraderConfig,
	logger core.DataHoraderLogger,
	data_horader core.DataHoraderServer,
) JWTMiddleware {
	return JWTMiddleware{config, logger, data_horader}
}

func (jm JWTMiddleware) InitMiddleware() {
	jm.logger.Info("Initializing JWT middleware")

	jm.data_horader.Api.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: jm.config.Server.JWT.Secret,
		},
	}))
}

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func generateToken(secret string, issuer string, id int, username string) (string, error) {
	claims := Claims{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func parseToken(secret string, tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}

func GenerateAccessToken(secret string, id int, username string) (string, error) {
	return generateToken(secret, "data_horader-access", id, username)
}

func GenerateRefreshToken(secret string, id int, username string) (string, error) {
	return generateToken(secret, "staragazer-refresh", id, username)
}

func ValidateToken(secret string, tokenString string) (*Claims, error) {
	claims, err := parseToken(secret, tokenString)
	if err != nil {
		return nil, err
	}

	if claims == nil {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
