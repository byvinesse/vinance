package entity

import (
	"time"
)

type User struct {
	ID           string     `json:"id" db:"id"`
	Email        string     `json:"email" db:"email"`
	PasswordHash string     `json:"password_hash" db:"password_hash"`
	Username     string     `json:"username" db:"username"`
	PhoneNumber  string     `json:"phone_number" db:"phone_number"`
	DateOfBirth  *time.Time `json:"date_of_birth" db:"date_of_birth"`
	Gender       string     `json:"gender" db:"gender"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}
