package repository

import (
	"FollowService/domain"
	"FollowService/dto"
	"context"
	"errors"
	"github.com/google/uuid"
	logger "github.com/jelena-vlajkov/logger/logger"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type FollowerRepo interface {
	CreateFollower(follower *domain.ProfileFollower) (*domain.ProfileFollower, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) *mongo.DeleteResult
	GetAllUsersFollowers(user dto.ProfileDTO) ([]bson.M, error)
	AlreadyFollowing(ctx context.Context, following *domain.ProfileFollowing) (bool, error)
	RemoveFollower(ctx context.Context, userToUnfollow string, userUnfollowing string) error
	GetFollowerByFollowerAndUser(ctx context.Context, follower string, user string) (*domain.ProfileFollower, error)
	UpdateCloseFriendStatus(ctx context.Context, profileFollowerId string, status bool) error
	GetAllUsersCloseFriends(ctx context.Context, userId string) ([]bson.M, error)
	GetAllUsersToWhomUserIsCloseFriend(ctx context.Context, userId string) ([]bson.M, error)
}

type followerRepo struct {
	collection *mongo.Collection
	db *mongo.Client
	logger *logger.Logger
}

func (f *followerRepo) GetAllUsersCloseFriends(ctx context.Context, userId string) ([]bson.M, error) {
	userBson := bson.M{"_id" : userId}

	filterCursor, err := f.collection.Find(ctx, bson.M{"user": userBson, "close_friend":true})

	if err != nil {
		f.logger.Logger.Errorf("error while finding all users , %v\n", err)
		//log.Fatal(err)
		return nil, err
	}
	var usersFollowersBson []bson.M
	if err = filterCursor.All(ctx, &usersFollowersBson); err != nil {
		f.logger.Logger.Errorf("error while decoding all, %v\n", err)
		//log.Fatal(err)
		return nil, err
	}
	return usersFollowersBson, nil
}


func (f *followerRepo) GetAllUsersToWhomUserIsCloseFriend(ctx context.Context, userId string) ([]bson.M, error) {
	followerBson := bson.M{"_id" : userId}

	filterCursor, err := f.collection.Find(ctx, bson.M{"follower": followerBson, "close_friend":true})

	if err != nil {
		f.logger.Logger.Errorf("error while finding all users , %v\n", err)
		//log.Fatal(err)
		return nil, err
	}
	var usersFollowersBson []bson.M
	if err = filterCursor.All(ctx, &usersFollowersBson); err != nil {
		f.logger.Logger.Errorf("error while decoding all, %v\n", err)
		//log.Fatal(err)
		return nil, err
	}
	return usersFollowersBson, nil
}

func (f *followerRepo) UpdateCloseFriendStatus(ctx context.Context, profileFollowerId string,status bool) error {

	profileFollowerToUpdate := bson.M{"_id" : profileFollowerId}
	updatedProfileFollower := bson.M{"$set": bson.M{
		"close_friend":      status,

	}}


	_, err := f.collection.UpdateOne(ctx, profileFollowerToUpdate, updatedProfileFollower)
	if err != nil {
		f.logger.Logger.Errorf("error while updating profile follower with id %v\n", profileFollowerId)
		return  err
	}
	return nil
}

func (f *followerRepo) GetFollowerByFollowerAndUser(ctx context.Context, follower string, user string) (*domain.ProfileFollower, error) {
	var fBson bson.M
	followerBson := bson.M{"_id" : follower}

	userBson := bson.M{"_id" : user}

	err := f.collection.FindOne(ctx, bson.M{"user": userBson, "follower":followerBson}).Decode(&fBson)
	if err !=nil{
		f.logger.Logger.Errorf("error while finding one and decoding, %v\n", err)
		return nil, err
	}

	var fusrodah domain.ProfileFollower

	bsonBytes, _ := bson.Marshal(fBson)
	err = bson.Unmarshal(bsonBytes, &fusrodah)
	if err != nil {
		f.logger.Logger.Errorf("unmarshal error, %v\n", err)
		return nil, err
	}
	return &fusrodah, nil
}

func (f *followerRepo) RemoveFollower(ctx context.Context, userToUnfollow string, userUnfollowing string) error {
	var follower bson.M
	unfollowerBson := bson.M{"_id" : userUnfollowing}
	
	userBson := bson.M{"_id" : userToUnfollow}

	err := f.collection.FindOne(ctx, bson.M{"user": userBson, "follower":unfollowerBson}).Decode(&follower)

	if err !=nil{
		f.logger.Logger.Errorf("error while finding one and decoding, %v\n", err)
		return err
	}

	var fusrodah dto.ProfileFollowerDTO

	bsonBytes, _ := bson.Marshal(follower)
	err = bson.Unmarshal(bsonBytes, &fusrodah)
	if err != nil {
		f.logger.Logger.Errorf("unmarshal error, %v\n", err)
		return err
	}


	if f.Delete(fusrodah.ID).DeletedCount==1{
		return nil
	}else{
		f.logger.Logger.Errorf("eleting error: no followers deleted")
		err1 := errors.New("deleting error: no followers deleted")
		return err1
	}
}

func (f *followerRepo) AlreadyFollowing(ctx context.Context, following *domain.ProfileFollowing) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	/*userBson := bson.M{"_id" : following.User}
	followingBson := bson.M{"_id" : following.Following}*/

	var val *domain.ProfileFollowing
	err := f.collection.FindOne(ctx, bson.M{"user._id" : following.User.ID, "following._id" : following.Following.ID}).Decode(&val)

	if err != nil {
		f.logger.Logger.Errorf("error while finding one and decoding, %v\n", err)
		return false, err
	}

	return true, err
}

func (f *followerRepo) CreateFollower(follower *domain.ProfileFollower) (*domain.ProfileFollower, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	follower.ID = uuid.NewString()

	_, err := f.collection.InsertOne(ctx, *follower)

	if err != nil {
		f.logger.Logger.Errorf("inserting one error %v\n", err)
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
	if result.DeletedCount==0{
		//objID, err := primitive.ObjectIDFromHex(id)
		if err!=nil{
			f.logger.Logger.Errorf("delete error, %v\n", err)
			return result
		}
		result, err = f.collection.DeleteOne(ctx, bson.M{"_id" :id})

	}

	if err != nil {
		f.logger.Logger.Errorf("delete error count, %v\n", err)
		return &mongo.DeleteResult{DeletedCount: 0}

		//log.Fatal("DeleteOne() ERROR:", err)
	}

	return result
}


func (f followerRepo) GetAllUsersFollowers(user dto.ProfileDTO) ([]bson.M, error){
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filterCursor, err := f.collection.Find(ctx, bson.M{"user": user})

	if err != nil {
		f.logger.Logger.Errorf("error while finding all users followers, %v\n", err)
		//log.Fatal(err)
		return nil, err
	}
	var usersFollowersBson []bson.M
	if err = filterCursor.All(ctx, &usersFollowersBson); err != nil {
		f.logger.Logger.Errorf("error while decoding all, %v\n", err)
		//log.Fatal(err)
		return nil, err
	}

	return usersFollowersBson, nil
}

func NewFollowerRepo(db *mongo.Client, logger *logger.Logger) FollowerRepo {
	return &followerRepo{
		db: db,
		collection : db.Database("follow_db").Collection("profile_followers"),
		logger: logger,
	}
}