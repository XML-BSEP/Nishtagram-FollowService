package dto

type Follow struct {
	UserToFollow  ProfileDTO `bson:"follower"`
	UserFollowing ProfileDTO `bson:"user"`
}