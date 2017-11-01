package main

import (
	"encoding/json"
	"net"
	"testing"
)

func TestNewRoom(t *testing.T) {
	room := newRoom("test")
	if rooms["test"] != room {
		t.Error("room not in slice")
	}
}

func newClients(t *testing.T) map[net.Conn]*client {
	t.Helper()

	clients := make(map[net.Conn]*client, 3)
	for i := 0; i < len(clients); i++ {
		serverConn, clientConn := net.Pipe()
		clients[clientConn] = newClient(serverConn)
	}
	return clients
}

func TestListen(t *testing.T) {
	room := newRoom("test")
	clients := newClients(t)

	// TODO
	t.Run("join", func(t *testing.T) {
		for _, client := range clients {
			room.join <- &join{client, nil}
		}
	})

	// TODO
	t.Run("leave", func(t *testing.T) {})

	t.Run("broadcast", func(t *testing.T) {
		msg := message{"test", json.RawMessage(`"test"`)}
		room.broadcast <- &msg

		for conn, _ := range clients {
			receiveMsg(t, conn, msg)
		}
	})
}
