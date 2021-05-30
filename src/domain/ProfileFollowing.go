package domain

import "time"

type ProfileFollowing struct {
	ID string `bson:"_id,omitempty" json:"id"`
	CloseFriend bool `json:"close-friend"`
	Timestamp time.Time `json:"timestamp"`
	User Profile `json:"user"`
	Following Profile `json:"following"`
}