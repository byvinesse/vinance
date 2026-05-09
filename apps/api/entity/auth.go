package entity

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthToken struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"expires_at"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
}
