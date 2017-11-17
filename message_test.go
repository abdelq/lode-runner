package main

import (
	"net"
	"testing"
)

// TODO Move sections to game package
func TestParse(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	client := newClient(serverConn)

	// Valid events
	for _, msg := range []message{{Event: "JOIN "}, {Event: " Move "}, {Event: " dig"}} {
		msg.parse(client)
		receiveMsg(t, clientConn, message{"error", []byte(`"unexpected end of JSON input"`)})
	}

	// Invalid event
	new(message).parse(client)
	receiveMsg(t, clientConn, message{"error", []byte(`"invalid event"`)})
}

func TestParseJoin(t *testing.T) {
	conn, _ := net.Pipe()
	spectator, runner := newClient(conn), newClient(conn)

	if _, ok := rooms["test"]; ok {
		t.Error("room already exists")
		return
	}

	// New room
	parseJoin([]byte(`{"name": "spectator", "room": "test", "role": 255}`), spectator)
	if _, ok := rooms["test"].clients[spectator]; !ok {
		t.Error("spectator not in room")
	}

	// Existing room
	parseJoin([]byte(`{"name": "runner", "room": "test", "role": 0}`), runner)
	if _, ok := rooms["test"].clients[runner]; !ok {
		t.Error("runner not in room")
	}
}

// TODO Move to game package
func TestParseMove(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	spectator := newClient(serverConn)

	parseMove([]byte(`{"direction": 0, "room": ""}`), spectator)
	receiveMsg(t, clientConn, message{"error", []byte(`"not in a room"`)})

	newRoom("test").clients[spectator] = nil

	parseMove([]byte(`{"direction": 0, "room": ""}`), spectator)
	receiveMsg(t, clientConn, message{"error", []byte(`"not in a game"`)})

	// TODO Not a player

	// TODO Verify player action
}

// TODO Move to game package
func TestParseDig(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	spectator := newClient(serverConn)

	parseDig([]byte(`{"direction": 0, "room": ""}`), spectator)
	receiveMsg(t, clientConn, message{"error", []byte(`"not in a room"`)})

	newRoom("test").clients[spectator] = nil

	parseDig([]byte(`{"direction": 0, "room": ""}`), spectator)
	receiveMsg(t, clientConn, message{"error", []byte(`"not in a game"`)})

	// TODO Not a runner

	// TODO Verify player action
}
