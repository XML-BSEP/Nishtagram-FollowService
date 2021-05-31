package usecase

import (
	"FollowService/domain"
	"FollowService/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type FollowRequestUseCase interface {
	CreateFollowRequest(req *domain.FollowRequest) (*domain.FollowRequest, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) *mongo.DeleteResult
}

type followRequestUseCase struct {
	FollowRequestRepo repository.FollowRequestRepo
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