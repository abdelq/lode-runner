package game

import "errors"

type Game struct {
	Level     *level
	Runner    *Runner
	Guards    map[*Guard]bool
	broadcast chan *message // TODO TODO TODO
}

func NewGame(broadcast chan *message) *Game { // TODO TODO TODO
	return &Game{Guards: make(map[*Guard]bool), broadcast: broadcast}
}

func (g *Game) start() {
	// Level
	g.Level, _ = newLevel(1)
	g.Level.game = g // TODO Temporary

	// Runner
	g.Runner.init(g)

	// Guards
	for guard := range g.Guards {
		guard.init(g)
	}

	//g.broadcast <- newMessage("start", "") // TODO
}

func (g *game) Started() bool {
	return g.Level != nil
}

func (g *Game) stop() {} // TODO TODO Cleanup + Broadcast

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

func (g *Game) AddPlayer(player player) error {
	switch p := player.(type) {
	case *Runner:
		if g.Runner != nil {
			return errors.New("runner already joined")
		}
		if g.hasPlayer(p.Name) {
			return errors.New("name already used")
		}

		g.Runner = p
		//g.broadcast <- newJoinMessage(p.name, 0) // TODO
	case *Guard:
		if len(g.Guards) == 1 { // TODO
			return errors.New("guards already joined")
		}
		if g.hasPlayer(p.Name) {
			return errors.New("name already used")
		}

		g.Guards[p] = true
		//g.broadcast <- newJoinMessage(p.name, 1) // TODO
	}

	if g.Runner != nil && len(g.Guards) == 1 { // TODO
		go g.start()
	}

	return nil
}

func (g *Game) RemovePlayer(player player) {
	switch p := player.(type) {
	case *Runner:
		g.Runner = nil
		//g.broadcast <- newLeaveMessage(p.name, 0) // TODO
	case *Guard:
		delete(g.Guards, p)
		//g.broadcast <- newLeaveMessage(p.name, 1) // TODO
	}

	if g.Started() && (g.Runner == nil || len(g.Guards) == 0) {
		go g.stop()
	}
}
