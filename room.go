package main

import (
	. "github.com/abdelq/lode-runner/game"
	msg "github.com/abdelq/lode-runner/message"
)

var rooms = make(map[string]*room)

type room struct {
	join      chan *join
	leave     chan *leave
	broadcast chan *msg.Message
	clients   map[*client]Player
	game      *Game
}

type leave = client
type join struct {
	client *client
	player Player
}

func newRoom(name string) *room {
	room := &room{
		join:      make(chan *join),
		leave:     make(chan *leave),
		broadcast: make(chan *msg.Message),
		clients:   make(map[*client]Player),
	}
	room.game = NewGame(room.broadcast) // TODO

	go room.listen()
	rooms[name] = room

	return room
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

			if err := r.game.AddPlayer(player); err != nil {
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
			if player == nil /*|| r.game.Stopped()*/ { // TODO
				continue
			}

			r.game.RemovePlayer(player)
		case mssage := <-r.broadcast: // TODO
			for client := range r.clients {
				newMsg := message(*mssage)
				client.out <- &newMsg
			}
		}
	}
}
