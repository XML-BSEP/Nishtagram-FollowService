package repository

import (
	"FollowService/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type FollowingRepo interface {
	CreateFollowing(following *domain.ProfileFollowing) (*domain.ProfileFollowing, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) *mongo.DeleteResult
}

type followingRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}

func (f followingRepo) CreateFollowing(following *domain.ProfileFollowing) (*domain.ProfileFollowing, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := f.collection.InsertOne(ctx, *following)

	if err != nil {
		panic(err)
	}
	return following, nil
}

func (f followingRepo) GetByID(id string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := f.collection.FindOne(ctx, bson.M{"_id": id})
	return result
}

func (f followingRepo) Delete(id string) *mongo.DeleteResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := f.collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		log.Fatal("DeleteOne() ERROR:", err)
	}

	return result
}
func NewFollowingRepo(db *mongo.Client) FollowingRepo {
	return &followingRepo{
		db: db,
		collection : db.Database("follow_db").Collection("profile_followings"),
	}
}