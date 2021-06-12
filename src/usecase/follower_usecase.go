package usecase

import (
	"FollowService/domain"
	"FollowService/dto"
	"FollowService/gateway"
	"FollowService/repository"
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FollowerUseCase interface {
	CreateFollower(follower *domain.ProfileFollower) (*domain.ProfileFollower, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) *mongo.DeleteResult
	GetAllUsersFollowers(user dto.ProfileDTO) ([]*domain.Profile, error)
	AlreadyFollowing (ctx context.Context, following *domain.ProfileFollowing) (bool, error)
	GetFollowersForFront(ctx context.Context, userId string) ([]dto.FollowerDTO, error)

}

type followerUseCase struct {
	FollowerRepo repository.FollowerRepo
	logger *logger.Logger
}

func (f followerUseCase) GetFollowersForFront(ctx context.Context, userId string) ([]dto.FollowerDTO, error) {
	f.logger.Logger.Infof("getting followers for front")

	following, _ := f.GetAllUsersFollowers(dto.ProfileDTO{ID: userId})
	var retVal []dto.FollowerDTO
	for _, follow := range following {
		profile, _ := gateway.GetUser(context.Background(), follow.ID)
		retVal = append(retVal, dto.FollowerDTO{Id: userId, ProfilePhoto: profile.ProfilePhoto, Username: profile.Username})
	}
		return retVal, nil
}


func (f followerUseCase) GetAllUsersFollowers(user dto.ProfileDTO) ([]*domain.Profile, error) {
	f.logger.Logger.Infof("getting all users followers")

	userFollowersBson, err :=  f.FollowerRepo.GetAllUsersFollowers(user)
	if err !=nil{
		f.logger.Logger.Errorf("failed cancel follow request, error: %v\n", err)
		return nil, err
	}

	var usersFollowers []*domain.Profile
	for _, uf := range userFollowersBson {
		bsonBytes, _ := bson.Marshal(uf)
		var follower *domain.ProfileFollower

		err := bson.Unmarshal(bsonBytes, &follower)
		if err != nil {
			f.logger.Logger.Errorf("failed cancel follow request, error: %v\n", err)
			return nil, err
		}
		usersFollowers = append(usersFollowers, &follower.Follower)
	}

	return usersFollowers,nil
}


func (f *followerUseCase) AlreadyFollowing(ctx context.Context, following *domain.ProfileFollowing) (bool, error) {
	f.logger.Logger.Infof("checking is already following")
	return f.FollowerRepo.AlreadyFollowing(ctx, following)
}

func (f followerUseCase) CreateFollower(follower *domain.ProfileFollower) (*domain.ProfileFollower, error) {
	f.logger.Logger.Infof("creating follower")
	return f.FollowerRepo.CreateFollower(follower)
}

func (f followerUseCase) GetByID(id string) *mongo.SingleResult {
	f.logger.Logger.Infof("getting by id %v\n", id)
	return f.FollowerRepo.GetByID(id)
}

func (f followerUseCase) Delete(id string) *mongo.DeleteResult {
	f.logger.Logger.Infof("deletting by id %v\n", id)
	return f.FollowerRepo.Delete(id)
}

func NewFollowerUseCase(repo repository.FollowerRepo, logger *logger.Logger) FollowerUseCase {
	return &followerUseCase{
		FollowerRepo: repo,
		logger: logger,
	}
}
