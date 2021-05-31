package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ProfileRepo interface {
	GetByID(id string) *mongo.SingleResult
}

type profileRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}

func (p profileRepo) GetByID(id string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := p.collection.FindOne(ctx, bson.M{"_id": id})
	return result
}

func NewProfileRepo(db *mongo.Client) ProfileRepo {
	return &profileRepo{
		db: db,
		collection : db.Database("follow_db").Collection("profiles"),
	}
}