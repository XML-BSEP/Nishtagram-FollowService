package domain

import "time"

type ProfileFollowing struct {
	ID string `bson:"_id,omitempty" json:"id"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	User Profile `bson:"user" json:"user"`
	Following Profile `bson:"following" json:"following"`
}