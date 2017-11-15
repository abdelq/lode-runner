package game

type state uint8
type Player interface {
	Add(*Game) error
	Remove(*Game)
	init(map[position]tile)
	move(direction, *Game)
	UpdateAction(string, direction)
}

// TODO Rename fields
type direction = uint8 // XXX
type action struct {
	actionType string
	direction  direction
}

// XXX States
const (
	ALIVE state = iota
	FALLING
	DIGGING
	TRAPPED
)

// XXX Directions
const (
	NONE direction = iota
	UP
	LEFT
	DOWN
	RIGHT
)

func NewPlayer(name string, role tile) Player {
	switch role {
	case 0, RUNNER: // XXX
		return &Runner{name: name, health: 5}
	case GUARD:
		return &Guard{name: name}
	}

	return nil
}
