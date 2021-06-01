package dto

import "time"

type ProfileFollowingDTO struct {
	ID string `bson:"_id" json:"id"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	User ProfileDTO `bson:"user" json:"user"`
	Following ProfileDTO `bson:"following" json:"following"`
}