package model

type SocketEventStruct struct {
	EventName    string      `json:"eventName"`
	EventPayload interface{} `json:"eventPayload"`
}

type UserDetailsResponsePayloadStruct struct {
	Username string `json:"username"`
	UserID   int `json:"userID"`
	Status   string `json:"status"`
}

type MessageRequest struct {
	OriginUserID			int				`json:"origin_user_id"`	
	DestinationUserID		int				`json:"destination_user_id"`	
	Message			 		string			`json:"message"`
}