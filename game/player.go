package game

type Player interface {
	Add(*Game) error
	Remove(*Game)
	init(*Game)
	Move(direction, *level)
}

type state uint8
type direction = uint8 // TODO

// States
const (
	ALIVE state = iota
	FALLING
	DIGGING
	TRAPPED
)

// Directions
const (
	NONE direction = iota
	UP
	LEFT
	DOWN
	RIGHT
)

func NewPlayer(name string, role uint8) Player {
	switch role {
	case 0:
		return &Runner{Name: name}
	case 1:
		return &Guard{Name: name}
	}

	return nil
}
