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

type FollowRequestRepo interface {
	CreateFollowRequest(req *domain.FollowRequest) (*domain.FollowRequest, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) (*mongo.DeleteResult, error)
	GetAllUsersFollowRequests(user dto.ProfileDTO) ([]bson.M, error)
	GetFollowRequestByUserAndFollower(ctx context.Context, req dto.FollowRequestDTO) (bson.M, error)
}

type followRequestRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}

func (f followRequestRepo) GetFollowRequestByUserAndFollower(ctx context.Context, req dto.FollowRequestDTO) (bson.M, error) {
	//panic("implement me")
	var reqBson bson.M
	followedBson := bson.M{"_id" :req.FollowedAccount}

	userRequestedBson := bson.M{"_id" : req.UserRequested}

	err := f.collection.FindOne(ctx, bson.M{"followed_account": followedBson, "user_requested":userRequestedBson}).Decode(&reqBson)
	if err != nil{
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
		log.Fatal(err)
		return nil, err
	}
	var usersFollowRequestsBson []bson.M
	if err = filterCursor.All(ctx, &usersFollowRequestsBson); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return usersFollowRequestsBson, nil
}

func (f followRequestRepo) CreateFollowRequest(req *domain.FollowRequest) (*domain.FollowRequest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := f.collection.InsertOne(ctx, *req)

	if err != nil {
		panic(err)
	}
	return req, nil
}

func (f followRequestRepo) GetByID(id string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result :=f.collection.FindOne(ctx, bson.M{"_id": id})
	return result
}

func (f followRequestRepo) Delete(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := f.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Fatal("DeleteOne() ERROR:", err)
		return nil, err
	}
	return result, nil
}

func NewFollowRequestRepo(db *mongo.Client) FollowRequestRepo {
	return &followRequestRepo{
		db: db,
		collection : db.Database("follow_db").Collection("follow_requests"),
	}
}