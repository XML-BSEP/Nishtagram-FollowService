package dto

type FollowDTO struct {
	Follower ProfileDTO `bson:"follower"`
	User     ProfileDTO `bson:"user"`
}

