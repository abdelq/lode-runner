package game

import msg "github.com/abdelq/lode-runner/message"

type Game struct {
	Level     *level
	Runner    *Runner
	Guards    map[*Guard]struct{}
	broadcast chan *msg.Message
}

func NewGame(broadcast chan *msg.Message) *Game {
	return &Game{Guards: make(map[*Guard]struct{}), broadcast: broadcast}
}

func (g *Game) Started() bool {
	return g.Level != nil
}

func (g *Game) Stopped() bool {
	return false // FIXME
}

func (g *Game) filled() bool {
	return g.Runner != nil && len(g.Guards) == 1 // FIXME
}

func (g *Game) start() {
	// Level
	g.Level, _ = newLevel(1)

	// Runner
	g.Runner.init(g.Level.landmarks)

	/* Guards */
	for guard := range g.Guards {
		guard.init(g.Level.landmarks)
	}

	// Remove unused landmarks
	for pos, _ := range g.Level.landmarks {
		for guard := range g.Guards {
			if *guard.pos == pos {
				continue
			}
		}
		delete(g.Level.landmarks, pos)
	}

	g.broadcast <- msg.NewMessage("start", g.Level.String()) // FIXME
}

func (g *Game) stop() {} // TODO

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
