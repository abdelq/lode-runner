package main

import "encoding/json"

var rooms = make(map[string]*room)

type room struct {
	join      chan *join
	leave     chan *leave
	broadcast chan *message
	clients   map[*client]player
	game      *game
}

type leave = client
type join struct {
	client *client
	player player
}

func newRoom(name string) *room {
	room := &room{
		join:      make(chan *join),
		leave:     make(chan *leave),
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
		case data := <-r.join:
			client := data.client // TODO

			// Game started or spectator
			if r.game.lvl != nil || data.player == nil {
				r.clients[client] = nil
				continue
			}

			guards := r.game.guards // TODO
			switch player := data.player.(type) {
			case *runner:
				if r.game.runner == nil {
					r.game.runner = player
					r.clients[client] = player
					r.broadcast <- &message{"join", json.RawMessage(`"runner ` + player.name + ` joined"`)}
				} else {
					r.clients[client] = nil
					client.out <- &message{"error", json.RawMessage(`""`)} // TODO
				}
			case *guard:
				if len(guards) < cap(guards) {
					guards = append(guards, player)
					r.clients[client] = player
					r.broadcast <- &message{"join", json.RawMessage(`"guard ` + player.name + ` joined"`)}
				} else {
					r.clients[client] = nil
					client.out <- &message{"error", json.RawMessage(`""`)} // TODO
				}
			}

			if r.game.runner != nil && len(guards) == cap(guards) {
				go r.game.start()
			}
		case client := <-r.leave:
			player := r.clients[client]
			delete(r.clients, client)
			if player == nil {
				continue
			}

			switch p := player.(type) {
			case *runner:
				r.game.runner = nil
				r.broadcast <- &message{"leave", json.RawMessage(`"runner ` + p.name + ` left"`)}
			case *guard:
				r.game.deleteGuard(p)
				r.broadcast <- &message{"leave", json.RawMessage(`"guard ` + p.name + ` left"`)}
			}

			if r.game.lvl != nil { // Game in progress
				if r.game.runner == nil || len(r.game.guards) == 0 {
					go r.game.stop()
				}
			}
		case message := <-r.broadcast:
			for client := range r.clients {
				client.out <- message
			}
		}
	}
}

// TODO
func (r *room) hasPlayer(name string) bool {
	for _, player := range r.clients {
		switch p := player.(type) {
		case *runner:
			if p.name == name {
				return true
			}
		case *guard:
			if p.name == name {
				return true
			}
		}
	}
	return false
}
