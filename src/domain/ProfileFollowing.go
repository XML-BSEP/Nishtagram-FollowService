package domain

import "time"

type ProfileFollowing struct {
	ID uint64 `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	User Profile `json:"user"`
	Following Profile `json:"following"`
}