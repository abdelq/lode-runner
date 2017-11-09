package game

type Player interface {
	init(*Game)
	Move(*level, direction)
}

// States
type state uint8 // TODO
const (
	ALIVE state = iota
	FALLING
	DIGGING
	TRAPPED
)

// Directions
type direction = uint8 // TODO
const (
	NONE direction = iota
	UP
	LEFT
	DOWN
	RIGHT
)
