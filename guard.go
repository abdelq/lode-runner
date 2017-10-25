package main

import "math/rand"

type guard struct {
	name  string
	pos   position
	state state
}

func (g *guard) init(grid [][]byte) {
	// TODO Move to level to stop the repeated calls
	var positions []position
	for i, row := range grid {
		for j, cell := range row {
			if cell == GUARD {
				positions = append(positions, position{j, i})
			}
		}
	}

	// TODO Verify it's not taken already
	// TODO Real random or no random ?
	g.pos = positions[rand.Intn(len(positions))]
}

func (g *guard) move(direction string, game *game) {
	// TODO
}
