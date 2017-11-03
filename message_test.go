package main

import (
	//"encoding/json"
	//"net"
	//"reflect"
	"testing"
)

// TODO TODO And other shit down this file
/*func TestNewMessage(t *testing.T) {
	errorMsg := *newMessage("error", "test")
	expectedMsg := message{"error", json.RawMessage(`"test"`)}

	if !reflect.DeepEqual(errorMsg, expectedMsg) {
		t.Errorf("%s, want: %s", errorMsg, expectedMsg)
	}
}*/

/*func TestNewJoinMessage(t *testing.T) {
	joinMsg := *newJoinMessage("test", 2)
	expectedMsg := message{"join", json.RawMessage(
		`{"name": "test", "role": 2}`,
	)}

	if !reflect.DeepEqual(joinMsg, expectedMsg) {
		t.Errorf("%s, want: %s", joinMsg, expectedMsg)
	}
}*/

/*func TestNewLeaveMessage(t *testing.T) {
	leaveMsg := *newLeaveMessage("test", 2)
	expectedMsg := message{"leave", json.RawMessage(
		`{"name": "test", "role": 2}`,
	)}

	if !reflect.DeepEqual(leaveMsg, expectedMsg) {
		t.Errorf("%s, want: %s", leaveMsg, expectedMsg)
	}
}*/

/*func TestMessageParse(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	client := newClient(serverConn)

	msg := &message{" TEST ", nil}
	msg.parse(client)

	// TODO Comment
	if msg.Event != "test" {
		t.Errorf(`%q, want: "test"`, msg.Event)
	}

	// TODO Comment
	receiveMsg(t, clientConn, *newMessage("error", "invalid event"))
}*/

// TODO
func TestJoinMessageParse(t *testing.T) {}

// TODO
func TestParseJoin(t *testing.T) {}

// TODO
func TestParseMove(t *testing.T) {}

// TODO
func TestParseDig(t *testing.T) {}

// TODO Supposedly message{"test", nil} when received is not the same
// TODO I think I test it somewhere else..
//receiveMsg(t, conn, message{"test", json.RawMessage(`null`)})
