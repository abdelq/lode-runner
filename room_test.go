package main

import (
	"testing"
)

func TestNewRoom(t *testing.T) {
	room := newRoom("test")

	if rooms["test"] != room {
		t.Fail() // TODO
	}
}

func TestRoomListen(t *testing.T) {
	//room := newRoom("test")

	// TODO
	t.Run("join", func(t *testing.T) {})

	// TODO
	t.Run("leave", func(t *testing.T) {})

	// TODO
	t.Run("broadcast", func(t *testing.T) {})
}

// TODO
func TestRoomJoin(t *testing.T) {}

// TODO
func TestRoomLeave(t *testing.T) {}

// TODO
func TestHasPlayer(t *testing.T) {}
