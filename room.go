package main

import . "github.com/abdelq/lode-runner/game"

var rooms = make(map[string]*room)

type room struct {
	join      chan *join
	leave     chan *leave
	broadcast chan *message
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
		broadcast: make(chan *message),
		clients:   make(map[*client]Player),
		game:      NewGame(),
	}

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

			r.clients[client] = player
			if player == nil || r.game.Started() {
				continue
			}

			// TODO Broadcast join
			if err := r.game.AddPlayer(player); err != nil {
				client.out <- newMessage("error", err.Error())
				continue
			}

			if r.game.Started() {
				//r.broadcast <- newMessage("start", "TODO") // TODO
			}
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

			// TODO Broadcast leave
			r.game.RemovePlayer(player)

			if r.game.Stopped() {
				//r.broadcast <- newMessage("stop", "TODO") // TODO
			}
		case message := <-r.broadcast:
			for client := range r.clients {
				client.out <- message
			}
		}
	}
}
