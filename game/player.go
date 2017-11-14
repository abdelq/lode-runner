package game

type Player interface {
	Add(*Game) error
	Remove(*Game)
	init(map[position]tile)
	Move(direction, *Game)
}

// TODO Rename fields
type Action struct {
	ActionType string
	Direction  direction
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
	case 0, RUNNER: // FIXME
		return &Runner{Name: name, health: 5}
	case GUARD:
		return &Guard{Name: name}
	}

	return nil
}
