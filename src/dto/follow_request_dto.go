package dto

type FollowRequestDTO struct {
	ID string `json:"id"`
	UserRequested string `bson:"user_requested" json:"user-requested"`
	FollowedAccount string `bson:"followed_account" json:"followed-account"`
}
