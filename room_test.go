package main

import (
	"net"
	"testing"

	"github.com/abdelq/lode-runner/game"
	msg "github.com/abdelq/lode-runner/message"
)

func TestFindRoom(t *testing.T) {
	client := new(client)
	for _, name := range []string{"Buzz", "Rex", "Bo"} {
		newRoom(name)
	}

	// Not in a room
	if room := findRoom(client); room != "" {
		t.Error("room found")
	}

	// In a room
	rooms["Rex"].clients[client] = nil
	if room := findRoom(client); room != "Rex" {
		t.Error("room not found")
	}
}

func TestListen(t *testing.T) {
	t.Run("Spectator", listenSpectator)
	t.Run("Runner", listenRunner)
	t.Run("Guard", listenGuard)
}

func listenSpectator(t *testing.T) {
	t.Parallel()

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
	room.broadcast <- &msg.Message{"test", []byte(`"spectator"`)}
	receiveMsg(t, clientConn, message{"test", []byte(`"spectator"`)})
	// TODO Quit message

	/* Leave */
	room.leave <- spectator // First
	room.leave <- spectator // Second
	receiveMsg(t, clientConn, message{"error", []byte(`"not in room"`)})

	if _, ok := room.clients[spectator]; ok {
		t.Error("still in room")
	}
}

// TODO Join when game is already started
// TODO Leave when game is already stopped
// TODO Check effets of player.Add and player.Remove
func listenRunner(t *testing.T) {
	t.Parallel()

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
	room.broadcast <- &msg.Message{"test", []byte(`"runner"`)}
	receiveMsg(t, clientConn, message{"test", []byte(`"runner"`)})
	// TODO Quit message

	/* Leave */
	room.leave <- runner // First
	room.leave <- runner // Second
	receiveMsg(t, clientConn, message{"error", []byte(`"not in room"`)})

	if _, ok := room.clients[runner]; ok {
		t.Error("still in room")
	}
}

// TODO Join when game is already started
// TODO Leave when game is already stopped
// TODO Check effets of player.Add and player.Remove
func listenGuard(t *testing.T) {
	t.Parallel()

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
	room.broadcast <- &msg.Message{"test", []byte(`"guard"`)}
	receiveMsg(t, clientConn, message{"test", []byte(`"guard"`)})
	// TODO Quit message

	/* Leave */
	room.leave <- guard // First
	room.leave <- guard // Second
	receiveMsg(t, clientConn, message{"error", []byte(`"not in room"`)})

	if _, ok := room.clients[guard]; ok {
		t.Error("still in room")
	}
}
