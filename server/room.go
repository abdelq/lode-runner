package main

var rooms = map[string]*Room{}

type Room struct {
	join      chan *Client
	leave     chan *Client
	broadcast chan string
	clients   map[*Client]bool
}

func newRoom(name string) *Room {
	room := &Room{
		join:      make(chan *Client),
		leave:     make(chan *Client),
		broadcast: make(chan string),
		clients:   make(map[*Client]bool),
	}

	go room.listen()
	rooms[name] = room

	return room
}

func (r *Room) listen() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
			}
		case msg := <-r.broadcast:
			for client := range r.clients {
				client.out <- msg
			}
		}
	}
}
