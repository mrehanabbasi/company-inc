package models

import "github.com/golang-jwt/jwt/v4"

type User struct {
	ID       string `json:"id,omitempty" bson:"_id"`
	Name     string `json:"name" bson:"user_name" binding:"required,min=4,max=30"`
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Password string `json:"password" bson:"password" binding:"required,max=20"`
}

type UserLogin struct {
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Password string `json:"password" bson:"password" binding:"required,max=20"`
}

type JwtClaims struct {
	jwt.RegisteredClaims
	TokenInfo
}

type TokenInfo struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"email"`
}
