package game

import (
	"errors"
	"sort"
)

type Guard struct {
	Name  string
	pos   position
	state state
}

// TODO Maybe just use landmarks directly
func (g *Guard) init(game *Game) {
	var positions []position
	for pos, tile := range game.Level.landmarks {
		if tile == GUARD {
			positions = append(positions, pos)
		}
	}

	sort.SliceStable(positions, func(i, j int) bool {
		return manhattanDist(positions[i], game.Runner.pos) >
			manhattanDist(positions[j], game.Runner.pos)
	})

	g.pos = positions[0]
	delete(game.Level.landmarks, positions[0])
}

func (g *Guard) Move(dir direction, lvl *level) {} // TODO TODO

func (g *Guard) Add(game *Game) error {
	if len(game.Guards) == 1 { // FIXME
		return errors.New("guards already joined")
	}
	if game.hasPlayer(g.Name) {
		return errors.New("name already used")
	}

	game.Guards[g] = struct{}{} // FIXME
	//game.broadcast <- newJoinMessage(p.name, 1) // TODO

	if game.filled() {
		go game.start()
	}

	return nil
}

func (g *Guard) Remove(game *Game) {
	delete(game.Guards, g)
	//game.broadcast <- newLeaveMessage(p.name, 1) // TODO

	if game.Started() && len(game.Guards) == 0 {
		go game.stop()
	}
}
