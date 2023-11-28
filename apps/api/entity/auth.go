package entity

import (
	"time"
)

type Auth struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Email       string    `json:"email" bson:"email"`
	Password    string    `json:"password" bson:"password"`
	PhoneNumber string    `json:"phone_number" bson:"phone_number"`
	IsMember    bool      `json:"is_member" bson:"is_member"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
