package usecase

import (
	"FollowService/domain"
	"FollowService/dto"
	"FollowService/repository"
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
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
	logger *logger.Logger

}
func (f *followRequestUseCase) CancelFollowRequest(ctx context.Context, request *dto.FollowRequestDTO) error {
	f.logger.Logger.Infof("cancel follow request")

	followReq, err := f.FollowRequestRepo.GetFollowRequest(ctx,request)
	if err!=nil{
		f.logger.Logger.Errorf("failed cancel follow request, error: %v\n", err)
		return err
	}
	if asd,_ := f.Delete(followReq.ID); asd.DeletedCount==1{
		return  nil
	}else{
		return err
	}
}

func (f *followRequestUseCase) IsCreated(ctx context.Context, request *domain.FollowRequest) bool {
	f.logger.Logger.Infof("checking is created")
	return f.FollowRequestRepo.IsCreated(ctx, request)
}

func (f followRequestUseCase) ApprofeFollowRequest(ctx context.Context, req dto.FollowRequestDTO) error {
	f.logger.Logger.Infof("approving follow request")
	request, err := f.FollowRequestRepo.GetFollowRequestByUserAndFollower(ctx, req)
	if err!=nil{
		f.logger.Logger.Errorf("failed to get following request by user and follower, error: %v\n", err)
		return err
	}

	var requestDTO domain.FollowRequest

	bsonBytes, _ := bson.Marshal(request)
	err = bson.Unmarshal(bsonBytes, &requestDTO)
	if err != nil {
		f.logger.Logger.Errorf("failed to unmarshal, error: %v\n", err)
		return err
	}
	_, del_err := f.Delete(requestDTO.ID)
	if del_err!=nil{
		f.logger.Logger.Errorf("failed to delete, error: %v\n", del_err)
		return del_err
	}

	follower := dto.ProfileFollowerDTO{Follower: dto.ProfileDTO{ID: requestDTO.UserRequested.ID}, User: dto.ProfileDTO{ID: requestDTO.FollowedAccount.ID}, Timestamp: time.Now()}
	following := dto.ProfileFollowingDTO{Following: dto.ProfileDTO{ID: requestDTO.FollowedAccount.ID},User: dto.ProfileDTO{ID: requestDTO.UserRequested.ID}, Timestamp: time.Now()}

	_, err = f.FollowingRepository.CreateFollowing(dto.NewProfileFollowingDTOToNewProfileFollowing(following))
	if err!=nil{
		f.logger.Logger.Errorf("faild to create following, error: %v\n", err)
		return err
	}
	_, err = f.FollowerRepository.CreateFollower(dto.NewProfileFollowerDTOToNewProfileFollower(follower))

	if err!=nil{
		f.logger.Logger.Errorf("failed to create following, error: %v\n", err)
		return err
	}

	return nil
}

func (f followRequestUseCase) GetAllUsersFollowRequests(user dto.ProfileDTO) ([]*domain.Profile, error) {
	f.logger.Logger.Infof("get all users follow request for user %v\n", user.ID)

	userFollowRequestsBson, err := f.FollowRequestRepo.GetAllUsersFollowRequests(user)
	if err !=nil{
		f.logger.Logger.Errorf("failed to create following, error: %v\n", err)
		return nil,err
	}
	var usersFollowRequests []*domain.Profile
	var followRequest *domain.FollowRequest
	for _, uf := range userFollowRequestsBson {
		bsonBytes, _ := bson.Marshal(uf)
		err := bson.Unmarshal(bsonBytes, &followRequest)
		if err != nil {
			f.logger.Logger.Errorf("failed unmarshal, error: %v\n", err)
			return nil, err
		}
		usersFollowRequests = append(usersFollowRequests, &followRequest.UserRequested)
	}
	return usersFollowRequests, nil
}

func (f followRequestUseCase) CreateFollowRequest(req *domain.FollowRequest) (*domain.FollowRequest, error) {
	f.logger.Logger.Infof("creating follow request")
	return f.FollowRequestRepo.CreateFollowRequest(req)
}

func (f followRequestUseCase) GetByID(id string) *mongo.SingleResult {
	f.logger.Logger.Infof("getting by id %v\n", id)
	return f.FollowRequestRepo.GetByID(id)
}

func (f followRequestUseCase) Delete(id string) (*mongo.DeleteResult, error) {
	f.logger.Logger.Infof("delete by id %v\n", id)
	return f.FollowRequestRepo.Delete(id)
}

func NewFollowRequestUseCase(repo repository.FollowRequestRepo, followerRepo repository.FollowerRepo, followingRepo repository.FollowingRepo, logger *logger.Logger) FollowRequestUseCase {
	return &followRequestUseCase{
		FollowRequestRepo:   repo,
		FollowerRepository:  followerRepo,
		FollowingRepository: followingRepo,
		logger: logger,
	}
}