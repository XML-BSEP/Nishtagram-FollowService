package dto

import (
	"FollowService/domain"
	"time"
)

type ProfileFollowerDTO struct {
	ID string `bson:"_id" json:"id"`
	CloseFriend bool `bson:"close_friend" json:"close-friend"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	User ProfileDTO `bson:"user" json:"user"`
	Follower ProfileDTO `bson:"follower" json:"follower"`
}


func NewProfileFollowerDTOToNewProfileFollower(dto ProfileFollowerDTO) *domain.ProfileFollower{
	return &domain.ProfileFollower{CloseFriend: dto.CloseFriend, Timestamp: dto.Timestamp, User: NewProfileDTOToNewProfile(dto.User), Follower: NewProfileDTOToNewProfile(dto.Follower)}
}