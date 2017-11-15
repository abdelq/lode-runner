package game

import "testing"

func TestNewGame(t *testing.T) {} // TODO

func TestStart(t *testing.T) {} // TODO

func TestStop(t *testing.T) {} // TODO

// TODO Improve + Replace t.Fail by t.Error
func TestHasPlayer(t *testing.T) {
	game := &Game{guards: make(map[*Guard]struct{})}

	if game.hasPlayer("runner") {
		t.Fail()
	}
	game.runner = &Runner{name: "runner"}
	if !game.hasPlayer("runner") {
		t.Fail()
	}

	if game.hasPlayer("guard") {
		t.Fail()
	}
	game.guards[&Guard{name: "guard"}] = struct{}{}
	if !game.hasPlayer("guard") {
		t.Fail()
	}
}
