package game

import "testing"

// TODO
func TestStart(t *testing.T) {}

// TODO
func TestStop(t *testing.T) {}

func TestDeleteGuard(t *testing.T) {
	game := new(Game)
	game.Guards = []*Guard{new(Guard), new(Guard), new(Guard)}
	// TODO
}

// TODO
func TestCheckCollisions(t *testing.T) {}
