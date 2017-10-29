package main

import "github.com/abdelq/lode-runner/game"

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

	rooms[name] = room // TODO
	go room.listen()

	return room
}

func (r *room) listen() {
	for {
		select {
		case join := <-r.join:
			go r.joinGame(join.client, join.player)
		case client := <-r.leave:
			go r.leaveGame(client)
		case message := <-r.broadcast:
			for client := range r.clients {
				client.out <- message
			}
		}
	}
}

// TODO Rename
func (r *room) joinGame(client *client, player game.Player) {
	if _, ok := r.clients[client]; ok {
		client.out <- newErrorMessage("already in room")
		return
	}
	if r.hasPlayer(player) {
		client.out <- newErrorMessage("name already used")
		return
	}

	r.clients[client] = nil
	if r.game.Lvl != nil || player == nil { // Game started or spectator
		return
	}

	switch p := player.(type) {
	case *game.Runner:
		if r.game.Runner != nil {
			client.out <- newErrorMessage("runner already joined")
			return
		}

		r.clients[client] = p
		r.game.Runner = p
		r.broadcast <- newJoinMessage(p.Name, 0)
	case *game.Guard:
		if len(r.game.Guards) == cap(r.game.Guards) {
			client.out <- newErrorMessage("guards already joined")
			return
		}

		r.clients[client] = p
		r.game.Guards = append(r.game.Guards, p)
		r.broadcast <- newJoinMessage(p.Name, 1)
	}

	if r.game.Runner != nil && len(r.game.Guards) == cap(r.game.Guards) {
		go r.game.Start()
	}
}

// TODO Rename
func (r *room) leaveGame(client *client) {
	player := r.clients[client]
	delete(r.clients, client)

	if player == nil {
		return
	}

	switch p := player.(type) {
	case *game.Runner:
		r.game.Runner = nil
		r.broadcast <- newLeaveMessage(p.Name, 0)
	case *game.Guard:
		r.game.DeleteGuard(p)
		r.broadcast <- newLeaveMessage(p.Name, 1)
	}

	if r.game.Lvl != nil {
		if r.game.Runner == nil {
			go r.game.Stop() // TODO
		} else if len(r.game.Guards) == 0 {
			go r.game.Stop() // TODO
		}
	}
}

// TODO Rewrite
func (r *room) hasPlayer(player game.Player) bool {
	var name string
	if runner, ok := player.(*game.Runner); ok {
		name = runner.Name
	} else if guard, ok := player.(*game.Guard); ok {
		name = guard.Name
	}

	for _, player := range r.clients {
		if runner, ok := player.(*game.Runner); ok && runner.Name == name {
			return true
		} else if guard, ok := player.(*game.Guard); ok && guard.Name == name {
			return true
		}
	}
	return false
}
