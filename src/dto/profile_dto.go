package dto

import "FollowService/domain"

type ProfileDTO struct {
	ID string `bson:"_id,omitempty" json:"id"`
}

func NewProfileDTOToNewProfile(dto ProfileDTO) domain.Profile{
	return domain.Profile{ID:dto.ID}
}