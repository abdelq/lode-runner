package game

import "math"

type Player interface {
	init(*Game)
	Move(uint8, *Game)
}

type position struct{ x, y int }
type state uint8

// States
const (
	ALIVE state = iota
	FALLING
	DIGGING
	TRAPPED
)

// Directions
const (
	UP = iota
	LEFT
	DOWN
	RIGHT
)

// TODO Rename
func (p *position) set(direction uint8) {
	switch direction {
	case UP:
		p.y--
	case LEFT:
		p.x--
	case DOWN:
		p.y++
	case RIGHT:
		p.x++
	}
}

// TODO Move somewhere else + Rename
func manhattanDistance(a, b position) float64 {
	return math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
}
