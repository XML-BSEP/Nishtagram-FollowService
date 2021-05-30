package domain

import "time"

type FollowRequest struct {
	ID string `bson:"_id,omitempty" json:"id"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	UserRequested Profile `bson:"user_requested" json:"user-requested"`
	FollowedAccount Profile `bson:"followed_account" json:"followed-account"`
}