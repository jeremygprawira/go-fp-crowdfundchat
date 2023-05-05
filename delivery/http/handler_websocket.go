package http

import (
	"go-fp-crowdfundchat/delivery/websockets"
	"go-fp-crowdfundchat/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	writeWait		= 10 * time.Second
	pongWait		= 60 * time.Second
	pingPeriod		= (pongWait * 9) / 10
	maxMessageSize	= 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitiateChat (c *gin.Context) {
	var request *model.InitiateChatRequest
	
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40010",
			"responseMessage": err.Error(),
		})
		return
	}

	websocketConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40010",
			"responseMessage": err.Error(),
		})
		return
	}

	hub := websockets.NewHub()
	go hub.Run()

	websockets.CreateSocketUser(hub, websocketConnection, request.UserID)
} 