package game

import (
	"testing"

	msg "github.com/abdelq/lode-runner/message"
)

// TODO
func TestStart(t *testing.T) {
	broadcast := make(chan *msg.Message, 1)
	defer close(broadcast)

	game := NewGame(broadcast)
	game.Runner = new(Runner)
	game.Guards[new(Guard)] = struct{}{}

	game.start()

	//t.Error(game.Level)
	//t.Error(string(<-game.broadcast))
}

// TODO
func TestStop(t *testing.T) {}

// TODO
func TestHasPlayer(t *testing.T) {}
