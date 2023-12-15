package http

import (
	"go-fp-crowdfundchat/model"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Connection		*websocket.Conn	
	Chat			chan *model.ChatRequest
	ID				string						`json:"id"`
	ProjectID		string						`json:"project_id"`
	UserName		string						`json:"user_name"`
}

func (c *Client) writeMessage() {
	defer func () {
		c.Connection.Close()
		
	}()

	for {
		chat, ok := <- c.Chat
		if !ok {
			return
		}

		c.Connection.WriteJSON(chat)
	}
}

func (c *Client) ReadMessage(h *Hub) {
	defer func() {
		h.Unregister <- c
		c.Connection.Close()
	}()

	for {
		_, m, err := c.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %s", err)
			}
			return
		}

		c := &model.ChatRequest{
			Chat: string(m),
			ProjectID: c.ProjectID,
			UserName: c.UserName,
			UserID: c.ID,
			CreatedAt: time.Now(),
		}

		h.Broadcast <- c
	}


}