package db

import (
	"context"

	"github.com/vincentkdeli/vinance-backend/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Member struct {
	db *mongo.Database
}

func NewMember(db *mongo.Database) *Member {
	return &Member{
		db: db,
	}
}

func (r *Member) InsertOne(ctx context.Context, member *entity.Member) (*entity.Member, error) {
	res, err := r.db.Collection(entity.TableNameMember).InsertOne(ctx, member)
	if err != nil {
		return nil, err
	}

	var newMember *entity.Member
	err = r.db.Collection(entity.TableNameMember).FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&newMember)
	if err != nil {
		return nil, err
	}

	return newMember, nil
}
