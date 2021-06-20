package implementation

import (
	"FollowService/dto"
	pb "FollowService/infrastructure/grpc/follow_service"
	"FollowService/usecase"
)

type FollowServiceImpl struct {
	pb.UnimplementedFollowServiceServer
	followingUsecase usecase.FollowingUseCase
}

func NewFollowServiceImpl(followingUsecase usecase.FollowingUseCase) *FollowServiceImpl {
	return &FollowServiceImpl{followingUsecase: followingUsecase}
}

func (f *FollowServiceImpl) SendUsers(in *pb.User, stream pb.FollowService_SendUsersServer) error {

	userId := in.UserId
	userFollowers, _ := f.followingUsecase.GetAllUsersFollowings(dto.ProfileDTO{ID: userId})

	for i := 0; i < len(userFollowers); i++ {
		followerMessage := &pb.Follower{FollowerUsername: userFollowers[i].ID}
		if err := stream.Send(followerMessage); err != nil {
			return err
		}

	}

	return nil
}
