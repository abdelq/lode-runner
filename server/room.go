package main

import "log"

var rooms = make(map[string]*room)

type room struct {
	join, leave chan *client
	broadcast   chan []byte
	clients     map[*client]bool
}

func newRoom(name string) *room {
	room := &room{
		join:      make(chan *client),
		leave:     make(chan *client),
		broadcast: make(chan []byte),
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
