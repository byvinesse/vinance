package model

import "time"

type RegisterRequest struct {
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phone_number"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
