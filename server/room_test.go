package main

import "testing"

// TODO
func TestListen(t *testing.T) {
	t.Run("join", func(t *testing.T) {})

	t.Run("leave", func(t *testing.T) {})

	t.Run("broadcast", func(t *testing.T) {})
}

func TestHasPlayer(t *testing.T) {
	room := newRoom("test")

	room.clients[new(client)] = nil
	if room.hasPlayer("") {
		t.Fail() // TODO
	}

	room.clients[new(client)] = &runner{name: "runner"}
	if !room.hasPlayer("runner") {
		t.Fail() // TODO
	}

	room.clients[new(client)] = &guard{name: "guard"}
	if !room.hasPlayer("guard") {
		t.Fail() // TODO
	}
}
