package main

import "log"

var rooms = make(map[string]*room)

type room struct {
	join, leave chan *client
	broadcast   chan *message
	clients     map[*client]bool
}

func newRoom(name string) *room {
	room := &room{
		join:      make(chan *client),
		leave:     make(chan *client),
		broadcast: make(chan *message),
		clients:   make(map[*client]bool),
	}

	go room.listen()
	rooms[name] = room

	return room
}

func (r *room) listen() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true

			log.Println(client.name, "joined", client.room)
			// TODO Broadcast to clients

			// TODO Start the game
		case client := <-r.leave:
			delete(r.clients, client)

			log.Println(client.name, "left", client.room)
			// TODO Broadcast to clients

			// TODO Stop the game
		case message := <-r.broadcast:
			for client := range r.clients {
				client.out <- message
			}
		}
	}
}

func (r *room) hasClient(name string) bool {
	for client := range r.clients {
		if client.name == name {
			return true
		}
	}
	return false
}
