package http

import (
	"fmt"
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/usecase"
	"go-fp-crowdfundchat/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	 hub *Hub
	 userUsecase usecase.UserUsecase
	 projectUsecase usecase.ProjectUsecase
	 chatUsecase usecase.ChatUsecase
}

func NewChatHandler(hub *Hub, userUsecase usecase.UserUsecase, projectUsecase usecase.ProjectUsecase, chatUsecase usecase.ChatUsecase) *ChatHandler {
	return &ChatHandler{
		hub: hub,
		userUsecase: userUsecase,
		projectUsecase: projectUsecase,
		chatUsecase: chatUsecase,
	}
}


func (h *ChatHandler) CreateChatRoom (c *gin.Context) {
	var request *model.CreateChatRoomRequest
	
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42205",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	stringProjectID := strconv.Itoa(request.ProjectID)

	h.hub.Rooms[stringProjectID] = &Room{
		ProjectID: stringProjectID,
		RoomName: request.Name,
		Clients: make(map[string]*Client),	
	}	

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Request has been successfully sent.",
		"project_id": request.ProjectID,
		"room_name": request.Name,
	})
}

func (h *ChatHandler) WebsocketJoinRoom(c *gin.Context) {
	websocketConnection, err := util.Upgrade(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "400101001",
			"responseMessage": err.Error(),
		})
		return
	}

	//projectRoomID, _ := strconv.Atoi(c.Param("room_id"))
	projectRoomID := c.Param("id")
	//clientID, _ := strconv.Atoi(c.Query("user_id"))
	//clientID := c.Query("user_id")
	//username := c.Query("user_name")

	userID := c.GetInt("currentUser")
	user, err := h.userUsecase.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42208",
			"responseMessage": "Could not get the User ID",
		})
		return
	}

	clientID := strconv.Itoa(userID)
	username := user.Name


	cl := &Client{
		Connection: websocketConnection,
		ID: clientID,
		Chat: make(chan *model.ChatRequest, 10),
		ProjectID: projectRoomID,
		UserName: username,
	}

	m := &model.ChatRequest{
		Chat: "a new user joined",
		ProjectID: projectRoomID,
		UserName: username,
	}

	a:= fmt.Sprintf("chat:%s, projectid:%s, username:%s", m.Chat, m.ProjectID, m.UserName)
	fmt.Println(a)

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.ReadMessage(h.hub)
}

func (h *ChatHandler) ChatHistory(c *gin.Context) {
	var request *model.ChatHistoryRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40013",
			"responseMessage": "ID is not found",
		})
		return
	}

	chatHistory, _ := h.chatUsecase.GetChatByProjectID(request.ID)
	chatHistoryList := make([]*model.ChatHistoryResponse, len(chatHistory))
	for i, p := range chatHistory{

		chatHistoryList[i] = &model.ChatHistoryResponse{
			UserID: p.UserID,
			ProjectID: p.ProjectID,
			UserName: p.UserName,
			Chat: util.StringToDecryptToString(p.Content),
			CreatedAt: p.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Request has been successfully sent.",
		"data": chatHistoryList,
	})
}

func (h *ChatHandler) ChatRoom(c *gin.Context) {
	var request *model.ProjectDetailRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40010",
			"responseMessage": "Failed to get project detail",
		})
		fmt.Println(err.Error())
		return
	}

	projectRoomID := strconv.Itoa(request.ID)

	if _, ok := h.hub.Rooms[projectRoomID]; ok {
		websocketConnection, err := util.Upgrade(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"responseCode": "400101001",
				"responseMessage": err.Error(),
			})
			fmt.Println(err.Error())
			return
		}

		userID, _ := strconv.Atoi(c.Query("user_id"))
		user, err := h.userUsecase.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"responseCode": "42208",
				"responseMessage": "Could not get the User ID",
			})
			fmt.Println(err.Error())
			return
		}

		clientID := strconv.Itoa(userID)
		username := user.Name


		cl := &Client{
			Connection: websocketConnection,
			ID: clientID,
			Chat: make(chan *model.ChatRequest, 10),
			ProjectID: projectRoomID,
			UserName: username,
		}

		/*m := &model.ChatRequest{
			Chat: "a new user joined",
			ProjectID: projectRoomID,
			UserName: username,
			UserID: clientID,
		}*/

		/*chatHistory, _ := h.chatUsecase.GetChatByProjectID(request.ID)
		chatHistoryList := make([]*model.ChatHistoryResponse, len(chatHistory))
		for i, p := range chatHistory{
			chatHistoryList[i] = &model.ChatHistoryResponse{
				UserID: p.UserID,
				ProjectID: p.ProjectID,
				Chat: util.StringToDecryptToString(p.Content),
				CreatedAt: p.CreatedAt,
			}
		}
		cl.Connection.WriteJSON(chatHistoryList)*/

		//a:= fmt.Sprintf("chat:%s, projectid:%s, username:%s", m.Chat, m.ProjectID, m.UserName)
		//fmt.Println(a)

		h.hub.Register <- cl
		//h.hub.Broadcast <- m

		go cl.writeMessage()
		cl.ReadMessage(h.hub)
	}

	project, err := h.projectUsecase.GetProjectDetail(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40011",
			"responseMessage": "Failed to get project detail",
		})
		fmt.Println(err.Error())
		return
	}

	h.hub.Rooms[projectRoomID] = &Room{
		ProjectID: projectRoomID,
		RoomName: project.ProjectTitle,
		Clients: make(map[string]*Client),	
	}	

	fmt.Println("room is created")

	websocketConnection, err := util.Upgrade(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "400101001",
			"responseMessage": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}

	userID, _ := strconv.Atoi(c.Query("user_id"))
	user, err := h.userUsecase.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42208",
			"responseMessage": "Could not get the User ID",
		})
		fmt.Println(err.Error())
		return
	}

	clientID := strconv.Itoa(userID)
	username := user.Name

	cl := &Client{
		Connection: websocketConnection,
		ID: clientID,
		Chat: make(chan *model.ChatRequest, 10),
		ProjectID: projectRoomID,
		UserName: username,
	}

	/*m := &model.ChatRequest{
		Chat: "a new user joined",
		ProjectID: projectRoomID,
		UserName: username,
		UserID: clientID,
	}*/

	//a:= fmt.Sprintf("chat:%s, projectid:%s, username:%s", m.Chat, m.ProjectID, m.UserName)
	//fmt.Println(a)

	h.hub.Register <- cl
	//h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.ReadMessage(h.hub)

	/*c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Request has been successfully sent.",
		"project_id": projectRoomID,
		"room_name": project.ProjectTitle,
	})*/
}

func (h *ChatHandler) GetRoom(c *gin.Context) {
	rooms := make([]*model.GetRoomResponse, 0)
	
	for _, r := range h.hub.Rooms {
		rooms = append(rooms, &model.GetRoomResponse{
			ID: r.ProjectID,
			RoomName: r.RoomName,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Request has been successfully sent.",
		"data": rooms,
	})
}

func (h *ChatHandler) GetClient (c *gin.Context) {
	var clients []*model.GetClientResponse
	projectID := c.Param("id")

	_, ok := h.hub.Rooms[projectID]
	if !ok {
		clients = make([]*model.GetClientResponse, 0)
		c.JSON(http.StatusOK, gin.H{
			"responseCode": "20000",
			"responseMessage": "Request has been successfully sent.",
			"data": clients,
		})
	}

	for _, c := range h.hub.Rooms[projectID].Clients {
		clients = append(clients, &model.GetClientResponse{
			ID: c.ID,
			UserName: c.UserName,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Request has been successfully sent.",
		"data": clients,
	})
}