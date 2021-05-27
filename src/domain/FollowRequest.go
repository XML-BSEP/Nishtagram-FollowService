package domain

import "time"

type FollowRequest struct {
	ID uint64 `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	UserRequested Profile `json:"user-requested"`
	FollowedAccount Profile `json:"followed-account"`
}