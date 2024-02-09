package middleware

import (
	"errors"
	"log"
	"time"

	"github.com/Shakezidin/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Id    string
	Email string
	Role  string
	Token string
	jwt.StandardClaims
}

type JwtClaims struct {
	cfg *config.Configure
}

var jwtKey = []byte("SECRETKEY")

func ValidateToken(ctx *gin.Context, role string) (string, string, error) {
	headerToken := ctx.GetHeader("Authorization")
	if headerToken == "" {
		log.Print("bearer token missing")
		return "", "", errors.New("bearer token missing")
	}

	claims := &Claims{}
	token := string([]byte(headerToken)[7:])
	parserToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil || !parserToken.Valid {
		log.Print("invalid token")
		return "", "", errors.New("invalid token")
	}

	if claims.Token != "Access" {
		log.Println("unauthorized user")
		return "", "", errors.New("not an access token")
	}

	userRole := claims.Role
	if userRole != role {
		log.Println("unauthorized user")
		return "", "", errors.New("unauthorized user")
	}

	expTime := claims.ExpiresAt
	if expTime < time.Now().Unix() {
		log.Print("token expired")
		return "", "", errors.New("token expired")
	}

	return claims.Email, claims.Id, nil
}

func ValidateRefreshToken(c *gin.Context, role string) (string, error) {
	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	tokenReq := tokenReqBody{}
	if err := c.Bind(&tokenReq); err != nil {
		return "", err
	}
	claims := &Claims{}
	parserToken, err := jwt.ParseWithClaims(tokenReq.RefreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil || !parserToken.Valid {
		log.Print("invalid token")
		return "", errors.New("invalid token")
	}

	if claims.Token != "Refresh" {
		log.Println("unauthorized user")
		return "", errors.New("not an refresh token")
	}

	expTime := claims.ExpiresAt
	if expTime < time.Now().Unix() {
		log.Print("token expired")
		return "", errors.New("token expired")
	}

	return claims.Id, nil

}
