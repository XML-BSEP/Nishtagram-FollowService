package dto

import (
	"FollowService/domain"
	"time"
)

type ProfileFollowerDTO struct {
	ID string `json:"id"`
	CloseFriend bool `json:"close-friend"`
	Timestamp time.Time `json:"timestamp"`
	User ProfileDTO `json:"user"`
	Follower ProfileDTO `json:"follower"`
}


func NewProfileFollowerDTOToNewProfileFollower(dto ProfileFollowerDTO) *domain.ProfileFollower{
	return &domain.ProfileFollower{CloseFriend: dto.CloseFriend, Timestamp: dto.Timestamp, User: NewProfileDTOToNewProfile(dto.User), Follower: NewProfileDTOToNewProfile(dto.Follower)}
}