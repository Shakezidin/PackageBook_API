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

	if err != nil {
		log.Printf("error parsing token: %v", err)
		return "", "", errors.New("error parsing token")
	}
	if !parserToken.Valid {
		log.Print("invalid token")
		return "", "", errors.New("token invalid")
	}

	expTime := claims.ExpiresAt
	if expTime < time.Now().Unix() {
		log.Print("token Expired")
		return "", "", errors.New("token expired")
	}

	userRole := claims.Role
	if userRole != role {
		log.Println("unauthorized user")
		return "", "", errors.New("unauthorized user")
	}

	return claims.Email, claims.Id, nil

}
