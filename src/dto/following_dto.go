package dto

type FollowingDTO struct {
	Id  string `json:"id"`
	Username string `json:"username"`
	ProfilePhoto  string `json:"profilePhoto"`
	CloseFriend  bool `json:"closeFriend"`
}
