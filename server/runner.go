package main

import (
	"bytes"
)

type runner struct {
	pos struct{ x, y int }
}

func (r *runner) init(grid [][]byte) {
	for i, line := range grid {
		if j := bytes.IndexRune(line, RUNNER); j != -1 {
			r.pos.x = i
			r.pos.y = j
			return
		}
	}
}

func (r *runner) move() {
}
