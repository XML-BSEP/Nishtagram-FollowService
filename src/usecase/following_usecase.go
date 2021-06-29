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
	"time"
)

type FollowingUseCase interface {
	CreateFollowing(ctx context.Context, following *domain.ProfileFollowing) (*domain.ProfileFollowing, error)
	GetByID(id string) *mongo.SingleResult
	Delete(ctx context.Context,id string) *mongo.DeleteResult
	GetAllUsersFollowings(user dto.ProfileDTO) ([]*domain.Profile, error)
	GetAllUsersProfileFollowings(user dto.ProfileDTO) ([]*domain.ProfileFollowing, error)

	Unfollow(ctx context.Context, userToUnfollow string, userUnfollowing string) error
	AlreadyFollowing(ctx context.Context, following *domain.ProfileFollowing) bool
	GetUserFollowingsForFrontend(ctx context.Context, userId string) ([]dto.FollowingDTO, error)
	BanUser(ctx context.Context, userId string)error

}
func (f followingUseCase) BanUser(ctx context.Context, userId string) error{
	followers, err := f.FollowerUseCase.GetAllUsersFollowers(dto.ProfileDTO{ID: userId})
	if err!=nil{
		return err
	}

	followings, err1 := f.GetAllUsersProfileFollowings(dto.ProfileDTO{ID: userId})
	if err1!=nil{
		return err1
	}

	for _, it :=range followers{
		err2:=f.Unfollow(ctx, it.User.ID, it.Follower.ID)
		if err2!=nil{
			return err
		}
	}

	for _, it :=range followings{
		err3:=f.Unfollow(ctx, it.Following.ID, it.User.ID)
		if err3!=nil{
			return err
		}
	}

	return nil
}


type followingUseCase struct {
	FollowingRepo       repository.FollowingRepo
	FollowerRepo 		repository.FollowerRepo
	//ProfileRepo          repository.ProfileRepo
	FollowRequestUseCase FollowRequestUseCase
	FollowerUseCase FollowerUseCase
	logger *logger.Logger
}

func (f followingUseCase) GetAllUsersProfileFollowings(user dto.ProfileDTO) ([]*domain.ProfileFollowing, error) {
	f.logger.Logger.Infof("getting all users following")

	userFollowingBson, err := f.FollowingRepo.GetAllUsersFollowings(user)
	if err!=nil{
		f.logger.Logger.Errorf("failed getting all users following, error: %v\n", err)
		return nil, err
	}
	var usersFollowings []*domain.ProfileFollowing

	for _, uf := range userFollowingBson {
		bsonBytes, _ := bson.Marshal(uf)
		var following *domain.ProfileFollowing

		err := bson.Unmarshal(bsonBytes, &following)
		if err != nil {
			f.logger.Logger.Errorf("failed to unmarshal: %v\n", err)
			return nil, err
		}

		usersFollowings = append(usersFollowings, following)
	}
	return usersFollowings, nil}

func (f followingUseCase) Unfollow(ctx context.Context, userToUnfollow string, userUnfollowing string) error {
	f.logger.Logger.Infof("user unfollowing %v unfollowing user %v\n",userUnfollowing ,userToUnfollow)
	var err error
	if err = f.FollowingRepo.RemoveFollowing(ctx, userToUnfollow, userUnfollowing);err !=nil{
		f.logger.Logger.Errorf("failed removing following, error: %v\n", err)

	}
	if err = f.FollowerRepo.RemoveFollower(ctx, userToUnfollow, userUnfollowing); err!=nil{
		f.logger.Logger.Errorf("failed removing follower, error: %v\n", err)

	}
	if err!=nil{
		return err
	}
	return nil
}

