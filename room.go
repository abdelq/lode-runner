package main

import (
	"sync"

	"github.com/abdelq/lode-runner/game"
	msg "github.com/abdelq/lode-runner/message"
)

var rooms = sync.Map{}

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
	room.game = game.NewGame(name, room.broadcast)

	go room.listen()

	return room
}

func findRoom(client *client) string {
	var name string
	rooms.Range(func(n, r interface{}) bool {
		if _, ok := r.(*room).clients[client]; ok {
			name = n.(string)
			return false
		}
		return true
	})
	return name
}

func (r *room) delete() {
	rooms.Range(func(n, r2 interface{}) bool {
		if r == r2.(*room) {
			rooms.Delete(n)
			return false
		}
		return true
	})
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
				if len(r.clients) == 0 {
					r.delete()
					return
				}
				continue
			}

			player.Leave(r.game)
		case msg := <-r.broadcast:
			switch msg.Event {
			case "next":
				for client, player := range r.clients {
					if player == nil {
						client.out <- msg
					}
				}
			case "quit":
				for client := range r.clients {
					client.out <- msg
					//client.close()
				}
				r.delete()
				return
			default:
				for client := range r.clients {
					client.out <- msg
				}
			}
		}
	}
}
