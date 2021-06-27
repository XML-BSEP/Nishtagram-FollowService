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
	GetAllUsersFollowers(user dto.ProfileDTO) ([]*domain.ProfileFollower, error)

	AlreadyFollowing (ctx context.Context, following *domain.ProfileFollowing) (bool, error)
	GetFollowersForFront(ctx context.Context, userId string) ([]dto.FollowerDTO, error)
	AddToCloseFriends(ctx context.Context, newCloseFriendId string, userId string) error
	RemoveFromCloseFriends(ctx context.Context, removedCloseFriendId string, userId string) error
	GetAllUsersCloseFriends(ctx context.Context, userId string)([]*domain.Profile, error)
	GetAllUsersToWhomUserIsCloseFriend(ctx context.Context, userId string)([]*domain.Profile, error)
	IsUserFollowingUser(ctx context.Context, userFollowing string, userFollowed string)(bool, error)
}

type followerUseCase struct {
	FollowerRepo repository.FollowerRepo
	logger *logger.Logger
}



func (f followerUseCase) IsUserFollowingUser(ctx context.Context, userFollowing string, userFollowed string) (bool, error) {
	f.logger.Logger.Infof("is user following user")
	profile_follower, err := f.FollowerRepo.GetFollowerByFollowerAndUser(ctx, userFollowing, userFollowed)

	if err !=nil{
		f.logger.Logger.Errorf("failed getting profile follower, error: %v\n", err)
		return false, err
	}

	if profile_follower !=nil{
		return true, nil
	}

	return false, nil

}

func (f followerUseCase) GetAllUsersCloseFriends(ctx context.Context, userId string) ([]*domain.Profile, error) {
	f.logger.Logger.Infof("get all users close friends")
	close_friends, err := f.FollowerRepo.GetAllUsersCloseFriends(ctx, userId)

	if err != nil {
		f.logger.Logger.Errorf("failed getting profile follower, error: %v\n", err)
		return nil, err
	}

	var followers []*domain.Profile
	for _, cf := range close_friends {
		bsonBytes, _ := bson.Marshal(cf)
		var follower *domain.ProfileFollower

		err := bson.Unmarshal(bsonBytes, &follower)
		if err != nil {
			f.logger.Logger.Errorf("failed cancel follow request, error: %v\n", err)
			return nil, err
		}
		followers = append(followers, &domain.Profile{ID: follower.Follower.ID})
	}
	return followers, nil
}
func (f followerUseCase) GetAllUsersToWhomUserIsCloseFriend(ctx context.Context, userId string) ([]*domain.Profile, error) {
	f.logger.Logger.Infof("get all users to whom user is in close friends")
	close_friends, err := f.FollowerRepo.GetAllUsersToWhomUserIsCloseFriend(ctx, userId)

	if err != nil {
		f.logger.Logger.Errorf("failed getting profile follower, error: %v\n", err)
		return nil, err
	}

	var followers []*domain.Profile
	for _, cf := range close_friends {
		bsonBytes, _ := bson.Marshal(cf)
		var follower *domain.ProfileFollower

		err := bson.Unmarshal(bsonBytes, &follower)
		if err != nil {
			f.logger.Logger.Errorf("failed cancel follow request, error: %v\n", err)
			return nil, err
		}
		followers = append(followers, &domain.Profile{ID: follower.User.ID})
	}
	return followers, nil
}

func (f followerUseCase) AddToCloseFriends(ctx context.Context, newCloseFriendId string, userId string) error {
	f.logger.Logger.Infof("adding to close friends")
	profile_follower, err := f.FollowerRepo.GetFollowerByFollowerAndUser(ctx, newCloseFriendId, userId)

	if err !=nil{
		f.logger.Logger.Errorf("failed getting profile follower, error: %v\n", err)
		return err
	}

	errRepo := f.FollowerRepo.UpdateCloseFriendStatus(ctx, profile_follower.ID, true)
	if errRepo != nil {
		f.logger.Logger.Errorf("error while editting user, error %v\n", errRepo)
		return errRepo
	}
	return nil

}

func (f followerUseCase) RemoveFromCloseFriends(ctx context.Context, removedCloseFriendId string, userId string) error {
	f.logger.Logger.Infof("adding to close friends")
	profile_follower, err := f.FollowerRepo.GetFollowerByFollowerAndUser(ctx, removedCloseFriendId, userId)

	if err !=nil{
		f.logger.Logger.Errorf("failed getting profile follower, error: %v\n", err)
		return err
	}

	errRepo := f.FollowerRepo.UpdateCloseFriendStatus(ctx, profile_follower.ID, false)
	if errRepo != nil {
		f.logger.Logger.Errorf("error while editting user, error %v\n", errRepo)
		return errRepo
	}
	return nil

}

func (f followerUseCase) GetFollowersForFront(ctx context.Context, userId string) ([]dto.FollowerDTO, error) {
	f.logger.Logger.Infof("getting followers for front")

	followers, _ := f.GetAllUsersFollowers(dto.ProfileDTO{ID: userId})
	var retVal []dto.FollowerDTO

	for _, follow := range followers {
		profile, _ := gateway.GetUser(context.Background(), follow.Follower.ID)
		retVal = append(retVal, dto.FollowerDTO{Id: follow.Follower.ID, ProfilePhoto: profile.ProfilePhoto, Username: profile.Username, CloseFriend: follow.CloseFriend})
	}


		return retVal, nil
}


func (f followerUseCase) GetAllUsersFollowers(user dto.ProfileDTO) ([]*domain.ProfileFollower, error) {
	f.logger.Logger.Infof("getting all users followers")

	userFollowersBson, err :=  f.FollowerRepo.GetAllUsersFollowers(user)
	if err !=nil{
		f.logger.Logger.Errorf("failed cancel follow request, error: %v\n", err)
		return nil, err
	}
	var followers []*domain.ProfileFollower
	for _, uf := range userFollowersBson {
		bsonBytes, _ := bson.Marshal(uf)
		var follower *domain.ProfileFollower

		err := bson.Unmarshal(bsonBytes, &follower)
		if err != nil {
			f.logger.Logger.Errorf("failed cancel follow request, error: %v\n", err)
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers,nil
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
