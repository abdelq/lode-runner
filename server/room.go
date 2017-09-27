package main

import (
	"encoding/json"
	"log"
)

var rooms = make(map[string]*room)

type room struct {
	join, leave chan *client
	broadcast   chan *message
	clients     map[*client]player
	game        *game
}

func newRoom(name string) *room {
	room := &room{
		join:      make(chan *client),
		leave:     make(chan *client),
		broadcast: make(chan *message),
		clients:   make(map[*client]player),
		game:      newGame(),
	}

	go room.listen()
	rooms[name] = room

	return room
}

func (r *room) listen() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = nil
			log.Println(client.name, "joined", client.room)
			client.out <- &message{"join", json.RawMessage(`"` + client.name + " joined " + client.room + `"`)}

			if r.game.lvl == nil {
				r.clients[client] = &runner{} // TODO
				r.game.players = append(r.game.players, r.clients[client])

				// Start the game
				if len(r.game.players) == cap(r.game.players) {
					go r.game.start()
				}
			}
		case client := <-r.leave:
			delete(r.clients, client)
			log.Println(client.name, "left", client.room)
			client.out <- &message{"leave", json.RawMessage(`"` + client.name + " left " + client.room + `"`)}

			// Stop the game
			if r.game.lvl != nil && r.clients[client] != nil {
				go r.game.stop()
			}
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
