package game

import "sort"

type Guard struct {
	Name  string
	pos   *position
	state state
}

// TODO Move to level to stop the repeated calls
func (g *Guard) init(game *Game) { // TODO Rename runnerPos
	var positions []position
	for i, row := range game.Lvl.grid {
		for j, cell := range row {
			if cell == GUARD {
				positions = append(positions, position{j, i})
			}
		}
	}

	sort.Slice(positions, func(i, j int) bool {
		return manhattanDistance(positions[i], *(game.Runner.pos)) > manhattanDistance(positions[j], *(game.Runner.pos))
	})

	g.pos = &positions[0]                   // TODO
	game.Lvl.grid[g.pos.y][g.pos.x] = EMPTY // TODO Try using cell = ?

	// TODO Maybe error management on positions
}

// TODO
func (g *Guard) Move(direction uint8, game *Game) {}
