package main

import (
	"encoding/json"

	"github.com/abdelq/lode-runner/game"
)

var rooms = make(map[string]*room)

type room struct {
	join      chan *join
	leave     chan *leave
	broadcast chan *message
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
		join:      make(chan *join),
		leave:     make(chan *leave),
		broadcast: make(chan *message),
		clients:   make(map[*client]game.Player),
		game:      game.NewGame(),
	}

	go room.listen()
	rooms[name] = room

	return room
}

func (r *room) joinGame(client *client, player game.Player) {
	if _, ok := r.clients[client]; ok {
		client.out <- &message{"error", json.RawMessage(`"already in room"`)}
		return
	}
	/*else if r.hasPlayer(joinData.Name) {
		// TODO Maybe move this logic in room.go
		sender.out <- &message{"error", json.RawMessage(`"name already used"`)}
		return
	}*/

	r.clients[client] = nil
	if r.game.Lvl != nil || player == nil { // Game started or spectator
		return
	}

	switch p := player.(type) {
	case *game.Runner:
		if r.game.Runner != nil {
			client.out <- &message{"error",
				json.RawMessage(`"runner already joined"`)}
			return
		}

		r.clients[client] = p
		r.game.Runner = p
		r.broadcast <- &message{"join",
			json.RawMessage(`{"name": ` + p.Name + `, "role": 0}`)}
	case *game.Guard:
		if len(r.game.Guards) == cap(r.game.Guards) {
			client.out <- &message{"error",
				json.RawMessage(`"guards already joined"`)}
			return
		}

		r.clients[client] = p
		r.game.Guards = append(r.game.Guards, p)
		r.broadcast <- &message{"join",
			json.RawMessage(`{"name": ` + p.Name + `, "role": 1}`)}
	}

	if r.game.Runner != nil && len(r.game.Guards) == cap(r.game.Guards) {
		go r.game.Start()
	}
}

func (r *room) leaveGame(client *client) {
	player := r.clients[client]
	delete(r.clients, client)

	if player == nil {
		return
	}

	switch p := player.(type) {
	case *game.Runner:
		r.game.Runner = nil
		r.broadcast <- &message{"leave",
			json.RawMessage(`{"name": ` + p.Name + `, "role": 0}`)}
	case *game.Guard:
		r.game.DeleteGuard(p)
		r.broadcast <- &message{"leave",
			json.RawMessage(`{"name": ` + p.Name + `, "role": 1}`)}
	}

	if r.game.Lvl != nil {
		if r.game.Runner == nil {
			go r.game.Stop() // TODO
		} else if len(r.game.Guards) == 0 {
			go r.game.Stop() // TODO
		}
	}
}

func (r *room) broadcastGame(message *message) {
	for client := range r.clients {
		client.out <- message
	}
}

func (r *room) listen() {
	for {
		select {
		case data := <-r.join:
			go r.joinGame(data.client, data.player) // TODO go or no go + Rename
		case client := <-r.leave:
			go r.leaveGame(client) // TODO go or no go + Rename
		case message := <-r.broadcast:
			go r.broadcastGame(message) // TODO go or no go + Rename
		}
	}
}

func (r *room) hasPlayer(name string) bool {
	for _, player := range r.clients {
		if runner, ok := player.(*game.Runner); ok && runner.Name == name {
			return true
		} else if guard, ok := player.(*game.Guard); ok && guard.Name == name {
			return true
		}
	}
	return false
}
