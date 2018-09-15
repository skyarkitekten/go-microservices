package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	pb "github.com/skyarkitekten/go-microservices/user-service/proto/auth"
)

var (
	key = []byte("ThatSecretKeyThoLOL")
)

// CustomClaims for JWT tokens
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

// Authable is an interface for token Service
type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

// TokenService encodes and decodes tokens
type TokenService struct {
	repo Repository
}

// Decode decodes tokens
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// Encode encodes tokens
func (srv *TokenService) Encode(user *pb.User) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "user.service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
