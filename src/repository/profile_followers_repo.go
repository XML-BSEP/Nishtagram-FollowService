package repository

import (
	"FollowService/domain"
	"FollowService/dto"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type FollowerRepo interface {
	CreateFollower(follower *domain.ProfileFollower) (*domain.ProfileFollower, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) *mongo.DeleteResult
	GetAllUsersFollowers(user dto.ProfileDTO) ([]bson.M, error)
}

type followerRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}

func (f *followerRepo) CreateFollower(follower *domain.ProfileFollower) (*domain.ProfileFollower, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := f.collection.InsertOne(ctx, *follower)

	if err != nil {
		panic(err)
	}

	return follower, nil
}
func (f followerRepo) GetByID(id string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := f.collection.FindOne(ctx, bson.M{"_id": id})
	return result
}

func (f followerRepo) Delete(id string) *mongo.DeleteResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := f.collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		log.Fatal("DeleteOne() ERROR:", err)
	}

	return result
}


func (f followerRepo) GetAllUsersFollowers(user dto.ProfileDTO) ([]bson.M, error){
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filterCursor, err := f.collection.Find(ctx, bson.M{"user": user})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var usersFollowersBson []bson.M
	if err = filterCursor.All(ctx, &usersFollowersBson); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return usersFollowersBson, nil
}

func NewFollowerRepo(db *mongo.Client) FollowerRepo {
	return &followerRepo{
		db: db,
		collection : db.Database("follow_db").Collection("profile_followers"),

	}
}