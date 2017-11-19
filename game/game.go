package game

import (
	"flag"
	"log"
	"time"

	msg "github.com/abdelq/lode-runner/message"
)

var tick = flag.String("tick", "250ms", "duration of game tick")

type Game struct {
	level     *level
	runner    *Runner
	guards    map[*Guard]struct{}
	ticker    *time.Ticker
	broadcast chan *msg.Message
}

func NewGame(broadcast chan *msg.Message) *Game {
	dur, err := time.ParseDuration(*tick)
	if err != nil {
		dur = 250 * time.Millisecond // XXX
	}

	game := &Game{
		guards:    make(map[*Guard]struct{}),
		ticker:    time.NewTicker(dur),
		broadcast: broadcast,
	}

	go func() {
		for range game.ticker.C {
			if game.Started() {
				// Runner
				switch runner := game.runner; runner.action.Type {
				case MOVE:
					runner.move(runner.action.Direction, game)
					runner.action = action{}
				case DIG:
					runner.dig(runner.action.Direction, game)
					runner.action = action{}
				}

				// Guards
				for guard := range game.guards {
					guard.move(guard.action.Direction, game)
					guard.action = action{}
				}

				game.runner.out <- &msg.Message{"next", []byte(`{}`)} // FIXME
				game.broadcast <- msg.NewMessage("next", game.level.String())
			}
		}
	}()

	return game
}

func (g *Game) Started() bool {
	return g.level != nil // XXX
}

func (g *Game) filled() bool {
	return g.runner != nil && len(g.guards) == 0 // XXX
}

// FIXME
func (g *Game) start(lvl int) {
	level, err := newLevel(lvl)
	if err != nil {
		log.Println(err)
		// TODO Broadcast error && Stop game
		return
	}

	/* Runner */
	g.runner.init(level.players)

	/* Guards */
	for guard := range g.guards {
		guard.init(level.players)
	}

	// Delete rest
PLAYERS:
	for pos, tile := range level.players {
		if tile == GUARD {
			for guard := range g.guards {
				if guard.pos == pos {
					continue PLAYERS
				}
			}
			delete(level.players, pos)
		}
	}

	g.broadcast <- msg.NewMessage("start", level.String())
	g.level = level // XXX
}

// FIXME
func (g *Game) stop(winner tile) {
	g.ticker.Stop() // XXX Verify garbage collection

	if winner == RUNNER {
		g.broadcast <- msg.NewMessage("quit", "runner wins")
	} else {
		g.broadcast <- msg.NewMessage("quit", "runner looses")
	}

	// switch winner {
	// case RUNNER:
	// 	g.broadcast <- msg.NewMessage("quit", "runner wins") // TODO
	// case GUARD:
	// 	g.broadcast <- msg.NewMessage("quit", "runner looses") // TODO
	// default:
	// 	g.broadcast <- msg.NewMessage("quit", "draw") // TODO
	// }
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
