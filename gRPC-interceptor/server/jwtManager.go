package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	// DeaultTokenDuration Default token duration
	DeaultTokenDuration = time.Duration(time.Minute * 15)
)

type userClaims struct {
	jwt.StandardClaims
	Email string
	Roles []string
}

// JWTManager ...
type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

// Create a JWTManager struct
func newJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	secretKey = strings.TrimSpace(secretKey)
	if tokenDuration <= 0 {
		tokenDuration = DeaultTokenDuration
	}

	return &JWTManager{
		secretKey:     secretKey,
		tokenDuration: tokenDuration,
	}
}

// Generate JWT access token
func (manager *JWTManager) generate(user *User) (string, error) {
	claims := &userClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		Email: user.Email,
		Roles: user.Roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString([]byte(manager.secretKey))
	if err != nil {
		return "", nil
	}

	return stringToken, nil
}

// Verify the access token that passed in and returns the userClaims sturct
func (manager *JWTManager) verify(accessToken string) (*userClaims, error) {
	t, err := jwt.ParseWithClaims(
		accessToken,
		&userClaims{},
		func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signature method")
			}
			return []byte(manager.secretKey), nil
		})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := t.Claims.(*userClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
