package main

import "bytes"

type runner struct {
	name string
	pos  position
}

func (r *runner) init(grid [][]byte) {
	for i := len(grid) - 1; i >= 0; i-- {
		if j := bytes.IndexRune(grid[i], RUNNER); j != -1 {
			r.pos.x, r.pos.y = j, i
			return
		}
	}
}

// TODO
func (r *runner) move(direction string) {}

// TODO
func (r *runner) dig(direction string) {}
