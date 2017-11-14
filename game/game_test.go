package game

import (
	"testing"
	//msg "github.com/abdelq/lode-runner/message"
)

func TestStart(t *testing.T) {
	/*game := &Game{
		Runner:    new(Runner),
		Guards:    map[*Guard]struct{}{new(Guard): struct{}{}},
		broadcast: make(chan *msg.Message, 1),
	}*/

	//game.start()

	// TODO TODO TODO
}

func TestStop(t *testing.T) {} // TODO

func TestHasPlayer(t *testing.T) {
	game := &Game{guards: make(map[*Guard]struct{})}

	// TODO Improve + Replace t.Fail by t.Error
	if game.hasPlayer("runner") {
		t.Fail()
	}
	game.runner = &Runner{Name: "runner"}
	if !game.hasPlayer("runner") {
		t.Fail()
	}

	if game.hasPlayer("guard") {
		t.Fail()
	}
	game.guards[&Guard{Name: "guard"}] = struct{}{}
	if !game.hasPlayer("guard") {
		t.Fail()
	}
}
