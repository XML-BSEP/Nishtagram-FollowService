package domain

import "time"

type ProfileFollowers struct {
	ID string `bson:"_id,omitempty" json:"id"`
	Timestamp time.Time `json:"timestamp"`
	User Profile `json:"user"`
	Follower Profile `json:"follower"`
}