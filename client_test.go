package main

import (
	"io"
	"net"
	"testing"
)

func TestRead(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	client := &client{conn: serverConn, out: make(chan *message)}

	go client.read()

	// Tests
	sendMsg(t, clientConn, message{})
	sendMsg(t, clientConn, message{"", []byte(`null`)})
	sendMsg(t, clientConn, message{"test", []byte(`"TestRead"`)})
}

func TestWrite(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	client := &client{conn: serverConn, out: make(chan *message)}

	go client.write()

	// Tests
	client.out <- nil
	receiveMsg(t, clientConn, message{})
	client.out <- new(message)
	receiveMsg(t, clientConn, message{"", []byte(`null`)})
	client.out <- &message{"test", []byte(`"TestWrite"`)}
	receiveMsg(t, clientConn, message{"test", []byte(`"TestWrite"`)})
}

func TestClose(t *testing.T) {
	conn, _ := net.Pipe()
	client := newClient(conn)

	// Join rooms
	rooms := []*room{newRoom("Buzz"), newRoom("Rex"), newRoom("Bo")}
	for _, room := range rooms {
		room.join <- &join{client, nil}
	}

	client.close() // First
	client.close() // Second

	// Verify rooms are left
	for _, room := range rooms {
		if _, ok := room.clients[client]; ok {
			t.Errorf("client still in a room") // FIXME
			break
		}
	}

	// Verify output channel is closed
	if _, ok := <-client.out; ok {
		t.Error("output channel not closed")
	}

	// Verify connection is closed
	if _, err := conn.Read(nil); err != io.ErrClosedPipe {
		t.Error(err)
	}
}
