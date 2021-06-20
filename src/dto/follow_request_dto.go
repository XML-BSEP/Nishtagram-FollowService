package dto

type FollowRequestDTO struct {
	ID string `json:"id"`
	UserRequested string `bson:"user_requested" json:"userrequested"`
	FollowedAccount string `bson:"followed_account" json:"followedaccount"`
}
