package model

import "time"

type LoginResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"expires_at"`
}
