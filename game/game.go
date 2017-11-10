package game

import "github.com/abdelq/lode-runner/message"

type Game struct {
	Level     *level
	Runner    *Runner
	Guards    map[*Guard]struct{}
	broadcast chan *message.Message
}

func NewGame(broadcast chan *message.Message) *Game {
	return &Game{Guards: make(map[*Guard]struct{}), broadcast: broadcast}
}

func (g *Game) Started() bool {
	return g.Level != nil
}

func (g *Game) Stopped() bool {
	return false // FIXME
	//return g.Level != nil
}

func (g *Game) filled() bool {
	return g.Runner != nil && len(g.Guards) == 1 // FIXME
}

func (g *Game) start() {
	// Level
	g.Level, _ = newLevel(1)
	g.Level.game = g // FIXME

	// Runner
	g.Runner.init(g)

	// Guards
	for guard := range g.Guards {
		guard.init(g)
	}

	g.broadcast <- message.NewMessage("start", g.Level.toString()) // XXX
}

func (g *Game) stop() {} // XXX Cleanup + Broadcast

func (g *Game) hasPlayer(name string) bool {
	// Runner
	if g.Runner != nil && g.Runner.Name == name {
		return true
	}

	// Guards
	for guard := range g.Guards {
		if guard.Name == name {
			return true
		}
	}

	return false
}
