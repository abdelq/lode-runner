package main

type game struct {
	lvl     *level
	players []player
}

type player interface {
	init([][]byte)
	move(string)
}

func newGame() *game {
	return &game{players: make([]player, 0, 2)}
}

// TODO
func (g *game) start() {
	g.lvl = newLevel()

	for _, player := range g.players {
		go player.init(g.lvl.grid)
	}
}

// TODO
func (g *game) stop() {}
