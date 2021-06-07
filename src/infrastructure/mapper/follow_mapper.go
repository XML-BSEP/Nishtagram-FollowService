package mapper

import (
	"FollowService/domain"
	"FollowService/dto"
	"time"
)
func FollowDtoToFollowRequest(dto dto.FollowDTO) *domain.FollowRequest {
	return &domain.FollowRequest{
		UserRequested: domain.Profile{ID: dto.User.ID},
		FollowedAccount: domain.Profile{ID: dto.Follower.ID},
		Timestamp: time.Now(),
	}
}
func FollowDtoToProfileFollowing(dto dto.FollowDTO) *domain.ProfileFollowing {
	return &domain.ProfileFollowing{
		User: domain.Profile{ID: dto.User.ID},
		Following: domain.Profile{ID: dto.Follower.ID},
		Timestamp: time.Now(),
	}
}
