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
	broadcast chan *msg.Message
}

func NewGame(broadcast chan *msg.Message) *Game {
	game := &Game{guards: make(map[*Guard]struct{}), broadcast: broadcast}

	// FIXME
	go func() {
		// https://gobyexample.com/tickers
		for now := range time.Tick(1 * time.Second) {
			_ = now
			if game.Started() && !game.Stopped() { // TODO Stupid
				// TODO Maybe order them (runner first + guard then)
				// Do the actions
				if game.runner.Action.ActionType == "move" {
					go game.runner.Move(game.runner.Action.Direction, game)
				} else if game.runner.Action.ActionType == "dig" {
					go game.runner.Dig(game.runner.Action.Direction, game)
				}

				for guard := range game.guards {
					if guard.Action.ActionType == "move" {
						go guard.Move(guard.Action.Direction, game)
					}
				}

				// Reset actions
				game.runner.Action = Action{"move", 0}
				for guard := range game.guards {
					guard.Action = Action{"move", 0}
				}

				game.broadcast <- msg.NewMessage("next", game.level.String()) // FIXME
			}
		}
	}()

	return game
}

func (g *Game) Started() bool { // FIXME
	return g.level != nil
}

func (g *Game) Stopped() bool { // FIXME
	return g.guards == nil // FIXME
}

func (g *Game) filled() bool { // FIXME
	return g.runner != nil && len(g.guards) == 1 // FIXME
}

func (g *Game) start(lvl int) { // FIXME
	var err error // TODO
	// Level
	g.level, err = newlevel(lvl)
	if err != nil {
		log.Println(err)
		return
		// TODO Crash people ?
	}

	// Runner
	g.runner.init(g.level.players)

	// guards
	for guard := range g.guards {
		guard.init(g.level.players)
	}

	// FIXME Remove unused landmarks
	for pos, tile := range g.level.players {
		if tile == GUARD {
			var found = false
			for guard := range g.guards {
				if *guard.pos == pos {
					found = true
				}
			}
			if !found {
				delete(g.level.players, pos)
			}
		}
	}

	g.broadcast <- msg.NewMessage("start", g.level.String()) // FIXME
}

// TODO Add argument w/ winner
func (g *Game) stop() { // FIXME
	// Force everyone to leave room
	if g.runner == nil || g.runner.health == 0 {
		g.broadcast <- msg.NewMessage("quit", "guards win") // FIXME
	} else if len(g.guards) == 0 {
		g.broadcast <- msg.NewMessage("quit", "runner wins") // FIXME
	} else { // TODO ?
		g.broadcast <- msg.NewMessage("quit", "") // FIXME
	}

	g.level = nil
	g.runner = nil
	g.guards = nil
	//close(g.broadcast)
}

func (g *Game) hasPlayer(name string) bool {
	// Runner
	if g.runner != nil && g.runner.Name == name {
		return true
	}

	// Guards
	for guard := range g.guards {
		if guard.Name == name {
			return true
		}
	}

	return false
}
