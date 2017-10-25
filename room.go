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
			if _, ok := r.clients[client]; ok {
				client.out <- &message{"error", json.RawMessage(`"already in room"`)}
				continue
			}
			/*else if r.hasPlayer(joinData.Name) {
				// TODO Maybe move this logic in room.go
				sender.out <- &message{"error", json.RawMessage(`"name already used"`)}
				return
			}*/

			r.clients[client] = nil
			if r.game.lvl != nil || player == nil { // Game started or spectator
				continue
			}

			game := r.game // TODO
			switch p := player.(type) {
			case *runner:
				if game.runner != nil {
					client.out <- &message{"error",
						json.RawMessage(`"runner already joined"`)}
					continue
				}

				r.clients[client] = p
				game.runner = p
				r.broadcast <- &message{"join",
					json.RawMessage(`{"name": ` + p.name + `, "role": 0}`)}
			case *guard:
				if len(game.guards) == cap(game.guards) {
					client.out <- &message{"error",
						json.RawMessage(`"guards already joined"`)}
					continue
				}

				r.clients[client] = p
				game.guards = append(game.guards, p)
				r.broadcast <- &message{"join",
					json.RawMessage(`{"name": ` + p.name + `, "role": 1}`)}
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

			switch p := player.(type) {
			case *runner:
				r.game.runner = nil
				r.broadcast <- &message{"leave",
					json.RawMessage(`{"name": ` + p.name + `, "role": 0}`)}
			case *guard:
				r.game.deleteGuard(p)
				r.broadcast <- &message{"leave",
					json.RawMessage(`{"name": ` + p.name + `, "role": 1}`)}
			}

			if r.game.lvl != nil {
				if r.game.runner == nil {
					go r.game.stop() // TODO
				} else if len(r.game.guards) == 0 {
					go r.game.stop() // TODO
				}
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
