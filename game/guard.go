package game

import "sort"

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
