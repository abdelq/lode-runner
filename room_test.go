package main

import (
	"net"
	"testing"

	. "github.com/abdelq/lode-runner/game"
)

func TestNewRoom(t *testing.T) {
	if room := newRoom("test"); rooms["test"] != room {
		t.Error("room not in slice")
	}
}

// TODO newClient() to struct
// TODO Second join/leave after verifications
func TestListen(t *testing.T) {
	room := newRoom("test")

	t.Run("spectator", func(t *testing.T) {
		serverConn, clientConn := net.Pipe()
		spectator := newClient(serverConn)

		// Join
		room.join <- &join{spectator, nil} // First
		room.join <- &join{spectator, nil} // Second
		receiveMsg(t, clientConn, message{"error", []byte(`"already in room"`)})

		player, ok := room.clients[spectator]
		if !ok {
			t.Fail() // TODO
		}
		if player != nil {
			t.Fail() // TODO
		}

		// Broadcast
		room.broadcast <- &message{"test", []byte(`"spectator"`)}
		receiveMsg(t, clientConn, message{"test", []byte(`"spectator"`)})

		// Leave
		room.leave <- spectator // First
		room.leave <- spectator // Second
		receiveMsg(t, clientConn, message{"error", []byte(`"not in room"`)})

		if _, ok := room.clients[spectator]; ok {
			t.Fail() // TODO
		}
	})

	t.Run("runner", func(t *testing.T) {
		serverConn, clientConn := net.Pipe()
		runner := newClient(serverConn)

		// Join
		room.join <- &join{runner, new(Runner)} // First
		room.join <- &join{runner, new(Runner)} // Second
		receiveMsg(t, clientConn, message{"error", []byte(`"already in room"`)})

		player, ok := room.clients[runner]
		if !ok {
			t.Fail() // TODO
		}
		if _, ok := player.(*Runner); !ok {
			t.Fail() // TODO
		}

		// Broadcast
		room.broadcast <- &message{"test", []byte(`"runner"`)}
		receiveMsg(t, clientConn, message{"test", []byte(`"runner"`)})

		// Leave
		room.leave <- runner // First
		room.leave <- runner // Second
		receiveMsg(t, clientConn, message{"error", []byte(`"not in room"`)})

		if _, ok := room.clients[runner]; ok {
			t.Fail() // TODO
		}
	})

	t.Run("guard", func(t *testing.T) {
		serverConn, clientConn := net.Pipe()
		guard := newClient(serverConn)

		// Join
		room.join <- &join{guard, new(Guard)} // First
		room.join <- &join{guard, new(Guard)} // Second
		receiveMsg(t, clientConn, message{"error", []byte(`"already in room"`)})

		player, ok := room.clients[guard]
		if !ok {
			t.Fail() // TODO
		}
		if _, ok := player.(*Guard); !ok {
			t.Fail() // TODO
		}

		// Broadcast
		room.broadcast <- &message{"test", []byte(`"guard"`)}
		receiveMsg(t, clientConn, message{"test", []byte(`"guard"`)})

		// Leave
		room.leave <- guard // First
		room.leave <- guard // Second
		receiveMsg(t, clientConn, message{"error", []byte(`"not in room"`)})

		if _, ok := room.clients[guard]; ok {
			t.Fail() // TODO
		}
	})

	// TODO Join: When game is started
	// TODO Join: Broadcasts
	// TODO Leave: Broadcasts
}
