package main

import (
	"github.com/abdelq/lode-runner/game"
	msg "github.com/abdelq/lode-runner/message"
)

var rooms = make(map[string]*room)

type room struct {
	join      chan *join
	leave     chan *leave
	broadcast chan *msg.Message
	clients   map[*client]game.Player
	game      *game.Game
}

type leave = client
type join struct {
	client *client
	player game.Player
}

func newRoom(name string) *room {
	room := &room{
		join:      make(chan *join, 7),
		leave:     make(chan *leave, 7),
		broadcast: make(chan *msg.Message, 10), // XXX
		clients:   make(map[*client]game.Player),
	}
	room.game = game.NewGame(room.broadcast)

	go room.listen()
	rooms[name] = room

	return room
}

func (r *room) delete() {
	for name, room := range rooms {
		if room == r {
			delete(rooms, name)
			return
		}
	}
}

func (r *room) listen() {
	for {
		select {
		case join := <-r.join:
			client, player := join.client, join.player
			if _, ok := r.clients[client]; ok {
				client.out <- msg.NewMessage("error", "already in room")
				continue
			}

			r.clients[client] = nil
			if player == nil || r.game.Started() {
				continue
			}

			if err := player.Join(r.game); err != nil {
				client.out <- msg.NewMessage("error", err.Error())
				continue
			}
			r.clients[client] = player
		case client := <-r.leave:
			player := r.clients[client]
			if _, ok := r.clients[client]; !ok {
				client.out <- msg.NewMessage("error", "not in room")
				continue
			}

			delete(r.clients, client)
			if player == nil {
				if len(r.clients) > 0 {
					continue
				}
				r.delete()
				break // XXX
			}

			player.Leave(r.game)
		case msg := <-r.broadcast:
			switch msg.Event {
			case "next": // XXX
				for client, player := range r.clients {
					if _, ok := player.(*game.Runner); !ok {
						client.out <- msg
					}
				}
			case "quit": // XXX
				for client := range r.clients {
					client.out <- msg
					client.close()
				}
				r.delete()
				break // XXX
			default:
				for client := range r.clients {
					client.out <- msg
				}
			}
		}
	}
}
