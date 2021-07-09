package dto

type UserDTO struct {
	Name     string `bson:"name" json:"name"`
	Surname  string `bson:"surname" json:"surname"`
	Email    string `bson:"email" json:"email"`
	Address  string `bson:"address" json:"address"`
	Phone    string `bson:"phone" json:"phone"`
	Birthday string `bson:"birthday" json:"birthday"`
	Gender   string `bson:"gender" json:"gender"`
	Web      string `bson:"web" json:"web"`
	Bio      string `bson:"bio" json:"bio"`
	Username string `bson:"username" json:"username"`
	Image    string `bson:"image" json:"image"`
	Private bool `bson:"private" json:"private"`
	Category string `bson:"category" json:"category"`
}

type UserIdsDto struct{
	Ids []string `json:"ids"`
}
