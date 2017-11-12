package game

type Player interface {
	Add(*Game) error
	Remove(*Game)
	init(map[position]tile)
	Move(direction, *level)
}

// States
type state uint8

const (
	ALIVE state = iota
	FALLING
	DIGGING
	TRAPPED
)

// Directions
type direction = uint8 // FIXME

const (
	NONE direction = iota
	UP
	LEFT
	DOWN
	RIGHT
)

func NewPlayer(name string, role tile) Player {
	switch role {
	case RUNNER: // FIXME Empty byte/default one from the marshalling
		return &Runner{Name: name}
	case GUARD:
		return &Guard{Name: name}
	}

	return nil
}
