package user

import "time"

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