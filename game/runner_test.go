package game

import (
	"testing"
	//"time"

	msg "github.com/abdelq/lode-runner/message"
)

// TODO
func TestMove(t *testing.T) {
	broadcast := make(chan *msg.Message, 1)
	defer close(broadcast)

	game := NewGame(broadcast)
	game.runner = new(Runner)
	game.guards[new(Guard)] = struct{}{}

	//game.start()
}
