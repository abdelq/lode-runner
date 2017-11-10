package main

import (
	"net"
	"testing"
)

func TestParse(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	client := newClient(serverConn)

	// Valid Events
	for _, msg := range []message{
		{Event: "JOIN "}, {Event: " Move "}, {Event: " dig"},
	} {
		msg.parse(client)
		receiveMsg(t, clientConn, message{"error", []byte(`"unexpected end of JSON input"`)})
	}

	// Invalid Event
	new(message).parse(client)
	receiveMsg(t, clientConn, message{"error", []byte(`"invalid event"`)})
}

func TestParseJoin(t *testing.T) {} // TODO

func TestParseMove(t *testing.T) {} // TODO

func TestParseDig(t *testing.T) {} // TODO
