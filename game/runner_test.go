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
	game.Runner = new(Runner)
	game.Guards[new(Guard)] = struct{}{}

	game.start()

	/*t.Error(game.Level)
	game.Runner.Move(DOWN, game.Level)
	t.Error(game.Level)
	game.Runner.Move(UP, game.Level)
	t.Error(game.Level)
	game.Runner.Move(LEFT, game.Level)
	game.Runner.Move(LEFT, game.Level)
	game.Runner.Move(LEFT, game.Level)
	game.Runner.Move(LEFT, game.Level)
	game.Runner.Move(LEFT, game.Level)
	game.Runner.Move(LEFT, game.Level)
	game.Runner.Move(LEFT, game.Level)
	game.Runner.Move(LEFT, game.Level)
	game.Runner.Move(LEFT, game.Level)
	game.Runner.Move(LEFT, game.Level)
	game.Runner.Move(UP, game.Level)
	game.Runner.Move(UP, game.Level)
	game.Runner.Move(RIGHT, game.Level)
	game.Runner.Move(RIGHT, game.Level)
	game.Runner.Move(RIGHT, game.Level)
	game.Runner.Dig(RIGHT, game.Level)
	game.Runner.Move(RIGHT, game.Level)
	//game.Runner.Move(RIGHT, game.Level)
	//game.Runner.Move(RIGHT, game.Level)
	//game.Runner.Dig(LEFT, game.Level)
	game.Runner.Move(NONE, game.Level)
	game.Runner.Move(NONE, game.Level)
	game.Runner.Dig(LEFT, game.Level)
	game.Runner.Move(RIGHT, game.Level)
	//game.Runner.Move(LEFT, game.Level)
	//game.Runner.Move(LEFT, game.Level)
	//game.Runner.Move(LEFT, game.Level)
	//game.Runner.Move(LEFT, game.Level)
	//game.Runner.Move(RIGHT, game.Level)
	//game.Runner.Move(RIGHT, game.Level)
	//game.Runner.Move(RIGHT, game.Level)
	//game.Runner.Move(RIGHT, game.Level)
	//game.Runner.Move(DOWN, game.Level)
	//game.Runner.Move(DOWN, game.Level)
	//game.Runner.Move(UP, game.Level)
	t.Error(game.Level)

	//time.Sleep(80000 * 5)
	//time.Sleep(600 * time.Millisecond)
	t.Error(game.Level)*/
	/*t.Error(game.Runner.pos)
	for pos, tile := range game.Level.landmarks {
		if tile == RUNNER {
			t.Error(pos)
		}
	}*/

	//t.Error(game.Level)
	//t.Error(string(<-game.broadcast))
}
