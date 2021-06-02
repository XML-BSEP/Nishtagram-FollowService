package repository

import (
	"FollowService/domain"
	"FollowService/dto"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type FollowingRepo interface {
	CreateFollowing(following *domain.ProfileFollowing) (*domain.ProfileFollowing, error)
	GetByID(id string) *mongo.SingleResult
	Delete(ctx context.Context, id string) *mongo.DeleteResult
	GetAllUsersFollowings(user dto.ProfileDTO) ([]bson.M, error)
	RemoveFollowing(ctx context.Context, unfollow dto.Unfollow) error

}

type followingRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}

func (f followingRepo) RemoveFollowing(ctx context.Context, unfollow dto.Unfollow) error {
	var following bson.M
	fmt.Println(unfollow)
	followingBson := bson.M{"_id" :unfollow.UserUnfollowing}

	userBson := bson.M{"_id" : unfollow.UserToUnfollow}

	// ja pratim peru i hocu peru da otpratim
	// u ovom slucaju ja sam userUnfollowing, a pera je UserToUnfollow

	//sto znaci da ja trebam da budem obrisan iz tabele kao follower
	//a pera kao user


	err := f.collection.FindOne(ctx, bson.M{"user": userBson, "following":followingBson}).Decode(&following)
	if err !=nil{
		log.Fatal(err)
		return err
	}

	var fusrodah dto.ProfileFollowingDTO

	bsonBytes, _ := bson.Marshal(following)
	err = bson.Unmarshal(bsonBytes, &fusrodah)
	if err != nil {
		return err
	}

	result := f.Delete(ctx, fusrodah.ID)

	if result.DeletedCount==1{
		return nil
	}else{
		err1 := errors.New("deleting error: no followings deleted")
		return err1
	}

}

func (f followingRepo) GetAllUsersFollowings(user dto.ProfileDTO) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterCursor, err := f.collection.Find(ctx, bson.M{"user": user})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var usersFollowingBson []bson.M
	if err = filterCursor.All(ctx, &usersFollowingBson); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return usersFollowingBson, nil
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

func (f followingRepo) Delete(ctx context.Context, id string) *mongo.DeleteResult {

	result, err := f.collection.DeleteOne(ctx, bson.M{"_id" :id})

	if result.DeletedCount==0{
		objID, err := primitive.ObjectIDFromHex(id)
		if err!=nil{
			return result
		}
		result, err = f.collection.DeleteOne(ctx, bson.M{"_id" :objID})

	}

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