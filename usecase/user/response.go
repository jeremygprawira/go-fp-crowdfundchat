package user

type UserBodyResponse struct {
	ID 					int
	Name 				string 			`json:"name"`
	PhoneNo 			string			`json:"phone_no"`
	Token 				string			`json:"token"`
}

func UserResponse (user User, token string) UserBodyResponse {
	response := UserBodyResponse {
		ID: user.ID,
		Name: user.Name,
		PhoneNo: user.PhoneNo,
		Token: token,
	}

	return response
}