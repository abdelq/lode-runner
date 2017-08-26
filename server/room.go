package main

var rooms = map[string]*Room{}

type Room struct {
	join      chan *Client
	leave     chan *Client
	broadcast chan []byte
	clients   map[*Client]bool
}

func newRoom(name string) *Room {
	room := &Room{
		join:      make(chan *Client),
		leave:     make(chan *Client),
		broadcast: make(chan []byte),
		clients:   make(map[*Client]bool),
	}

	rooms[name] = room
	go room.run()

	return room
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			if _, ok := r.clients[client]; ok {
				close(client.send)
				delete(r.clients, client)
			}
		case message := <-r.broadcast:
			for client := range r.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(r.clients, client)
				}
			}
		}
	}
}
