package game

type game struct{}

func newGame() *game {
	return &game{}
}

func (g *game) start() {}

func (g *game) stop() {}
