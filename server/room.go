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
			client, player := data.client, data.player
			r.clients[client] = nil

			// Game started or spectator
			if r.game.lvl != nil || player == nil {
				continue
			}

			game := r.game
			switch p := player.(type) {
			case *runner:
				if game.runner != nil {
					client.out <- &message{"error", json.RawMessage(`""`)} // TODO
					continue
				}

				r.clients[client] = p
				game.runner = p
				r.broadcast <- &message{"join", json.RawMessage(`"runner ` + p.name + ` joined"`)} // TODO
			case *guard:
				if len(game.guards) == cap(game.guards) {
					client.out <- &message{"error", json.RawMessage(`""`)} // TODO
					continue
				}

				r.clients[client] = p
				game.guards = append(game.guards, p)
				r.broadcast <- &message{"join", json.RawMessage(`"guard ` + p.name + ` joined"`)} // TODO
			}

			if game.runner != nil && len(game.guards) == cap(game.guards) {
				go game.start()
			}
		case client := <-r.leave:
			player := r.clients[client]
			delete(r.clients, client)

			if player == nil {
				continue
			}

			game := r.game
			switch p := player.(type) {
			case *runner:
				game.runner = nil
				r.broadcast <- &message{"leave", json.RawMessage(`"runner ` + p.name + ` left"`)} // TODO
			case *guard:
				// TODO
				game.deleteGuard(p)
				r.broadcast <- &message{"leave", json.RawMessage(`"guard ` + p.name + ` left"`)} // TODO
			}

			if game.lvl != nil && (game.runner == nil || len(game.guards) == 0) {
				go game.stop()
			}
		case message := <-r.broadcast:
			for client := range r.clients {
				client.out <- message
			}
		}
	}
}

func (r *room) hasPlayer(name string) bool {
	for _, player := range r.clients {
		if runner, ok := player.(*runner); ok && runner.name == name {
			return true
		} else if guard, ok := player.(*guard); ok && guard.name == name {
			return true
		}
	}
	return false
}
