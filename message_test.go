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
	for _, msg := range []message{
		{Event: "JOIN "},
		{Event: " Move "},
		{Event: " dig"},
		{Event: "list"},
	} {
		msg.parse(client)
		receiveMsg(t, clientConn, message{
			"error", []byte(`"unexpected end of JSON input"`),
		})
	}

	// TODO Test the list command

	// Invalid event
	new(message).parse(client)
	receiveMsg(t, clientConn, message{"error", []byte(`"invalid event"`)})
}

func TestParseJoin(t *testing.T) {
	conn, _ := net.Pipe()
	spectator, runner := newClient(conn), newClient(conn)

	if _, ok := rooms.Load("test"); ok {
		t.Error("room already exists")
		return
	}

	// New room
	parseJoin([]byte(`{"room": "test", "role": 42}`), spectator)
	if _, ok := rooms.Load("test").(*room).clients[spectator]; !ok {
		t.Error("spectator not in room")
	}

	// Existing room
	parseJoin([]byte(`{"name": "runner", "room": "test", "level": 1}`), runner)
	if _, ok := rooms.Load("test").(*room).clients[runner]; !ok {
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
	receiveMsg(t, clientConn, message{"error", []byte(`"game not started"`)})

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
	receiveMsg(t, clientConn, message{"error", []byte(`"game not started"`)})

	// TODO Not a runner

	// TODO Verify runner action
}
