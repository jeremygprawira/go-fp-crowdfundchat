package user

type RegisterUserRequest struct {
	Name		string		`json:"name" binding:"required"`
	PhoneNo 	string		`json:"phoneNo" binding:"required"`
	PIN			string		`json:"pin" binding:"required"`
}