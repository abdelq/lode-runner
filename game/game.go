package game

import (
	"encoding/json"
	"flag"
	"time"

	msg "github.com/abdelq/lode-runner/message"
)

var tick = flag.String("tick", "250ms", "duration of game tick")

type Game struct {
	level     *level
	runner    *Runner
	guards    map[*Guard]struct{}
	ticker    chan bool
	broadcast chan *msg.Message
}

func NewGame(broadcast chan *msg.Message) *Game {
	return &Game{guards: make(map[*Guard]struct{}), broadcast: broadcast}
}

func (g *Game) Started() bool {
	return g.level != nil
}

func (g *Game) filled() bool {
	return g.runner != nil && len(g.guards) == 0 // XXX
}

func (g *Game) start(lvl int) {
	/* Level */
	level, err := newLevel(lvl)
	if err != nil {
		g.broadcast <- msg.NewMessage("error", err.Error())
		return // XXX Check if client will stay open. Maybe load lvl 1
	}
	g.level = level

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
	g.ticker = startTicker(g.tick)
}

func (g *Game) restart() {
	close(g.ticker)
	g.start(g.level.num)
}

func (g *Game) stop(winner tile) {
	close(g.ticker)
	g.broadcast <- msg.NewMessage("quit", "game over")
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

// TODO Improve https://stackoverflow.com/questions/17797754/ticker-stop-behaviour-in-golang
func startTicker(f func()) chan bool {
	done := make(chan bool, 1)
	go func() {
		dur, err := time.ParseDuration(*tick)
		if err != nil {
			dur = 250 * time.Millisecond // XXX
		}
		ticker := time.NewTicker(dur)

		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				f()
			case <-done:
				return
			}
		}
	}()
	return done
}

func (g *Game) tick() {
	// Fill up holes
	for pos, ticks := range g.level.holes {
		if ticks > 0 {
			g.level.holes[pos] = ticks - 1
			continue
		}

		// XXX Guard
		if tile, ok := g.level.players[pos]; ok && tile == RUNNER {
			g.runner.health--
			if g.runner.health == 0 {
				g.stop(GUARD)
				return
			}

			g.restart()
			return
		}

		g.level.tiles[pos.y][pos.x] = BRICK
		delete(g.level.holes, pos)
	}

	switch runner := g.runner; runner.action.Type {
	case MOVE:
		runner.move(runner.action.Direction, g)
		runner.action = action{}
	case DIG:
		runner.dig(runner.action.Direction, g)
		runner.action = action{}
	}

	// Guards
	for guard := range g.guards {
		guard.move(guard.action.Direction, g)
		guard.action = action{}
	}

	// XXX Need to be an array when speaking about guards?
	next := struct {
		Runner struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"runner"`
	}{}
	next.Runner.X = g.runner.pos.x
	next.Runner.Y = g.runner.pos.y

	stuff, _ := json.Marshal(next)

	g.runner.out <- &msg.Message{"next", stuff}
	g.broadcast <- msg.NewMessage("next", g.level.String())
}
