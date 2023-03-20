package model

import (
	"time"
)

type User struct {
	ID 					int
	Name 				string 			`json:"name"`
	PhoneNo 			string			`json:"phone_no"`
	PIN 				string			`json:"pin"`
	AvatarFileName 		string			`json:"avatar_file_name"`
	Role 				string			`json:"role"`
	CreatedAt 			time.Time		`json:"created_at"`
	UpdatedAt 			time.Time		`json:"updated_at"`
}

type RegisterUserRequest struct {
	Name		string		`json:"name" binding:"required"`
	PhoneNo 	string		`json:"phoneNo" binding:"required"`
	PIN			string		`json:"pin" binding:"required"`
}

type LoginUserRequest struct {
	PhoneNo		string		`json:"phoneNo" binding:"required"`
	PIN			string		`json:"pin" binding:"required"`
}

type PhoneNoBodyRequest struct {
	PhoneNo		string		`json:"phoneNo" binding:"required"`
}

type PinValidationRequest struct {
	PIN			string		`json:"pin" binding:"required"`
}