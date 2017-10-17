package main

import (
	"math/rand"
	"time"
)

type guard struct {
	name string
	pos  position
}

func (g *guard) init(grid [][]byte) {
	positions := make([]position, 0, 6)
	for i, row := range grid {
		for j, cell := range row {
			if cell == GUARD {
				positions = append(positions, position{j, i})
			}
		}
	}

	// Pick at random
	rand.Seed(time.Now().UnixNano())
	g.pos = positions[rand.Intn(len(positions))]
}

// TODO
func (g *guard) move(direction string) {}
