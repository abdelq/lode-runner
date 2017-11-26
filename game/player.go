package game

import msg "github.com/abdelq/lode-runner/message"

type state uint8
type Player interface {
	Join(*Game) error
	Leave(*Game)
	//init(map[position]tile)
	//move(direction, *Game)
	UpdateAction(uint8, direction)
}

// TODO Rename fields
type direction = uint8 // XXX
type action struct {
	Type      uint8
	Direction direction
}

// XXX Actions
const (
	MOVE = iota
	DIG
)

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

func NewPlayer(name string, role tile, out chan *msg.Message) Player {
	switch role {
	case 0, RUNNER: // XXX
		return &Runner{name: name, action: action{}, health: 5, out: out}
		//return &Runner{name: name, action: action{}, health: 5}
	case GUARD:
		return &Guard{name: name, action: action{}, out: out}
		//return &Guard{name: name, action: action{}}
	}

	return nil
}
