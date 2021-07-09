package dto


type SearchUserDTO struct {
	Id		 string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Surname  string `bson:"surname" json:"surname"`
	Username string `bson:"username" json:"username"`
	Private bool `bson:"private" json:"private"`
	Image	string `json:"image" bson:"image"`
}

