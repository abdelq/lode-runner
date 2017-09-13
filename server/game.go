package main

type Game struct {
	player *Client
	enemy  *Client
}

func newGame(clients map[*Client]bool) *Game {
	game := &Game{}

	for client, _ := range clients {
		if game.player == nil {
			game.player = client
		} else if game.enemy == nil {
			game.enemy = client
		}
	}

	go game.start()

	return game
}

func (g *Game) start() {
	room := rooms[g.player.room]

	room.broadcast <- "Game started"
}
