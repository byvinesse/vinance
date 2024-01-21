package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email       string             `json:"email" bson:"email"`
	AccountID   string             `json:"account_id" bson:"account_id"`
	Username    string             `json:"username" bson:"username"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	DateOfBirth time.Time          `json:"date_of_birth" bson:"date_of_birth"`
	Gender      string             `json:"gender" bson:"gender"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
