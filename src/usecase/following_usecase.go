package usecase

import (
	"FollowService/domain"
	"FollowService/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type FollowingUseCase interface {
	CreateFollowing(following *domain.ProfileFollowing) (*domain.ProfileFollowing, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) *mongo.DeleteResult
}

type followingUseCase struct {
	FollowingRepo        repository.FollowingRepo
	ProfileRepo          repository.ProfileRepo
	FollowRequestUseCase FollowRequestUseCase
	FollowerUseCase FollowerUseCase
}

func (f followingUseCase) CreateFollowing(following *domain.ProfileFollowing) (*domain.ProfileFollowing, error) {
	var followedUser domain.Profile
	result := f.ProfileRepo.GetByID(following.Following.ID)
	decodeError := result.Decode(&followedUser)
	if decodeError !=nil{
		return nil, decodeError
	}
	if !followedUser.IsPrivate{
		newfollowing, err := f.FollowingRepo.CreateFollowing(following)
		if err != nil{
			return nil, err
		}
		newfollower := domain.ProfileFollower{Follower: following.User, User:following.Following, Timestamp: time.Now()}
		_, err = f.FollowerUseCase.CreateFollower(&newfollower)
		if err!=nil{
			return nil, err
		}
		return newfollowing, nil
	}else {
		req := domain.FollowRequest{Timestamp: time.Now(), FollowedAccount: following.Following, UserRequested: following.User}
		_,err :=f.FollowRequestUseCase.CreateFollowRequest(&req)
		return following, err
	}

}

func (f followingUseCase) GetByID(id string) *mongo.SingleResult {
	return f.FollowingRepo.GetByID(id)
}

func (f followingUseCase) Delete(id string) *mongo.DeleteResult {
	if f.FollowerUseCase.Delete(id).DeletedCount==1{
		return f.FollowingRepo.Delete(id)
	}else{
		return &mongo.DeleteResult{DeletedCount: 0}
	}
}

func NewFollowingUseCase(followRepo repository.FollowingRepo,
						profileRepo repository.ProfileRepo,
						followReqUseCase FollowRequestUseCase,
						followerUseCase FollowerUseCase) FollowingUseCase {
	return &followingUseCase{
		FollowingRepo:         followRepo,
		ProfileRepo:           profileRepo,
		FollowRequestUseCase : followReqUseCase,
		FollowerUseCase: followerUseCase,
	}
}
