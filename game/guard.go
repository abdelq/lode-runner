package game

import (
	"errors"
	"sort"

	msg "github.com/abdelq/lode-runner/message"
)

type Guard struct {
	name   string
	pos    position
	state  state
	action action
	out    chan *msg.Message
}

func (g *Guard) Join(game *Game) error {
	if len(game.guards) == 1 { // XXX
		return errors.New("guards already joined")
	}
	if game.hasPlayer(g.name) {
		return errors.New("name already used")
	}

	game.guards[g] = struct{}{}
	/*game.broadcast <- &msg.Message{"join",
		[]byte(`{"name": "", "room": "", "role": "guard"}`), // FIXME
	}*/

	if game.filled() {
		go game.start(1)
	}

	return nil
}

func (g *Guard) Leave(game *Game) {
	delete(game.guards, g)
	/*game.broadcast <- &msg.Message{"leave",
		[]byte(`{"name": "", "room": "", "role": "guard"}`), // FIXME
	}*/

	if game.Started() && len(game.guards) == 0 {
		game.stop(RUNNER)
	}
}

func (g *Guard) init(players map[position]tile) {
	var runnerPos position
	var positions []position
	for pos, tile := range players { // FIXME Don't loop over taken positions
		if tile == RUNNER {
			runnerPos = pos
		} else if tile == GUARD {
			positions = append(positions, pos)
		}
	}

	sort.Slice(positions, func(i, j int) bool {
		return manhattanDist(positions[i], runnerPos) >
			manhattanDist(positions[j], runnerPos)
	})

	g.pos = positions[0]
}

func (g *Guard) move(dir direction, game *Game) {} // TODO

// FIXME FIXME FIXME
func (g *Guard) UpdateAction(actionType uint8, direction direction) {
	g.action = action{actionType, direction}
}
