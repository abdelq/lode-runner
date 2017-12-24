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

func NewPlayer(joinMsg *msg.JoinMessage, out chan *msg.Message) Player {
	switch joinMsg.Role {
	case RUNNER, 0: // XXX
		return &Runner{
			name:     joinMsg.Name,
			action:   action{},
			health:   5,
			startLvl: joinMsg.Level,
			out:      out,
		}
	case GUARD:
		return &Guard{
			name:   joinMsg.Name,
			action: action{},
			out:    out,
		}
	}

	return nil
}
