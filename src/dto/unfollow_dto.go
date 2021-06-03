package dto

type Unfollow struct {
	UserToUnfollow  string `bson:"following" json:"following"`
	UserUnfollowing string `bson:"user" json:"user"`
}