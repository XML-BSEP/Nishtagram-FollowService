package repository

import (
	"FollowService/domain"
	"FollowService/dto"
	"context"
	"github.com/google/uuid"
	logger "github.com/jelena-vlajkov/logger/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type FollowRequestRepo interface {
	CreateFollowRequest(req *domain.FollowRequest) (*domain.FollowRequest, error)
	GetByID(ctx context.Context,id string) *mongo.SingleResult
	Delete(id string, ctx context.Context) (*mongo.DeleteResult, error)
	GetAllUsersFollowRequests(user dto.ProfileDTO) ([]bson.M, error)
	GetFollowRequestByUserAndFollower(ctx context.Context, req dto.FollowRequestDTO) (bson.M, error)
	IsCreated(ctx context.Context, request *domain.FollowRequest) bool
	ExistsProfileIds(ctx context.Context, following *domain.FollowRequest) error
	GetFollowRequest(ctx context.Context, following *dto.FollowRequestDTO) (*domain.FollowRequest, error)
}

type followRequestRepo struct {
	collection *mongo.Collection
	db *mongo.Client
	logger *logger.Logger
}

func (f *followRequestRepo) GetFollowRequest(ctx context.Context, following *dto.FollowRequestDTO) (*domain.FollowRequest, error) {
	var val *domain.FollowRequest
	err := f.collection.FindOne(ctx, bson.M{"user_requested._id": following.UserRequested, "followed_account._id": following.FollowedAccount}).Decode(&val)
	if err!=nil{
		f.logger.Logger.Errorf("error while finding one and decoding, %v\n", err)
		return nil, err
	}else{
		return val,err
	}
}

func (f *followRequestRepo) ExistsProfileIds(ctx context.Context, following *domain.FollowRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var val *domain.FollowRequest
	return f.collection.FindOne(ctx, bson.M{"user_requested._id": following.UserRequested.ID, "followed_account._id": following.FollowedAccount.ID}).Decode(&val)
}

func (f *followRequestRepo) IsCreated(ctx context.Context, request *domain.FollowRequest) bool {
	err := f.ExistsProfileIds(ctx, request)

	if err != nil {
		f.logger.Logger.Errorf("exists profile failed, %v\n", err)
		return false
	}

	return true
}

func (f followRequestRepo) GetFollowRequestByUserAndFollower(ctx context.Context, req dto.FollowRequestDTO) (bson.M, error) {
	//panic("implement me")
	var reqBson bson.M
	followedBson := bson.M{"_id" :req.FollowedAccount}

	userRequestedBson := bson.M{"_id" : req.UserRequested}

	err := f.collection.FindOne(ctx, bson.M{"followed_account": followedBson, "user_requested":userRequestedBson}).Decode(&reqBson)
	if err != nil{
		f.logger.Logger.Errorf("error while finding one and decoding, %v\n", err)
		return nil, err
	}
	return reqBson,nil
	//var request dto.FollowRequestDTO
	//
	//bsonBytes, _ := bson.Marshal(reqBson)
	//err = bson.Unmarshal(bsonBytes, &request)
	//if err != nil {
	//	return nil, err
	//}

}

func (f followRequestRepo) GetAllUsersFollowRequests(user  dto.ProfileDTO) ( []bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filterCursor, err := f.collection.Find(ctx, bson.M{"followed_account": user})

	if err != nil {
		f.logger.Logger.Errorf("error while finding, %v\n", err)
		//log.Fatal(err)
		return nil, err
	}
	var usersFollowRequestsBson []bson.M
	if err = filterCursor.All(ctx, &usersFollowRequestsBson); err != nil {
		f.logger.Logger.Errorf("error while decoding all users follow requests, %v\n", err)
		//log.Fatal(err)
		return nil, err
	}

	return usersFollowRequestsBson, nil
}

func (f followRequestRepo) CreateFollowRequest(req *domain.FollowRequest) (*domain.FollowRequest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	req.ID = uuid.NewString()
	_, err := f.collection.InsertOne(ctx, *req)

	if err != nil {
		f.logger.Logger.Errorf("inster one failed  %v\n", err)
		//panic(err)
	}
	return req, nil
}

func (f followRequestRepo) GetByID(ctx context.Context,id string) *mongo.SingleResult {
	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//defer cancel()

	result :=f.collection.FindOne(ctx, bson.M{"_id": id})

	return result
}

func (f followRequestRepo) Delete(id string, ctx context.Context) (*mongo.DeleteResult, error) {
	//oid, err := primitive.ObjectIDFromHex(id)

	result, err := f.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		f.logger.Logger.Errorf("delete one failed, %v\n", err)
		//log.Fatal("DeleteOne() ERROR:", err)
		return nil, err
	}
	return result, nil
}

func NewFollowRequestRepo(db *mongo.Client, logger *logger.Logger) FollowRequestRepo {
	return &followRequestRepo{
		db: db,
		collection : db.Database("follow_db").Collection("follow_requests"),
		logger: logger,
	}
}