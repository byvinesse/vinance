package entity

import "time"

type Auth struct {
	ID          string `bson:"id"`
	Email       string `bson:"email"`
	Password    string `bson:"password"`
	PhoneNumber string `bson:"phone_number"`
	CreatedUpdated
}

type CreatedUpdated struct {
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
