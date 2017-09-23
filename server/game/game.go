package game

type game struct {
	lvl     *level
	players []player
}

type player interface {
	move()
}

func newGame() *game {
	return &game{
		lvl:     newLevel(1),
		players: make([]player, 0, 2),
	}
}

func (g *game) start() {
}

func (g *game) stop() {
}
