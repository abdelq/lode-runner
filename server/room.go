package main

import "log"

var rooms = make(map[string]*room)

type room struct {
	join, leave chan *client
	broadcast   chan *message
	clients     map[*client]bool
	game        *game
}

func newRoom(name string) *room {
	room := &room{
		join:      make(chan *client),
		leave:     make(chan *client),
		broadcast: make(chan *message),
		clients:   make(map[*client]bool),
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
			r.clients[client] = true
			log.Println(client.name, "joined", client.room)
			client.out <- &message{"join", json.RawMessage(`"` + client.name + " joined " + client.room + `"`)}

			// TODO Improve this section
			if r.game.lvl.num == 0 {
				if len(r.game.players) == 0 {
					r.game.players[client] = &runner{}
				} else if len(r.game.players) == 1 {
					r.game.players[client] = &guard{}
				}

				// Start the game
				if len(r.game.players) == 2 {
					go r.game.start()
				}
			}
		case client := <-r.leave:
			delete(r.clients, client)
			log.Println(client.name, "left", client.room)
			client.out <- &message{"leave", json.RawMessage(`"` + client.name + " left " + client.room + `"`)}

			// Stop the game
			if _, ok := r.game.players[client]; ok {
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
