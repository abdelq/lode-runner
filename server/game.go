package main

type game struct {
	lvl     *level
	players map[*client]player
}

type player interface {
	init([][]byte)
	move()
}

func newGame() *game {
	return &game{
		lvl:     &level{},
		players: make(map[*client]player),
	}
}

func (g *game) start() {
	g.lvl.init(1)
	for _, player := range g.players {
		go player.init(g.lvl.grid)
	}
}

func (g *game) stop() {
}
