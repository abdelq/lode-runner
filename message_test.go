package main

import (
	"encoding/json"
	"net"
	"testing"
)

// TODO Rename
var errors = map[string]message{
	"invalidEvent":     message{"error", json.RawMessage(`"invalid event"`)},
	"invalidName":      message{"error", json.RawMessage(`"invalid name"`)},
	"invalidDirection": message{"error", json.RawMessage(`"invalid direction"`)},
	"invalidRoom":      message{"error", json.RawMessage(`"invalid room"`)},
	"notAPlayer":       message{"error", json.RawMessage(`"not a player"`)},
	"notARunner":       message{"error", json.RawMessage(`"not a runner"`)},
}

func TestParse(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	client := newClient(serverConn)

	new(message).parse(client)
	testMessageReception(t, clientConn, errors["invalidEvent"])
}

// TODO
func TestParseJoin(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	clients := map[string]*client{
		"spectator": newClient(serverConn),
		"runner":    newClient(serverConn),
		"guard":     newClient(serverConn),
	}

	// No name
	parseJoin(json.RawMessage(`{"name": "", "room": "test"}`), clients["spectator"])
	testMessageReception(t, clientConn, errors["invalidName"])

	// No room
	parseJoin(json.RawMessage(`{"name": "test", "room": ""}`), clients["spectator"])
	testMessageReception(t, clientConn, errors["invalidRoom"])

	// TODO
	/*parseJoin(json.RawMessage(`{"name": "test", "room": "test"}`), client)
	if room, ok := rooms["test"]; !ok { // New room
		t.Fail() // TODO
	} else if player, ok := room.clients[client]; !ok || player != nil { // No specific role
		t.Fail() // TODO
	}*/
	//newRoom("test")

	// Spectator
	/*parseJoin(json.RawMessage(`{"name": "spectator", "room": "test", role: 2}`), clients["spectator"])
	if player, ok := rooms["test"].clients[clients["spectator"]]; !ok || player != nil {
		t.Fail() // TODO
	}*/

	// Runner
	//parseJoin(json.RawMessage(`{"name": "runner", "room": "test", role: 0}`), clients["runner"])
	/*if _, ok := rooms["test"].clients[clients["runner"]].(*runner); !ok {
		t.Fail() // TODO
	}*/

	// Guard
	//parseJoin(json.RawMessage(`{"name": "spectator", "room": "test", role: 1}`), clients["guard"])
	/*if _, ok := rooms["test"].clients[clients["guard"]].(*guard); !ok {
		t.Fail() // TODO
	}*/
}

// TODO Improve
func TestParseMove(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	client := newClient(serverConn)

	// Inexsting room
	parseMove(json.RawMessage(`{"direction": "up", "room": "test"}`), client)

	// Existing room
	newRoom("test")

	parseMove(json.RawMessage(`{"direction": "", "room": "test"}`), client)
	testMessageReception(t, clientConn, errors["invalidDirection"])

	parseMove(json.RawMessage(`{"direction": "up", "room": ""}`), client)
	testMessageReception(t, clientConn, errors["invalidRoom"])

	parseMove(json.RawMessage(`{"direction": "up", "room": "test"}`), client)
	testMessageReception(t, clientConn, errors["notAPlayer"])

	// TODO Try joining as spectator
}

// TODO Improve
func TestParseDig(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	client := newClient(serverConn)

	newRoom("test")

	parseDig(json.RawMessage(`{"direction": "", "room": "test"}`), client)
	testMessageReception(t, clientConn, errors["invalidDirection"])

	parseDig(json.RawMessage(`{"direction": "up", "room": ""}`), client)
	testMessageReception(t, clientConn, errors["invalidRoom"])

	parseDig(json.RawMessage(`{"direction": "up", "room": "test"}`), client)
	testMessageReception(t, clientConn, errors["notARunner"])
}