func (f followingUseCase) GetUserFollowingsForFrontend(ctx context.Context, userId string) ([]dto.FollowingDTO, error) {
	f.logger.Logger.Infof("getting user followings for fronend")

	following, _ := f.GetAllUsersFollowings(dto.ProfileDTO{ID: userId})
	var retVal []dto.FollowingDTO
	for _, follow := range following {
		profile, _ := gateway.GetUser(context.Background(), follow.ID)
		retVal = append(retVal, dto.FollowingDTO{Id: follow.ID, ProfilePhoto: profile.ProfilePhoto, Username: profile.Username})
	}
	return retVal, nil
}

func (f followingUseCase) AlreadyFollowing(ctx context.Context, following *domain.ProfileFollowing) bool {
	f.logger.Logger.Infof("checking if already following")
	return f.FollowingRepo.AlreadyFollowing(ctx, following)

}


func (f followingUseCase) GetAllUsersFollowings(user dto.ProfileDTO) ([]*domain.Profile, error) {
	f.logger.Logger.Infof("getting all users following")

	userFollowingBson, err := f.FollowingRepo.GetAllUsersFollowings(user)
	if err!=nil{
		f.logger.Logger.Errorf("failed getting all users following, error: %v\n", err)
		return nil, err
	}
	var usersFollowings []*domain.Profile

	for _, uf := range userFollowingBson {
		bsonBytes, _ := bson.Marshal(uf)
		var following *domain.ProfileFollowing

		err := bson.Unmarshal(bsonBytes, &following)
		if err != nil {
			f.logger.Logger.Errorf("failed to unmarshal: %v\n", err)
			return nil, err
		}

		usersFollowings = append(usersFollowings, &following.Following)
	}
	return usersFollowings, nil
}

func (f *followingUseCase) CreateFollowing(ctx context.Context, following *domain.ProfileFollowing) (*domain.ProfileFollowing, error) {
	f.logger.Logger.Infof("creating following for user %v\n", following.ID)


	isPrivate, err := gateway.IsProfilePrivate(ctx, following.Following.ID)
	if err != nil {
		f.logger.Logger.Errorf("is profile private error, : %v\n", err)
		return nil, err
	}

	//TODO: A call towards profile microservice, because all data about profiles is stored in that db because of consistency
	if !isPrivate{
		newfollowing, err := f.FollowingRepo.CreateFollowing(following)
		if err != nil{
			f.logger.Logger.Errorf("create following error: %v\n", err)
			return nil, err
		}
		newfollower := domain.ProfileFollower{Follower: following.User, User:following.Following, Timestamp: time.Now()}
		_, err = f.FollowerUseCase.CreateFollower(&newfollower)
		if err!=nil{
			f.logger.Logger.Errorf("create follower error: %v\n", err)
			return nil, err
		}
		return newfollowing, nil
	} else {

		req := domain.FollowRequest{Timestamp: time.Now(), FollowedAccount: following.Following, UserRequested: following.User}

		_,err :=f.FollowRequestUseCase.CreateFollowRequest(&req)
		return following, err
	}

	return nil,nil
}

func (f followingUseCase) GetByID(id string) *mongo.SingleResult {
	f.logger.Logger.Infof("getting following by id %v\n", id)
	return f.FollowingRepo.GetByID(id)
}

func (f followingUseCase) Delete(ctx context.Context, id string) *mongo.DeleteResult {
	f.logger.Logger.Infof("deleting following by id %v\n", id)
	//if f.FollowerUseCase.Delete(id).DeletedCount==1{
		return f.FollowingRepo.Delete(ctx, id)
	//}else{
	//	return &mongo.DeleteResult{DeletedCount: 0}
	//}

}

func NewFollowingUseCase(followingRepo repository.FollowingRepo,
						//profileRepo repository.ProfileRepo,
						followReqUseCase FollowRequestUseCase,
						followerUseCase FollowerUseCase,
						followerRepo repository.FollowerRepo,
						logger *logger.Logger) FollowingUseCase {
	return &followingUseCase{
		FollowingRepo: followingRepo,
		FollowerRepo:  followerRepo,
		//ProfileRepo:           profileRepo,
		FollowRequestUseCase : followReqUseCase,
		FollowerUseCase: followerUseCase,
		logger: logger,
	}
}
