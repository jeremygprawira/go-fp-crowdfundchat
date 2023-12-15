package model

import "time"

type Chat struct {
	ID          	int  				`json:"id"`
	Content			string				`json:"chat"`
	UserID			int					`json:"user_id"`
	UserName		string				`json:"user_name"`
	ProjectID		int					`json:"project_id"`
	CreatedAt 		time.Time			`json:"created_at"`				
}

type CreateChatRoomRequest struct {
	ProjectID		int 				`json:"project_id"`
	Name			string 				`json:"name"`
}

type ChatRequest struct {
	Chat			string 				`json:"chat"`
	ProjectID		string				`json:"room_id"`
	UserName		string				`json:"user_name"`
	UserID			string				`json:"user_id"`
	CreatedAt 		time.Time			`json:"created_at"`			
}

type GetRoomResponse struct {
	ID				string				`json:"room_id"`
	RoomName		string				`json:"user_name"`
}

type GetClientResponse struct {
	ID				string				`json:"client_id"`
	UserName		string				`json:"user_name"`
}

type InputChatRequest struct {
	UserID			int					`json:"user_id"`	
	ProjectID		int					`json:"project_id"`	
	Message			string				`json:"message"`
}

type ChatHistoryRequest struct {	
	ID				int			`uri:"id" binding:"required"`
}

type ChatHistoryResponse struct {
	UserID			int					`json:"user_id"`
	UserName		string				`json:"user_name"`	
	ProjectID		int					`json:"project_id"`	
	Chat			string				`json:"chat"`
	CreatedAt 		time.Time			`json:"created_at"`
}