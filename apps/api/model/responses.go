package model

import "time"

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
