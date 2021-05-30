package dto

import "time"

type ProfileFollowerDTO struct {
	ID string `json:"id"`
	CloseFriend bool `json:"close-friend"`
	Timestamp time.Time `json:"timestamp"`
	User ProfileDTO `json:"user"`
	Follower ProfileDTO `json:"follower"`
}