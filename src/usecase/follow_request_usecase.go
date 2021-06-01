package usecase

import (
	"FollowService/domain"
	"FollowService/dto"
	"FollowService/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FollowRequestUseCase interface {
	CreateFollowRequest(req *domain.FollowRequest) (*domain.FollowRequest, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) *mongo.DeleteResult
	GetAllUsersFollowRequests(user dto.ProfileDTO) ([]*domain.Profile, error)

}

type followRequestUseCase struct {
	FollowRequestRepo repository.FollowRequestRepo
}

func (f followRequestUseCase) GetAllUsersFollowRequests(user dto.ProfileDTO) ([]*domain.Profile, error) {
	userFollowRequestsBson, err := f.FollowRequestRepo.GetAllUsersFollowRequests(user)
	if err !=nil{
		return nil,err
	}
	var usersFollowRequests []*domain.Profile
	var followRequest *domain.FollowRequest
	for _, uf := range userFollowRequestsBson {
		bsonBytes, _ := bson.Marshal(uf)
		err := bson.Unmarshal(bsonBytes, &followRequest)
		if err != nil {
			return nil, err
		}
		usersFollowRequests = append(usersFollowRequests, &followRequest.UserRequested)
	}
	return usersFollowRequests, nil
}

func (f followRequestUseCase) CreateFollowRequest(req *domain.FollowRequest) (*domain.FollowRequest, error) {
	return f.FollowRequestRepo.CreateFollowRequest(req)
}

func (f followRequestUseCase) GetByID(id string) *mongo.SingleResult {
	return f.FollowRequestRepo.GetByID(id)
}

func (f followRequestUseCase) Delete(id string) *mongo.DeleteResult {
	return f.FollowRequestRepo.Delete(id)
}

func NewFollowRequestUseCase(repo repository.FollowRequestRepo) FollowRequestUseCase {
	return &followRequestUseCase{
		FollowRequestRepo: repo,
	}
}