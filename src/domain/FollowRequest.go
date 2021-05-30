package domain

import "time"

type FollowRequest struct {
	ID string `bson:"_id,omitempty" json:"id"`
	Timestamp time.Time `json:"timestamp"`
	UserRequested Profile `json:"user-requested"`
	FollowedAccount Profile `json:"followed-account"`
}