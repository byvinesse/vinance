package model

import (
	"time"

	"github.com/byvinesse/vinance-backend/entity"
)

type RegisterRequest struct {
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Username    string     `json:"username"`
	PhoneNumber string     `json:"phone_number"`
	Gender      string     `json:"gender"`
	DateOfBirth *time.Time `json:"date_of_birth"`
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

type CreateRecordRequest struct {
	AccountID     string               `json:"account_id"`
	SubCategoryID string               `json:"subcategory_id"`
	Amount        float64              `json:"amount"`
	Currency      entity.CurrencyType  `json:"currency"`
	BaseAmount    float64              `json:"base_amount"`
	Type          entity.RecordType    `json:"type"`
	Labels        []string             `json:"labels"`
	Name          string               `json:"name"`
	Payee         string               `json:"payee"`
	PaymentType   entity.PaymentType   `json:"payment_type"`
	PaymentStatus entity.PaymentStatus `json:"payment_status"`
	IsExcluded    bool                 `json:"is_excluded"`
	RecordedAt    *time.Time           `json:"recorded_at"`
}
