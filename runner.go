package main

type runner struct {
	name string
	pos  position
}

func (r *runner) init(grid [][]byte) {
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] == RUNNER {
				r.pos.x, r.pos.y = j, i
				return
			}
		}
	}
}

// TODO
func (r *runner) move(direction string) {}

// TODO
func (r *runner) dig(direction string) {}
