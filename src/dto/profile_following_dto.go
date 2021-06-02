package dto

import (
	"FollowService/domain"
	"time"
)

type ProfileFollowingDTO struct {
	ID string `bson:"_id" json:"id"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	User ProfileDTO `bson:"user" json:"user"`
	Following ProfileDTO `bson:"following" json:"following"`
}

func NewProfileFollowingDTOToNewProfileFollowing(dto ProfileFollowingDTO) *domain.ProfileFollowing{
	return &domain.ProfileFollowing{Timestamp: dto.Timestamp, User: NewProfileDTOToNewProfile(dto.User), Following: NewProfileDTOToNewProfile(dto.Following)}
}
