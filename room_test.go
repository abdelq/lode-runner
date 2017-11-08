package main

import (
	"testing"

	. "github.com/abdelq/lode-runner/game"
)

// TODO Join when game is started
func TestListen(t *testing.T) {
	room := newRoom("test")

	t.Run("spectator", func(t *testing.T) {
		serverConn, clientConn := net.Pipe()
		spectator := newClient(serverConn)

		/* Join */
		room.join <- &join{spectator, nil}
		player, ok := room.clients[spectator]
		if !ok {
			t.Fail() // TODO
			return
		}
		if player != nil {
			t.Fail() // TODO
			return
		}

		// Second join
		room.join <- &join{spectator, nil}
		receiveMsg(t, clientConn, message{"error", []byte(`"already in room"`)})

		/* Broadcast */
		room.broadcast <- &message{"test", []byte(`"spectator"`)}
		receiveMsg(t, clientConn, message{"test", []byte(`"spectator"`)})

		/* Leave */
		room.leave <- spectator
		if _, ok := room.clients[spectator]; ok {
			t.Fail() // TODO
		}

		// Second leave
		room.leave <- spectator
		receiveMsg(t, clientConn, message{"error", []byte(`"not in room"`)})
	})

	t.Run("runner", func(t *testing.T) {
		serverConn, clientConn := net.Pipe()
		runner := newClient(serverConn)

		/* Join */
		room.join <- &join{runner, new(Runner)}
		player, ok := room.clients[runner]
		if !ok {
			t.Fail() // TODO
			return
		}
		if _, ok := player.(*Runner); !ok {
			t.Fail() // TODO
			return
		}

		// Second join
		room.join <- &join{runner, new(Runner)}
		receiveMsg(t, clientConn, message{"error", []byte(`"already in room"`)})

		/* Broadcast */
		room.broadcast <- &message{"test", []byte(`"runner"`)}
		receiveMsg(t, clientConn, message{"test", []byte(`"runner"`)})

		/* Leave */
		room.leave <- runner
		if _, ok := room.clients[runner]; ok {
			t.Fail() // TODO
		}

		// Second leave
		room.leave <- runner
		receiveMsg(t, clientConn, message{"error", []byte(`"not in room"`)})
	})

	t.Run("guard", func(t *testing.T) {
		serverConn, clientConn := net.Pipe()
		guard := newClient(serverConn)

		/* Join */
		room.join <- &join{guard, new(Guard)}
		player, ok := room.clients[guard]
		if !ok {
			t.Fail() // TODO
			return
		}
		if _, ok := player.(*Guard); !ok {
			t.Fail() // TODO
			return
		}

		// Second join
		room.join <- &join{guard, new(Guard)}
		receiveMsg(t, clientConn, message{"error", []byte(`"already in room"`)})

		/* Broadcast */
		room.broadcast <- &message{"test", []byte(`"guard"`)}
		receiveMsg(t, clientConn, message{"test", []byte(`"guard"`)})

		/* Leave */
		room.leave <- guard
		if _, ok := room.clients[guard]; ok {
			t.Fail() // TODO
		}

		// Second leave
		room.leave <- guard
		receiveMsg(t, clientConn, message{"error", []byte(`"not in room"`)})
	})
}
