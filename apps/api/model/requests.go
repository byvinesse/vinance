package model

import (
	"time"

	"github.com/byvinesse/vinance-backend/entity"
)

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

type CreateAccountRequest struct {
	Name     string              `json:"name"`
	Balance  float64             `json:"balance"`
	Currency entity.CurrencyType `json:"currency"`
	Type     entity.AccountType  `json:"type"`
	Color    string              `json:"color"`
}
