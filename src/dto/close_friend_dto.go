package dto

type CloseFriendDTO struct {
	CloseFriend string `bson:"follower" json:"follower"`
	User        string `bson:"user" json:"user"`
}
