package usecase

import (
	"FollowService/domain"
	"FollowService/dto"
	"FollowService/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type FollowRequestUseCase interface {
	CreateFollowRequest(req *domain.FollowRequest) (*domain.FollowRequest, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) (*mongo.DeleteResult, error)
	GetAllUsersFollowRequests(user dto.ProfileDTO) ([]*domain.Profile, error)
	ApprofeFollowRequest(ctx context.Context,req dto.FollowRequestDTO) error
	IsCreated(ctx context.Context, request *domain.FollowRequest) bool
	CancelFollowRequest(ctx context.Context, request *dto.FollowRequestDTO) error
}

type followRequestUseCase struct {
	FollowRequestRepo repository.FollowRequestRepo
	FollowerRepository repository.FollowerRepo
	FollowingRepository repository.FollowingRepo

}
func (f *followRequestUseCase) CancelFollowRequest(ctx context.Context, request *dto.FollowRequestDTO) error {
	followReq, err := f.FollowRequestRepo.GetFollowRequest(ctx,request)
	if err!=nil{
		return err
	}
	if asd,_ := f.Delete(followReq.ID); asd.DeletedCount==1{
		return  nil
	}else{
		return err
	}
}

func (f *followRequestUseCase) IsCreated(ctx context.Context, request *domain.FollowRequest) bool {
	return f.FollowRequestRepo.IsCreated(ctx, request)
}

func (f followRequestUseCase) ApprofeFollowRequest(ctx context.Context, req dto.FollowRequestDTO) error {
	request, err := f.FollowRequestRepo.GetFollowRequestByUserAndFollower(ctx, req)
	if err!=nil{
		return err
	}

	var requestDTO domain.FollowRequest

	bsonBytes, _ := bson.Marshal(request)
	err = bson.Unmarshal(bsonBytes, &requestDTO)
	if err != nil {
		return err
	}
	_, del_err := f.Delete(requestDTO.ID)
	if del_err!=nil{
		return del_err
	}

	follower := dto.ProfileFollowerDTO{Follower: dto.ProfileDTO{ID: requestDTO.UserRequested.ID}, User: dto.ProfileDTO{ID: requestDTO.FollowedAccount.ID}, Timestamp: time.Now()}
	following := dto.ProfileFollowingDTO{Following: dto.ProfileDTO{ID: requestDTO.FollowedAccount.ID},User: dto.ProfileDTO{ID: requestDTO.UserRequested.ID}, Timestamp: time.Now()}

	_, err = f.FollowingRepository.CreateFollowing(dto.NewProfileFollowingDTOToNewProfileFollowing(following))
	if err!=nil{
		return err
	}
	_, err = f.FollowerRepository.CreateFollower(dto.NewProfileFollowerDTOToNewProfileFollower(follower))

	if err!=nil{
		return err
	}

	return nil
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

func (f followRequestUseCase) Delete(id string) (*mongo.DeleteResult, error) {
	return f.FollowRequestRepo.Delete(id)
}

func NewFollowRequestUseCase(repo repository.FollowRequestRepo, followerRepo repository.FollowerRepo, followingRepo repository.FollowingRepo) FollowRequestUseCase {
	return &followRequestUseCase{
		FollowRequestRepo:   repo,
		FollowerRepository:  followerRepo,
		FollowingRepository: followingRepo,
	}
}