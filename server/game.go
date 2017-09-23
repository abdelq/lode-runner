package main

type game struct {
	lvl     *level
	players map[*client]player
}

type player interface {
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
}

func (g *game) stop() {
}
