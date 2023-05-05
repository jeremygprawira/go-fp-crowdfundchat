package http

import (
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/usecase"
	"go-fp-crowdfundchat/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	chatUsecase usecase.ChatUsecase
}

func NewChatHandler (chatUsecase usecase.ChatUsecase) *ChatHandler {
	return &ChatHandler{chatUsecase}
}

func (h *ChatHandler) InputChat(c *gin.Context) {
	var chatRequest model.InputChatRequest

	err := c.ShouldBindJSON(&chatRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42210",
			"responseMessage": err.Error(),
		})
		return
	}

	chat, err := h.chatUsecase.PostInputChat(&chatRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40030",
			"responseMessage": "Usecase PostInputChat is not working properly",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "User has been registered successfully",
		"originUserID": chat.OriginUserID,
		"destinationUserID": chat.DestinationUserID,
		"message": chat.Message,
	})
}

func (h *ChatHandler) ChatHistory(c *gin.Context) {
	var request *model.ChatHistoryRequest
	
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40010",
			"responseMessage": err.Error(),
		})
		return
	}

	chat, err := h.chatUsecase.GetChatHistory(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40011",
			"responseMessage": err.Error(),
		})
		return
	}

	chatHistoryList := make([]*model.ChatHistoryResponse, len(chat))
    for i, p := range chat {
        chatHistoryList[i] = &model.ChatHistoryResponse{
            OriginUserID: p.OriginUserID,
			DestinationUserID: p.DestinationUserID,
			Message: util.StringToDecryptToString(p.Message),
            CreatedAt: p.CreatedAt,
        }
    }

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Project detail has been successfully retrieved.",
		"data": chatHistoryList,
	})
}
/*
var clients = make(map[*model.Client]bool)
var broadcast = make(chan *model.Chat)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {return true},
}*/

/*func (h *ChatHandler) receiver(client *model.Client) {
//func receiver(client *model.Client) {
	for {
		_, p, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		m := &model.Message{}
		err = json.Unmarshal(p, m)
		if err != nil {
			log.Println("error while unmarshaling chat", err)
			continue
		}

		if m.Type == "bootup" {
			fmt.Println("client successfully mapped", &client, client)
		} else {
			fmt.Println("received message", m.Type, m.Chat)
			c := m.Chat
			id, err := h.chatUsecase.PostInputChat(&c)
			if err != nil {
				log.Println("error while saving chat in redis", err)
				return
			}

			broadcast <- id
		}
	}
}

func Broadcaster() {
	for {
		message := <- broadcast
		fmt.Println("new message: ", message)

		for client := range clients {
			fmt.Println("from:", message.OriginUserID, "to:", message.DestinationUserID)
			err := client.Conn.WriteJSON(message)
				if err != nil {
					log.Printf("websocket error: %s", err)
					client.Conn.Close()
					delete(clients, client)
				}
		}
	}
}

func HandleConnection(c *gin.Context) {
	fmt.Println(c.Request.Host, c.Request.URL.Query())

	websocket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("err")
	}

	client := &model.Client{Conn: websocket}
	clients[client] = true
	fmt.Println("clients", len(clients), clients, websocket.RemoteAddr())

	(*ChatHandler).receiver(&ChatHandler{}, client)
	fmt.Println("exiting", websocket.RemoteAddr().String())
}*/
/*
func (h *ChatHandler) ChatOperation(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42210",
			"responseMessage": err.Error(),
		})
		return
	}


}*/