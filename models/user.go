package models

type User struct {
	ID       string `json:"id,omitempty" bson:"_id"`
	Name     string `json:"name" bson:"user_name" binding:"required,min=4,max=30"`
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Password string `json:"password" bson:"password" binding:"required"`
}
