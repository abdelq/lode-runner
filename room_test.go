package main

import (
	"net"
	"testing"

	"github.com/abdelq/lode-runner/game"
	. "github.com/abdelq/lode-runner/message"
)

func TestNewRoom(t *testing.T) {} // TODO

func TestFindRoom(t *testing.T) {
	client := new(client)
	for _, name := range []string{"Buzz", "Rex", "Bo"} {
		newRoom(name)
	}

	if room := findRoom(client); room != "" {
		t.Errorf("expected: %s, receveived: %s", "", room) // TODO
	}

	rooms["Rex"].clients[client] = nil
	if room := findRoom(client); room != "Rex" {
		t.Errorf("expected: %s, receveived: %s", "Rex", room) // TODO
	}
}

// TODO Join when game is already started
// TODO Leave when game is already stopped
// TODO Check effets of player.Add and player.Remove
func TestListen(t *testing.T) {
	t.Run("listen", func(t *testing.T) {
		t.Run("spectator", listenSpectator)
		t.Run("runner", listenRunner)
		t.Run("guard", listenGuard)
	})
}

func listenSpectator(t *testing.T) {
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

func listenRunner(t *testing.T) {
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

func listenGuard(t *testing.T) {
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
