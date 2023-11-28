package model

type RegisterResponse struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	IsMember    bool   `json:"is_member"`
}

type LoginResponse struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	IsMember    bool   `json:"is_member"`
}
