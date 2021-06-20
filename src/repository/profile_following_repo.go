package repository

import (
	"FollowService/domain"
	"FollowService/dto"
	"context"
	"errors"
	logger "github.com/jelena-vlajkov/logger/logger"
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
	RemoveFollowing(ctx context.Context, userToUnfollow string, userUnfollowing string) error
	AlreadyFollowing(ctx context.Context, following *domain.ProfileFollowing) bool
	ExistsProfileIds(ctx context.Context, following *domain.ProfileFollowing) error
}

type followingRepo struct {
	collection *mongo.Collection
	db *mongo.Client
	logger *logger.Logger
}

func (f *followingRepo) RemoveFollowing(ctx context.Context, userToUnfollow string, userUnfollowing string) error {
	var following bson.M

	followingBson := bson.M{"_id" :userToUnfollow}

	userBson := bson.M{"_id" : userUnfollowing}

	err := f.collection.FindOne(ctx, bson.M{"user": userBson, "following":followingBson}).Decode(&following)
	if err !=nil{
		f.logger.Logger.Errorf("error while finding one and decoding, %v\n", err)
		//log.Fatal(err)
		return err
	}

	var fusrodah dto.ProfileFollowingDTO

	bsonBytes, _ := bson.Marshal(following)
	err = bson.Unmarshal(bsonBytes, &fusrodah)
	if err != nil {
		f.logger.Logger.Errorf("error while unmarshaling, %v\n", err)
		return err
	}

	result := f.Delete(ctx, fusrodah.ID)

	if result.DeletedCount==1{
		return nil
	}else{
		f.logger.Logger.Errorf("eno followings deleted, %v\n", err)
		err1 := errors.New("deleting error: no followings deleted")
		return err1
	}}

func (f *followingRepo) ExistsProfileIds(ctx context.Context, following *domain.ProfileFollowing) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var val *domain.ProfileFollowing
	return f.collection.FindOne(ctx, bson.M{"user._id": following.User.ID, "following._id": following.Following.ID}).Decode(&val)
}

func (f *followingRepo) AlreadyFollowing(ctx context.Context, following *domain.ProfileFollowing) bool {
	err := f.ExistsProfileIds(ctx, following)

	if err != nil {
		f.logger.Logger.Errorf("existing in profile error %v\n", err)
		return false
	}

	return true
}


func (f followingRepo) GetAllUsersFollowings(user dto.ProfileDTO) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterCursor, err := f.collection.Find(ctx, bson.M{"user": user})

	if err != nil {
		f.logger.Logger.Errorf("error while finding, %v\n", err)
		//log.Fatal(err)
		return nil, err
	}
	var usersFollowingBson []bson.M
	if err = filterCursor.All(ctx, &usersFollowingBson); err != nil {
		f.logger.Logger.Errorf("error while decoding all, %v\n", err)
		//log.Fatal(err)
		return nil, err
	}

	return usersFollowingBson, nil
}

func (f followingRepo) CreateFollowing(following *domain.ProfileFollowing) (*domain.ProfileFollowing, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := f.collection.InsertOne(ctx, *following)

	if err != nil {
		f.logger.Logger.Errorf("error while inserting following, %v\n", err)
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
			f.logger.Logger.Errorf("error while deleting, %v\n", err)
			return result
		}
		result, err = f.collection.DeleteOne(ctx, bson.M{"_id" :objID})

	}

	if err != nil {
		f.logger.Logger.Errorf("error while deleting, %v\n", err)
		log.Fatal("DeleteOne() ERROR:", err)
	}

	return result
}
func NewFollowingRepo(db *mongo.Client, logger *logger.Logger) FollowingRepo {
	return &followingRepo{
		db: db,
		collection : db.Database("follow_db").Collection("profile_followings"),
		logger: logger,
	}
}