package game

type Player interface {
	init(*Game)
	Move(direction, *level)
}

func NewPlayer(name string, role uint8) Player { // TODO *Player vs Player
	switch role {
	case 0:
		return &Runner{Name: name}
	case 1:
		return &Guard{Name: name}
	}

	return nil
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
