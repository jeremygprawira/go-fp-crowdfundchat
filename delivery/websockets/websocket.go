package websockets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	hub *Hub
	userUsecase usecase.UserUsecase
	chatUsecase usecase.ChatUsecase
}

func NewWebsocketHandler (h *Hub, userUsecase usecase.UserUsecase, chatUsecase usecase.ChatUsecase) *WebsocketHandler {
	return &WebsocketHandler{
		hub: h,
		userUsecase: userUsecase,
		chatUsecase: chatUsecase,
	}
}

func (h *WebsocketHandler) CreateRoom(c *gin.Context) {
	var request model.CreateRoomRequest
	
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42210",
			"responseMessage": err.Error(),
		})
		return
	}

	/*b, _ := strconv.Atoi(request.OriginUserID)

	a, err := h.userUsecase.GetUserByID(b)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42210",
			"responseMessage": err.Error(),
		})
		return
	}*/

	s := fmt.Sprintf("%s%s", request.OriginUserID, request.DestinationUserID)

	h.hub.Rooms[s] = &Room{
		ID: s,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusOK, h.hub.Rooms)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func readConfig(c *Client) {
	c.Conn.SetReadLimit(512)
	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second)); return nil})
}

func (h *WebsocketHandler) SocketEvents(c *Client, s *model.SocketEventStruct) {
	type chatlistResponseStruct struct {
		Type     string      `json:"type"`
		Chatlist interface{} `json:"chatlist"`
	}
	
	switch s.EventName {
	case "join":
		userID := (s.EventPayload).(int)
		user, err := h.userUsecase.GetUserByID(userID)
		if err != nil {
			panic(err)
		}

		if user == (&model.User{}) {
			panic(err)
		}

		if user.Status == "offline" {
			panic(err)
		}

		newUserOnline := model.SocketEventStruct{
			EventName: "contact-list-response",
			EventPayload: chatlistResponseStruct{
				Type: "new-user-joined",
				Chatlist: model.UserDetailsResponsePayloadStruct{
					UserID: user.ID,
					Username: user.Name,
					Status: user.Status,
				},
			},
		}

		BroadcastEventOutsideSelfUser(c.Hub, &newUserOnline, user.ID)

		onlineUser, err := h.userUsecase.GetOnlineUserByStatus("online")
		if err != nil {
			panic(err)
		}

		allOnlineUser := model.SocketEventStruct{
			EventName: "contact-list-response",
			EventPayload: chatlistResponseStruct{
				Type: "self-contact-list",
				Chatlist: onlineUser,
			},
		}

		EmitToSpecificClient(c.Hub, &allOnlineUser, user.ID)
	case "disconnect":
		if s.EventPayload != nil {
			userID := (s.EventPayload).(int)
			user, err := h.userUsecase.GetUserByID(userID)
			if err != nil {
				panic(err)
			}
			h.userUsecase.PostUpdateUserStatus(user.ID, "offline")

			BroadcastEvent(c.Hub, &model.SocketEventStruct{
				EventName: "contact-list-response",
				EventPayload: chatlistResponseStruct{
					Type: "user-disconnected",
					Chatlist: model.UserDetailsResponsePayloadStruct{
						UserID: user.ID,
						Username: user.Name,
						Status: "offline",
					},
				},
			})
		}
	case "message":
		message := (s.EventPayload.(map[string]interface{})["message"]).(string)
		originUserID := (s.EventPayload.(map[string]interface{})["originUserID"]).(string)
		destinationUserID := (s.EventPayload.(map[string]interface{})["destinationUserID"]).(string)
	
		if message != "" && originUserID != "" && destinationUserID != "" {
			
			intOriginUserID, _ := strconv.Atoi(originUserID)
			intDestinationUserID, _ := strconv.Atoi(destinationUserID)

			messagePacket := model.InputChatRequest{
				Message: message,
				OriginUserID: intOriginUserID,
				DestinationUserID: intDestinationUserID,
			}
			
			h.chatUsecase.PostInputChat(&messagePacket)
			allOnlineUser := model.SocketEventStruct{
				EventName: "message-response",
				EventPayload: messagePacket,
			}

			EmitToSpecificClient(c.Hub, &allOnlineUser, intDestinationUserID)
		}
	}
}

func CreateSocketUser(h *Hub, cn *websocket.Conn, userID int) {
	client := &Client{
		Hub: h,
		Conn: cn,
		Send: make(chan *model.SocketEventStruct),
		UserID: userID,
	}

	go client.readPump()

	client.Hub.Register <- client
}

func UserRegisterEvent(h *Hub, client *Client) {
	h.Clients[client] = true

	
}

func EmitToSpecificClient(h *Hub, payload *model.SocketEventStruct, userID int) {
	for client := range h.Clients {
		if client.UserID == userID {
			select {
			case client.Send <- payload:
			default:
				close(client.Send)
				delete(h.Clients, client)
			}
		}
	}
}

func BroadcastEvent(h *Hub, s *model.SocketEventStruct) {
	for client := range h.Clients{
		select {
		case client.Send <- s:
		default:
			close(client.Send)
			delete(h.Clients, client)
		}	
	}
}

func BroadcastEventOutsideSelfUser(h *Hub, s *model.SocketEventStruct, userID int) {
	for client := range h.Clients{
		if client.UserID != userID {
			select {
			case client.Send <- s:
			default:
				close(client.Send)
				delete(h.Clients, client)
			}
		}
	}
}

func unregisterCloseConnection(c *Client) {
	c.Hub.Unregister <- c
	c.Conn.Close()
}

func (c *Client) readPump() {
	var s *model.SocketEventStruct
	defer unregisterCloseConnection(c)
	readConfig(c)

	for {
		_, payload, err := c.Conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		
		decoder := json.NewDecoder(bytes.NewReader(payload))
		err = decoder.Decode(&s)
		if err != nil {
			panic(err)
		}
		
		//go *NewWebsocketHandler()
	}
	
}