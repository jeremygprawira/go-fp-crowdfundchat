package http

import (
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/usecase"
	"log"
)

type Room struct {
	ProjectID		string 							`json:"project_id"`
	RoomName		string							`json:"room_name"`
	Clients			map[string]*Client				`json:"clients"`
}

type Hub struct {
	Rooms			map[string]*Room
	Register		chan *Client
	Unregister		chan *Client
	Broadcast		chan *model.ChatRequest
	ChatUsecase		usecase.ChatUsecase
}

func NewHub(ChatUsecase usecase.ChatUsecase) *Hub {
	return &Hub{
		Rooms: make(map[string]*Room),
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast: make(chan *model.ChatRequest, 5),
		ChatUsecase: ChatUsecase,
	}
}

/*func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*Room),
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast: make(chan *model.ChatRequest, 5),
		ChatUsecase: NewHub().ChatUsecase,
	}
}*/

func (h *Hub) Run() {
	for {
		select {
		case cl := <- h.Register:
			/*_, ok := h.Rooms[cl.ProjectID]
			if ok {
				r := h.Rooms[cl.ProjectID]

				_, ok := r.Clients[cl.ID]
				if ok {
					r.Clients[cl.ID] = cl
				}
			}*/
			if _, ok := h.Rooms[cl.ProjectID]; ok {
				r := h.Rooms[cl.ProjectID]

				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
				}
			}
		case cl := <- h.Unregister:
			_, ok := h.Rooms[cl.ProjectID]
			if ok {
				_, ok := h.Rooms[cl.ProjectID].Clients[cl.ID]
				if ok {
					/*if len(h.Rooms[cl.ProjectID].Clients) != 0 {
						h.Broadcast <- &model.ChatRequest{
							ProjectID: cl.ProjectID,
							UserName: cl.UserName,
							Chat: "user has left the chat",
						}
					}*/
					
					delete(h.Rooms[cl.ProjectID].Clients, cl.ID)
					close(cl.Chat)
				}
			}
		case m := <- h.Broadcast:
			if _, ok := h.Rooms[m.ProjectID]; ok {
				for _, cl := range h.Rooms[m.ProjectID].Clients {
					_, err := h.ChatUsecase.PostInputChat(m)
					if err != nil {
						log.Println(err)
					}
					cl.Chat <- m
				}
			}
		}
	}
}

