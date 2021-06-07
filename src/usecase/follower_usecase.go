package usecase

import (
	"FollowService/domain"
	"FollowService/dto"
	"FollowService/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FollowerUseCase interface {
	CreateFollower(follower *domain.ProfileFollower) (*domain.ProfileFollower, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) *mongo.DeleteResult
	GetAllUsersFollowers(user dto.ProfileDTO) ([]*domain.Profile, error)
	AlreadyFollowing (ctx context.Context, following *domain.ProfileFollowing) (bool, error)
}

type followerUseCase struct {
	FollowerRepo repository.FollowerRepo
}

func (f followerUseCase) GetAllUsersFollowers(user dto.ProfileDTO) ([]*domain.Profile, error) {
	userFollowersBson, err :=  f.FollowerRepo.GetAllUsersFollowers(user)
	if err !=nil{
		return nil, err
	}

	var usersFollowers []*domain.Profile
	for _, uf := range userFollowersBson {
		bsonBytes, _ := bson.Marshal(uf)
		var follower *domain.ProfileFollower

		err := bson.Unmarshal(bsonBytes, &follower)
		if err != nil {
			return nil, err
		}
		usersFollowers = append(usersFollowers, &follower.Follower)
	}

	return usersFollowers,nil
}


func (f *followerUseCase) AlreadyFollowing(ctx context.Context, following *domain.ProfileFollowing) (bool, error) {
	return f.FollowerRepo.AlreadyFollowing(ctx, following)
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
