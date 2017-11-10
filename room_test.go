package main

import (
	"net"
	"testing"

	"github.com/abdelq/lode-runner/game"
	. "github.com/abdelq/lode-runner/message"
)

func testListenSpectator(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	spectator := newClient(serverConn)

	room := newRoom("test")

	/* Join */
	room.join <- &join{spectator, nil} // First
	room.join <- &join{spectator, nil} // Second
	receiveMsg(t, clientConn, message{"error", []byte(`"already in room"`)})

	player, ok := room.clients[spectator]
	if !ok {
		t.Error("not in room")
		return
	}
	if player != nil {
		t.Error("not a spectator")
		return
	}

	/* Broadcast */
	room.broadcast <- &Message{"test", []byte(`"spectator"`)}
	receiveMsg(t, clientConn, message{"test", []byte(`"spectator"`)})

	/* Leave */
	room.leave <- spectator // First
	room.leave <- spectator // Second
	receiveMsg(t, clientConn, message{"error", []byte(`"not in room"`)})

	if _, ok := room.clients[spectator]; ok {
		t.Error("still in room")
	}
}

func testListenRunner(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	runner := newClient(serverConn)

	room := newRoom("test")

	/* Join */
	room.join <- &join{runner, new(game.Runner)} // First
	room.join <- &join{runner, new(game.Runner)} // Second
	receiveMsg(t, clientConn, message{"error", []byte(`"already in room"`)})

	player, ok := room.clients[runner]
	if !ok {
		t.Error("not in room")
		return
	}
	if _, ok := player.(*game.Runner); !ok {
		t.Error("not a runner")
		return
	}

	/* Broadcast */
	room.broadcast <- &Message{"test", []byte(`"runner"`)}
	receiveMsg(t, clientConn, message{"test", []byte(`"runner"`)})

	/* Leave */
	room.leave <- runner // First
	room.leave <- runner // Second
	receiveMsg(t, clientConn, message{"error", []byte(`"not in room"`)})

	if _, ok := room.clients[runner]; ok {
		t.Error("still in room")
	}
}

func testListenGuard(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	guard := newClient(serverConn)

	room := newRoom("test")

	/* Join */
	room.join <- &join{guard, new(game.Guard)} // First
	room.join <- &join{guard, new(game.Guard)} // Second
	receiveMsg(t, clientConn, message{"error", []byte(`"already in room"`)})

	player, ok := room.clients[guard]
	if !ok {
		t.Error("not in room")
		return
	}
	if _, ok := player.(*game.Guard); !ok {
		t.Error("not a guard")
		return
	}

	/* Broadcast */
	room.broadcast <- &Message{"test", []byte(`"guard"`)}
	receiveMsg(t, clientConn, message{"test", []byte(`"guard"`)})

	/* Leave */
	room.leave <- guard // First
	room.leave <- guard // Second
	receiveMsg(t, clientConn, message{"error", []byte(`"not in room"`)})

	if _, ok := room.clients[guard]; ok {
		t.Error("still in room")
	}
}

// TODO Join when game is already started
func TestListen(t *testing.T) {
	t.Run("listen", func(t *testing.T) {
		t.Run("spectator", testListenSpectator)
		t.Run("runner", testListenRunner)
		t.Run("guard", testListenGuard)
	})
}
