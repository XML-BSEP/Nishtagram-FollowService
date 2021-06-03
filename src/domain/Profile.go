package domain

type Profile struct {
	ID string `bson:"_id,omitempty" json:"id"`
	//IsPrivate bool `bson:"private" json:"private"`
}