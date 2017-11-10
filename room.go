package main

import (
	"github.com/abdelq/lode-runner/game"
	. "github.com/abdelq/lode-runner/message"
)

var rooms = make(map[string]*room)

type room struct {
	join      chan *join
	leave     chan *leave
	broadcast chan *Message
	clients   map[*client]game.Player
	game      *game.Game
}

type leave = client
type join struct {
	client *client
	player game.Player
}

func newRoom(name string) *room {
	broadcast := make(chan *Message)
	room := &room{
		join:      make(chan *join),
		leave:     make(chan *leave),
		broadcast: broadcast,
		clients:   make(map[*client]game.Player),
		game:      game.NewGame(broadcast),
	}

	go room.listen()
	rooms[name] = room

	return room
}

func findRoom(client *client) string {
	for name, room := range rooms {
		if _, ok := room.clients[client]; ok {
			return name
		}
	}
	return ""
}

func (r *room) listen() {
	for {
		select {
		case join := <-r.join:
			client, player := join.client, join.player
			if _, ok := r.clients[client]; ok {
				client.out <- newMessage("error", "already in room")
				continue
			}

			r.clients[client] = nil
			if player == nil || r.game.Started() {
				continue
			}

			if err := player.Add(r.game); err != nil {
				client.out <- newMessage("error", err.Error())
				continue
			}
			r.clients[client] = player
		case client := <-r.leave:
			player := r.clients[client]
			if _, ok := r.clients[client]; !ok {
				client.out <- newMessage("error", "not in room")
				continue
			}

			delete(r.clients, client)
			if player == nil || r.game.Stopped() {
				continue
			}

			player.Remove(r.game)
		case msg := <-r.broadcast:
			for client := range r.clients { // FIXME
				msg2 := message(*msg)
				client.out <- &msg2
			}
		}
	}
}
