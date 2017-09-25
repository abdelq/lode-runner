package main

type game struct {
	lvl     *level
	players []player
}

type player interface {
	init([][]byte)
	move()
}

func newGame() *game {
	return &game{
		players: make([]player, 0, 2),
	}
}

func (g *game) start() {
	g.lvl = &level{}
	g.lvl.init(1)

	for _, player := range g.players {
		go player.init(g.lvl.grid)
	}
}

func (g *game) stop() {
}
