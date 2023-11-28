package db

import (
	"context"

	"github.com/vincentkdeli/vinance-backend/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth struct {
	db *mongo.Database
}

func NewAuth(db *mongo.Database) *Auth {
	return &Auth{
		db: db,
	}
}

func (r *Auth) InsertOne(ctx context.Context, auth *entity.Auth) (*entity.Auth, error) {
	res, err := r.db.Collection(entity.TableNameAuth).InsertOne(ctx, auth)
	if err != nil {
		return nil, err
	}

	var newAuth *entity.Auth
	err = r.db.Collection(entity.TableNameAuth).FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&newAuth)
	if err != nil {
		return nil, err
	}

	return newAuth, nil
}

func (r *Auth) FindOneByEmail(ctx context.Context, email string) (*entity.Auth, error) {
	var data *entity.Auth

	if err := r.db.Collection(entity.TableNameAuth).FindOne(ctx, bson.M{"email": email}).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
