package model

import "time"

type Chat struct {
	ID 						int
	OriginUserID			int				`json:"origin_user_id"`	
	DestinationUserID		int				`json:"destination_user_id"`	
	Message			 		string			`json:"message"`
	CreatedAt 				time.Time		`json:"created_at"`
	
	User					User			`json:"user" gorm:"foreignKey:ID"`
}

type InputChatRequest struct {
	OriginUserID			int				`json:"origin_user_id"`	
	DestinationUserID		int				`json:"destination_user_id"`	
	Message			 		string			`json:"message"`
}

type ChatHistoryRequest struct {
	OriginUserID			int				`uri:"origin_id" binding:"required"`
	DestinationUserID		int				`uri:"destination_id" binding:"required"`	
}

type ChatHistoryResponse struct {
	OriginUserID			int				`json:"origin_user_id"`	
	DestinationUserID		int				`json:"destination_user_id"`	
	Message			 		string			`json:"message"`
	CreatedAt				time.Time		`json:"created_at"`
}

/*type ContactListRequest struct {
	OriginUserID			int				`uri:"origin_id" binding:"required"`
}

type UserContactListResponse struct {
	Name					string			`json:"name"`
	UserID					int				`json:"id"`
	AvatarFileName			string			`json:"avatar_file_name"`
}*/

type CreateRoomRequest struct {
	OriginUserID			string			`json:"origin_id"`
	DestinationUserID		string			`json:"destination_id"`
}

type InitiateChatRequest struct {
	UserID					int				`uri:"user_id" binding:"required"`
}