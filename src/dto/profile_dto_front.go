package dto

type ProfileDTOFront struct {
	Id string `json:"id"`
	ProfilePhoto string `json:"profilePhoto"`
	Username string `json:"username"`
	IsCloseFriend string `json:"close_friend"`
}
