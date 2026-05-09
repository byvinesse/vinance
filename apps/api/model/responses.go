package model

import (
	"time"

	"github.com/byvinesse/vinance-backend/entity"
)

type LoginResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"expires_at"`
}

type GetProfileResponse struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	DateOfBirth int64  `json:"date_of_birth"`
}

type GetCompleteCategoriesResponse struct {
	CategoryID      string `json:"category_id"`
	SubCategoryID   string `json:"subcategory_id"`
	CategoryName    string `json:"category_name"`
	SubCategoryName string `json:"subcategory_name"`
}

type RecordResponse struct {
	ID            string               `json:"id"`
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
	RecordedAt    time.Time            `json:"recorded_at"`
	CreatedAt     time.Time            `json:"created_at"`
}

type PaginatedRecordsResponse struct {
	Records    []RecordResponse `json:"records"`
	NextCursor *string          `json:"next_cursor"` // nil when there are no more pages
	Limit      int              `json:"limit"`
}
