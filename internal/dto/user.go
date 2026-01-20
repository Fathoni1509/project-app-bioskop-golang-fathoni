package dto

import "time"

type UserRegister struct {
	Name      string `json:"name" validate:"required,min=0"`
	Email     string `json:"email" validate:"required,min=0,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Token     string `json:"token"`
	TokenTime time.Time
}

type UserLogin struct {
	Name      string `json:"name" validate:"required,min=0"`
	Password  string `json:"password" validate:"required,min=6"`
	Token     string `json:"token"`
	TokenTime time.Time
}

type UserResponse struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}
