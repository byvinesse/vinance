package entity

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Auth struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	IsMember  bool      `json:"is_member" db:"is_member"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type AuthToken struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"expires_at"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
}
