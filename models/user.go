package models

type User struct {
	Username string `json:"username" bson:"username"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Picture  string `json:"picture" bson:"picture"`
	Job      string `json:"job" bson:"job"`
}
