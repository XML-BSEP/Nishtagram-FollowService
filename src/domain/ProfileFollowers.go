package domain

import "time"

type ProfileFollowers struct {
	ID string `bson:"_id,omitempty" json:"id"`
	CloseFriend bool `bson:"close_friend" json:"close-friend"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	User Profile `bson:"user" json:"user"`
	Follower Profile `bson:"follower" json:"follower"`
}