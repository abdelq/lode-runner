package game

import (
	"log"
	"time"

	msg "github.com/abdelq/lode-runner/message"
)

type Game struct {
	level     *level
	runner    *Runner
	guards    map[*Guard]struct{}
	ticker    *time.Ticker
	broadcast chan *msg.Message
}

func NewGame(broadcast chan *msg.Message) *Game {
	game := &Game{
		guards:    make(map[*Guard]struct{}),
		ticker:    time.NewTicker(time.Second), // TODO Choose appropriate duration
		broadcast: broadcast,
	}

	go func() {
		for range game.ticker.C {
			if game.Started() {
				// Runner
				switch action := game.runner.Action; action.ActionType {
				case "move":
					game.runner.Move(action.Direction, game) // XXX
				case "dig":
					game.runner.Dig(action.Direction, game) // XXX
				}

				// Guards
				for guard := range game.guards {
					guard.Move(guard.Action.Direction, game) // XXX
				}
			}
		}
	}()

	return game
}

func (g *Game) Started() bool {
	return g.level != nil // XXX
}

/*func (g *Game) Stopped() bool {
	return g.guards == nil // XXX
}*/

func (g *Game) filled() bool {
	return g.runner != nil && len(g.guards) == 1 // XXX
}

func (g *Game) start(lvl int) {
	level, err := newLevel(lvl)
	if err != nil {
		log.Println(err)
		// TODO Broadcast error & Stop game
		return // XXX
	}

	/* Runner */
	g.runner.init(level.players)

	/* Guards */
	for guard := range g.guards {
		guard.init(level.players)
	}

	// Delete rest
OUTER: // TODO Rename
	for pos, tile := range level.players {
		if tile == GUARD {
			for guard := range g.guards {
				if *guard.pos == pos {
					continue OUTER
				}
			}
			delete(level.players, pos)
		}
	}

	// XXX
	g.broadcast <- msg.NewMessage("start", level.String())
	g.level = level // XXX Placement + Possible ticker issue
}

func (g *Game) stop(winner tile) {
	switch winner {
	case RUNNER:
		g.broadcast <- msg.NewMessage("quit", "runner wins") // TODO
	case GUARD:
		g.broadcast <- msg.NewMessage("quit", "guards win") // TODO
	default:
		g.broadcast <- msg.NewMessage("quit", "draw") // TODO
	}

	// TODO Verify garbage collection
	g.ticker.Stop() // XXX
}

func (g *Game) hasPlayer(name string) bool {
	// Runner
	if g.runner != nil && g.runner.name == name {
		return true
	}

	// Guards
	for guard := range g.guards {
		if guard.name == name {
			return true
		}
	}

	return false
}
