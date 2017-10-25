package main

type game struct {
	lvl    *level
	runner *runner
	guards []*guard
}

func newGame() *game {
	return &game{guards: make([]*guard, 0, 1)}
}

// TODO
func (g *game) start() {
	g.lvl, _ = newLevel(1)
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
func (g *game) stop() {
	// TODO Add argument to know if it's lost by runner or guards
	// TODO Decide the winner
	// TODO Call for a leave of everyone in room
	// TODO Delete room with its game
}

func (g *game) deleteGuard(guard *guard) {
	// TODO Rename v
	for i, v := range g.guards {
		if guard == v {
			copy(g.guards[i:], g.guards[i+1:])
			g.guards[len(g.guards)-1] = nil
			g.guards = g.guards[:len(g.guards)-1]
			return
		}
	}
}

// TODO
func (g *game) checkCollisions() {}
