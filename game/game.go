package game

import (
	"encoding/json"
	"flag"
	"strings"
	"time"

	msg "github.com/abdelq/lode-runner/message"
)

var tick = flag.String("tick", "250ms", "duration of game tick")

type Game struct {
	room      string
	level     *level
	runner    *Runner
	guards    map[*Guard]struct{}
	ticker    chan bool
	broadcast chan *msg.Message
}

func NewGame(room string, broadcast chan *msg.Message) *Game {
	return &Game{
		room:      room,
		guards:    make(map[*Guard]struct{}),
		broadcast: broadcast,
	}
}

func (g *Game) Started() bool {
	return g.level != nil
}

func (g *Game) filled() bool {
	return g.runner != nil && len(g.guards) == 0 // XXX
}

func (g *Game) start(lvl int) {
	if g.ticker != nil {
		g.ticker <- true
	}

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
	/*for guard := range g.guards {
		guard.init(level.players)
	}*/

	// Delete rest
	/*PLAYERS:
	for pos, tile := range level.players {
		if tile == GUARD {
			for guard := range g.guards {
				if guard.pos == pos {
					continue PLAYERS
				}
			}
			delete(level.players, pos)
		}
	}*/

	tiles := g.level.stringTiles()
	for i := range tiles {
		tiles[i] = strings.Replace(tiles[i], string(TRAP), string(BRICK), -1)
	}

	start := struct {
		Tiles []string `json:"tiles"`
		Room  string   `json:"room"`
		Lives uint8    `json:"lives"`
		Level int      `json:"level"`
	}{tiles, g.room, g.runner.health, g.level.num}
	stuff, _ := json.Marshal(start)

	g.broadcast <- &msg.Message{"start", stuff}
	g.ticker = startTicker(g.tick)
}

func (g *Game) restart() {
	//close(g.ticker)
	//g.ticker <- true
	g.start(g.level.num)
}

func (g *Game) stop(winner tile) {
	//close(g.ticker)
	g.ticker <- true
	g.broadcast <- msg.NewMessage("quit", g.room)
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

func (g *Game) Kill() {
	g.runner.health--
	if g.runner.health == 0 {
		g.stop(GUARD)
		return
	}

	g.restart()
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

		//defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				f()
			case <-done:
				ticker.Stop()
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
		if tile, ok := g.level.players.Load(pos); ok &&
			tile.(uint8) == RUNNER {
			g.Kill()
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

	tiles := g.level.stringTiles()
	for i := range tiles {
		tiles[i] = strings.Replace(tiles[i], string(TRAP), string(BRICK), -1)
	}

	next2 := struct {
		Tiles []string `json:"tiles"`
		Room  string   `json:"room"`
		Lives uint8    `json:"lives"`
		Level int      `json:"level"`
	}{tiles, g.room, g.runner.health, g.level.num}
	stuff2, _ := json.Marshal(next2)

	g.broadcast <- &msg.Message{"next", stuff2}
}
