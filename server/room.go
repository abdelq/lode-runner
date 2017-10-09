package main

//import "log"

var rooms = make(map[string]*room)

type room struct {
	join, leave chan *client
	broadcast   chan *message
	clients     []*client
	game        *game
}

func newRoom(name string) *room {
	room := &room{
		join:      make(chan *client),
		leave:     make(chan *client),
		broadcast: make(chan *message),
		clients:   make([]*client, 0, 2), // TODO Is it necessary, maybe add 0, 2
		game:      newGame(),
	}

	go room.listen()
	rooms[name] = room

	return room
}

func (r *room) listen() {
	for {
		select {
		case client := <-r.join:
			client.room = r
			r.clients = append(r.clients, client)
			// TODO Problematic bc specatator client etc. always a player by default fucks up the name...
			//log.Println(client.player.name, "joined", client.room) // TODO client.room is no longer a string
			//client.out <- &message{"join", &json.RawMessage(`"` + client.name + " joined " + client.room + `"`)}

			if r.game.lvl == nil {
				//r.clients[client] = &runner{} // TODO
				r.game.players = append(r.game.players, client.player)

				// Start the game
				if len(r.game.players) == cap(r.game.players) {
					go r.game.start()
				}
			}
		case client := <-r.leave:
			//delete(r.clients, client)
			// client.room = nil
			// client.player = nil
			//log.Println(client.name, "left", client.room)
			//client.out <- &message{"leave", json.RawMessage(`"` + client.name + " left " + client.room + `"`)}

			// Stop the game
			if r.game.lvl != nil && client.player != nil {
				go r.game.stop()
			}
		case message := <-r.broadcast:
			for _, client := range r.clients {
				client.out <- message
			}
		}
	}
}

func (r *room) hasPlayer(name string) bool {
	for _, client := range r.clients {
		if client.player != nil {
			switch p := client.player.(type) {
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
	}
	return false
}
