package game

import (
	"errors"
	"sort"
	//msg "github.com/abdelq/lode-runner/message"
)

type Guard struct {
	Name  string
	pos   *position
	state state
}

func (g *Guard) Add(game *Game) error {
	if len(game.Guards) == 1 { // FIXME
		return errors.New("guards already joined")
	}
	if game.hasPlayer(g.Name) {
		return errors.New("name already used")
	}

	game.Guards[g] = struct{}{} // FIXME
	//game.broadcast <- msg.NewMessage("join", g.Name) // FIXME

	if game.filled() {
		go game.start(1)
	}

	return nil
}

func (g *Guard) Remove(game *Game) {
	delete(game.Guards, g)
	//game.broadcast <- msg.NewMessage("leave", g.Name) // FIXME

	if game.Started() && len(game.Guards) == 0 {
		go game.stop()
	}
}

func (g *Guard) init(landmarks map[position]tile) { // XXX
	var runnerPos position
	var positions []position
	for pos, tile := range landmarks {
		if tile == RUNNER {
			runnerPos = pos
		} else if tile == GUARD {
			positions = append(positions, pos)
		}
	}

	sort.SliceStable(positions, func(i, j int) bool {
		return manhattanDist(positions[i], runnerPos) >
			manhattanDist(positions[j], runnerPos)
	})

	g.pos = &positions[0]
}

// TODO Broadcast
func (g *Guard) Move(dir direction, game *Game) {} // TODO
