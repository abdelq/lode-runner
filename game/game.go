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
		ticker:    time.NewTicker(250 * time.Millisecond), // TODO Right duration
		broadcast: broadcast,
	}

	go func() {
		for range game.ticker.C {
			if game.Started() {
				// Runner
				switch action := game.runner.action; action.actionType {
				case "move":
					game.runner.move(action.direction, game)
				case "dig":
					game.runner.dig(action.direction, game)
				}
				game.runner.action = action{"move", NONE} // XXX

				// Guards
				for guard := range game.guards {
					guard.move(guard.action.direction, game)
					guard.action = action{"move", NONE} // XXX
				}

				game.broadcast <- msg.NewMessage("next", game.level.String()) // XXX
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

func (g *Game) stop(winner tile) {
	g.ticker.Stop() // XXX Verify garbage collection

	if winner == RUNNER {
		g.broadcast <- msg.NewMessage("quit", "runner wins")
	} else {
		g.broadcast <- msg.NewMessage("quit", "runner looses")
	}
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
