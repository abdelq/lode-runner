package main

import (
	"net"
	"testing"
)

func TestParse(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	client := newClient(serverConn)

	// Valid events
	messages := []message{{Event: "JOIN "}, {Event: " Move "}, {Event: " dig"}}
	for _, msg := range messages {
		msg.parse(client)
		receiveMsg(t, clientConn, message{"error", []byte(`"unexpected end of JSON input"`)})
	}

	// Invalid event
	new(message).parse(client)
	receiveMsg(t, clientConn, message{"error", []byte(`"invalid event"`)})
}

func TestParseJoin(t *testing.T) {} // TODO

func TestParseMove(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	spectator := newClient(serverConn)

	parseDig([]byte(`{"direction": 0, "room": ""}`), spectator)
	receiveMsg(t, clientConn, message{"error", []byte(`"not in a room"`)})

	newRoom("test").clients[spectator] = nil // FIXME

	parseDig([]byte(`{"direction": 0, "room": "test"}`), spectator)
	receiveMsg(t, clientConn, message{"error", []byte(`"not in a game"`)})

	// TODO Not a player
}

func TestParseDig(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	spectator := newClient(serverConn)

	parseDig([]byte(`{"direction": 0, "room": ""}`), spectator)
	receiveMsg(t, clientConn, message{"error", []byte(`"not in a room"`)})

	newRoom("test").clients[spectator] = nil // FIXME

	parseDig([]byte(`{"direction": 0, "room": "test"}`), spectator)
	receiveMsg(t, clientConn, message{"error", []byte(`"not in a game"`)})

	// TODO Not a runner
}
