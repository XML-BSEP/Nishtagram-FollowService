package usecase

import (
	"FollowService/domain"
	"FollowService/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type FollowerUseCase interface {
	CreateFollower(follower *domain.ProfileFollower) (*domain.ProfileFollower, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) *mongo.DeleteResult
}

type followerUseCase struct {
	FollowerRepo repository.FollowerRepo
}

func (f followerUseCase) CreateFollower(follower *domain.ProfileFollower) (*domain.ProfileFollower, error) {
	return f.FollowerRepo.CreateFollower(follower)
}

func (f followerUseCase) GetByID(id string) *mongo.SingleResult {
	return f.FollowerRepo.GetByID(id)
}

func (f followerUseCase) Delete(id string) *mongo.DeleteResult {
	return f.FollowerRepo.Delete(id)
}

func NewFollowerUseCase(repo repository.FollowerRepo) FollowerUseCase {
	return &followerUseCase{
		FollowerRepo: repo,
	}
}
