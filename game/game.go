package game

import (
	"errors"
	msg "github.com/abdelq/lode-runner/message"
)

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

	g.broadcast <- msg.NewMessage("start", g.Level.toString()) // XXX
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

// XXX Try player.Add(g)
func (g *Game) AddPlayer(player Player) error {
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

		g.Guards[p] = struct{}{} // FIXME
		//g.broadcast <- newJoinMessage(p.name, 1) // TODO
	}

	if g.Runner != nil && len(g.Guards) == 1 { // TODO
		go g.start()
	}

	return nil
}

// XXX Try player.Remove(g)
func (g *Game) RemovePlayer(player Player) {
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
