package dto

type FollowerDTO struct {
	Id  			string 	`bson:"_id" json:"id"`
	Username 		string 	`json:"username"`
	ProfilePhoto  	string 	`json:"profilePhoto"`
	CloseFriend 	bool 	`bson:"close_friend" json:"close_friend"`
}

