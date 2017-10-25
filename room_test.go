package main

import (
	"net"
	"testing"

	"github.com/abdelq/lode-runner/game"
)

// TODO Uppercase
func testListen(t *testing.T) {
	room := newRoom("test")

	conn, _ := net.Pipe()
	clients := map[string]*client{
		"spectator": newClient(conn),
		"runner":    newClient(conn),
		"guard":     newClient(conn),
	}

	t.Run("join", func(t *testing.T) {
		t.Run("spectator", func(t *testing.T) {
			room.join <- &join{clients["spectator"], nil}

			player, ok := room.clients[clients["spectator"]]
			if !ok {
				t.Fail() // TODO
			}

			if player != nil {
				t.Fail() // TODO
			}
		})

		t.Run("runner", func(t *testing.T) {
			room.join <- &join{clients["runner"], new(game.Runner)}

			player, ok := room.clients[clients["runner"]]
			if !ok {
				t.Fail() // TODO
			}

			runner, ok := player.(*game.Runner)
			if !ok {
				t.Fail() // TODO
			}

			if room.game.Runner != runner {
				t.Fail() // TODO
			}
		})

		t.Run("guard", func(t *testing.T) {
			room.join <- &join{clients["guard"], new(game.Guard)}

			player, ok := room.clients[clients["guard"]]
			if !ok {
				t.Fail() // TODO
			}

			guard, ok := player.(*game.Guard)
			if !ok {
				t.Fail() // TODO
			}

			// TODO Verify guard is in array
			t.Log(guard)
		})

		// TODO Already in room
	})

	// TODO
	t.Run("broadcast", func(t *testing.T) {})

	t.Run("leave", func(t *testing.T) {
		// TODO
		room.leave <- nil
		if _, ok := room.clients[nil]; ok {
			t.Fail()
		}

		// Spectator
		room.leave <- clients["spectator"]
		if _, ok := room.clients[clients["spectator"]]; ok {
			t.Fail() // TODO
		}

		// Runner
		room.leave <- clients["runner"]
		if _, ok := room.clients[clients["runner"]]; ok {
			t.Fail() // TODO
		}
		if room.game.Runner != nil {
			t.Fail()
		}

		// TODO Guard
		room.leave <- clients["guard"]
		if _, ok := room.clients[clients["guard"]]; ok {
			t.Fail() // TODO
		}
	})
}

func TestHasPlayer(t *testing.T) {
	room := newRoom("test")

	room.clients[new(client)] = nil
	if room.hasPlayer("") {
		t.Fail() // TODO
	}

	room.clients[new(client)] = &game.Runner{Name: "runner"}
	if !room.hasPlayer("runner") {
		t.Fail() // TODO
	}

	room.clients[new(client)] = &game.Guard{Name: "guard"}
	if !room.hasPlayer("guard") {
		t.Fail() // TODO
	}
}
