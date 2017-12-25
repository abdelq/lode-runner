package game

import msg "github.com/abdelq/lode-runner/message"

type Player interface {
	Join(*Game) error
	Leave(*Game)
	Move(uint8)
}

type action struct {
	Type, Direction uint8
}

// Actions
const (
	MOVE = iota
	DIG
)

// States
const (
	ALIVE = iota
	FALLING
)

// Directions
const (
	NONE = iota
	UP
	LEFT
	DOWN
	RIGHT
)

func NewPlayer(joinMsg *msg.JoinMessage, out chan *msg.Message) Player {
	switch joinMsg.Role {
	case RUNNER, 0:
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
