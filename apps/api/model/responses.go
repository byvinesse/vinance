package model

import "time"

type RegisterResponse struct {
	Email    string `json:"email"`
	IsMember bool   `json:"is_member"`
}

type LoginResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"expires_at"`
}
