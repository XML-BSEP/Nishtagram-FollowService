package domain

import "time"

type ProfileFollowers struct {
	ID uint64 `json:"id"`
	CloseFriend bool `json:"close-friend"`
	Timestamp time.Time `json:"timestamp"`
	User Profile `json:"user"`
	Follower Profile `json:"follower"`
}