package dto

type FollowReqDTO struct {
	Id  			string 	`bson:"_id" json:"id"`
	UserFollowedId	string	`json:"userFollowedId"`
	UserId			string 	`json:"userId"`
	Username 		string 	`json:"username"`
	ProfilePhoto  	string 	`json:"profilePhoto"`

}