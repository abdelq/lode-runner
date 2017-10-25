package game

type Game struct {
	Lvl    *level
	Runner *Runner
	Guards []*Guard
}

func NewGame() *Game {
	return &Game{Guards: make([]*Guard, 0, 1)}
}

// TODO
func (g *Game) Start() {
	g.Lvl, _ = newLevel(1)
	/*if err != nil {
		log.Println(err)
		return
	}*/
	/*if g.lvl == nil {
		// TODO Broadcast failure ?
	}*/

	//g.runner.init(g.lvl.grid)
	//g.guards[0].init(g.lvl.grid) // TODO Garbage

	//for _, guard := range g.guards {}
	//g.runner.init(g.lvl.grid)

	// TODO Check for nil ?
	/*for _, player := range g.players {
		go player.init(g.lvl.grid)
	}*/

	// TODO Send a start message
}

// TODO
func (g *Game) Stop() {
	// TODO Add argument to know if it's lost by runner or guards
	// TODO Decide the winner
	// TODO Call for a leave of everyone in room
	// TODO Delete room with its game
}

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
