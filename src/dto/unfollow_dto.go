package dto

type Unfollow struct {
	UserToUnfollow  string	`json:"following"`
	UserUnfollowing string	`json:"user"`
}