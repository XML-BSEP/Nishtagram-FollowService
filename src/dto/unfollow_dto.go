package dto

type Unfollow struct {
	UserToUnfollow  ProfileDTO `bson:"following"`
	UserUnfollowing ProfileDTO `bson:"user"`
}