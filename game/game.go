package game

type Game struct {
	Lvl    *level
	Runner *Runner
	Guards []*Guard
}

func NewGame() *Game {
	// TODO Capability should be a command-line flag
	// TODO Auto min?/max on invalid values
	return &Game{Guards: make([]*Guard, 0, 1)}
}

func (g *Game) Start() {
	// Level
	g.Lvl, _ = newLevel(1)

	// Runner
	g.Runner.init(g)

	// Guards
	for _, guard := range g.Guards {
		guard.init(g)
	}
}

// TODO
func (g *Game) Stop() {}

func (g *Game) DeleteGuard(guard *Guard) {
	// TODO Rename v
	for i, v := range g.Guards {
		if guard == v {
			copy(g.Guards[i:], g.Guards[i+1:])
			g.Guards[len(g.Guards)-1] = nil
			g.Guards = g.Guards[:len(g.Guards)-1]
			return
		}
	}
}

// TODO
func (g *Game) checkCollisions() {}
