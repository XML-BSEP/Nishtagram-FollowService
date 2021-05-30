package repository

import (
	"FollowService/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type FollowerRepo interface {
	CreateFollower(follower *domain.ProfileFollowers) (*domain.ProfileFollowers, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) *mongo.DeleteResult
}

type followerRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}

func (f *followerRepo) CreateFollower(follower *domain.ProfileFollowers) (*domain.ProfileFollowers, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_, err := f.collection.InsertOne(ctx, *follower)

	if err != nil {
		panic(err)
	}

	return follower, nil
}
func (f followerRepo) GetByID(id string) *mongo.SingleResult {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	result := f.collection.FindOne(ctx, bson.M{"_id": id})
	return result
}

func (f followerRepo) Delete(id string) *mongo.DeleteResult {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	result, err := f.collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		log.Fatal("DeleteOne() ERROR:", err)
	}

	return result
}
func NewFollowerRepo(db *mongo.Client) FollowerRepo {
	return &followerRepo{
		db: db,
		collection : db.Database("follow_db").Collection("profile_followers"),

	}
}