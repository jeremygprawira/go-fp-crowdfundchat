package websockets

import (
	"go-fp-crowdfundchat/model"

	"github.com/gorilla/websocket"
)

/*type Client struct {
	Conn		*websocket.Conn
	Msg			chan *Message
	ID			string				`json:"id"`
	RoomID		string				`json:"room_id"`
	Username	string				`json:"username"`
}*/

type Client struct {
	Hub			*Hub
	Conn		*websocket.Conn
	Send		chan *model.SocketEventStruct
	UserID		int
}

type Message struct {
	Content		string				`json:"content"`
	RoomID		string				`json:"room_id"`
	Username	string				`json:"username"`
}