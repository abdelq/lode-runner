package game

import (
	"errors"
	"sort"
)

type Guard struct {
	name   string
	pos    *position
	state  state
	Action Action
}

func (g *Guard) Add(game *Game) error {
	if len(game.guards) == 1 { // FIXME
		return errors.New("guards already joined")
	}
	if game.hasPlayer(g.name) {
		return errors.New("name already used")
	}

	game.guards[g] = struct{}{} // FIXME
	//game.broadcast <- msg.NewMessage("join", g.name) // FIXME

	if game.filled() {
		go game.start(1)
	}

	return nil
}

func (g *Guard) Remove(game *Game) {
	delete(game.guards, g)
	//game.broadcast <- msg.NewMessage("leave", g.name) // FIXME

	if game.Started() && len(game.guards) == 0 {
		go game.stop(RUNNER)
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

func (g *Guard) UpdateAction(actionType string, direction direction) {
	g.Action = Action{actionType, direction}
}
