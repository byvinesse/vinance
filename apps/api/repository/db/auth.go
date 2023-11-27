package db

import "go.mongodb.org/mongo-driver/mongo"

type Auth struct {
	db *mongo.Database
}

func NewAuth(db *mongo.Database) *Auth {
	return &Auth{
		db: db,
	}
}
