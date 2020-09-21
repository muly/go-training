package model

type Student struct {
	ID      string `bson:"_id" json:"id"`
	Name    string `bson:"name" json:"name"`
	Email   string `bson:"email" json:"email"`
	English int    `bson:"english" json:"english"`
	Maths   int    `bson:"maths" json:"maths"`
	Total   int    `bson:"total" json:"total"`
}
