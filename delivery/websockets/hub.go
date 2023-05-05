package websockets

import "go-fp-crowdfundchat/model"

type Room struct {
	ID 			string 					`json:"id"`
	//Name		string					`json:"name"`
	Clients		map[string]*Client		`json:"clients"`

	User		model.User				`json:"user"`
}

type Hub struct {
	Rooms		map[string]*Room
	Clients	   map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
		Clients: 	make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
			case client := <-h.Register:
				UserRegisterEvent(h, client)
			//case client := <-h.Unregister:
				//HandleUserDisconnectEvent(h, client)
		}
	}
}
/*func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}

					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Msg)
				}
			}

		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {

				for _, cl := range h.Rooms[m.RoomID].Clients {
					cl.Msg <- m
				}
			}
		}
	}
}*/
