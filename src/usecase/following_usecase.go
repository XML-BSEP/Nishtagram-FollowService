package usecase

import (
	"FollowService/domain"
	"FollowService/dto"
	"FollowService/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FollowingUseCase interface {
	CreateFollowing(following *domain.ProfileFollowing) (*domain.ProfileFollowing, error)
	GetByID(id string) *mongo.SingleResult
	Delete(ctx context.Context,id string) *mongo.DeleteResult
	GetAllUsersFollowings(user dto.ProfileDTO) ([]*domain.Profile, error)
	Unfollow(ctx context.Context, unfollow dto.Unfollow) error
}

type followingUseCase struct {
	FollowingRepo       repository.FollowingRepo
	FollowerRepo 		repository.FollowerRepo
	//ProfileRepo          repository.ProfileRepo
	//FollowRequestUseCase FollowRequestUseCase
	//FollowerUseCase FollowerUseCase
}

func (f followingUseCase) Unfollow(ctx  context.Context, unfollow dto.Unfollow) error {
	if err := f.FollowingRepo.RemoveFollowing(ctx, unfollow);err !=nil{
		return err
	}
	if err := f.FollowerRepo.RemoveFollower(ctx,unfollow); err!=nil{
		return err
	}
	return nil
}

func (f followingUseCase) GetAllUsersFollowings(user dto.ProfileDTO) ([]*domain.Profile, error) {
	userFollowingBson, err := f.FollowingRepo.GetAllUsersFollowings(user)
	if err!=nil{
		return nil, err
	}
	var usersFollowings []*domain.Profile
	for _, uf := range userFollowingBson {
		bsonBytes, _ := bson.Marshal(uf)
		var following *domain.ProfileFollowing

		err := bson.Unmarshal(bsonBytes, &following)
		if err != nil {
			return nil, err
		}
		usersFollowings = append(usersFollowings, &following.Following)
	}
	return usersFollowings, nil
}

func (f followingUseCase) CreateFollowing(following *domain.ProfileFollowing) (*domain.ProfileFollowing, error) {
	//var followedUser domain.Profile
	//result := f.ProfileRepo.GetByID(following.Following.ID)
	//decodeError := result.Decode(&followedUser)
	//if decodeError !=nil{
	//	return nil, decodeError
	//}
	//TODO: A call towards profile microservice, because all data about profiles is stored in that db because of consistency
	//if !followedUser.IsPrivate{
	//	newfollowing, err := f.FollowingRepo.CreateFollowing(following)
	//	if err != nil{
	//		return nil, err
	//	}
	//	newfollower := domain.ProfileFollower{Follower: following.User, User:following.Following, Timestamp: time.Now()}
	//	_, err = f.FollowerUseCase.CreateFollower(&newfollower)
	//	if err!=nil{
	//		return nil, err
	//	}
	//	return newfollowing, nil
	//}else {
	//	req := domain.FollowRequest{Timestamp: time.Now(), FollowedAccount: following.Following, UserRequested: following.User}
	//	_,err :=f.FollowRequestUseCase.CreateFollowRequest(&req)
	//	return following, err
	//}

	return nil,nil
}

func (f followingUseCase) GetByID(id string) *mongo.SingleResult {
	return f.FollowingRepo.GetByID(id)
}

func (f followingUseCase) Delete(ctx context.Context, id string) *mongo.DeleteResult {
	//if f.FollowerUseCase.Delete(id).DeletedCount==1{
		return f.FollowingRepo.Delete(ctx, id)
	//}else{
	//	return &mongo.DeleteResult{DeletedCount: 0}
	//}

}

func NewFollowingUseCase(followingRepo repository.FollowingRepo,
						//profileRepo repository.ProfileRepo,
						//followReqUseCase FollowRequestUseCase,
						//followerUseCase FollowerUseCase,
						followerRepo repository.FollowerRepo,) FollowingUseCase {
	return &followingUseCase{
		FollowingRepo: followingRepo,
		FollowerRepo:  followerRepo,
		//ProfileRepo:           profileRepo,
		//FollowRequestUseCase : followReqUseCase,
		//FollowerUseCase: followerUseCase,
	}
}
