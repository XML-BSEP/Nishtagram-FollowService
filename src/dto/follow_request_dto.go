package dto

import (
	"time"
)

type FollowRequestDTO struct {
	ID string `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	UserRequested ProfileDTO `json:"user-requested"`
	FollowedAccount ProfileDTO `json:"followed-account"`
}