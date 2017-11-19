package main

import (
	"io"
	"net"
	"testing"

	msg "github.com/abdelq/lode-runner/message"
)

func TestRead(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	client := &client{conn: serverConn, out: make(chan *msg.Message, 5)} // XXX

	go client.read()

	/* TODO Tests */
	sendMsg(t, clientConn, message{})
	sendMsg(t, clientConn, message{"test", []byte(`"TestRead"`)})
}

func TestWrite(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	client := &client{conn: serverConn, out: make(chan *msg.Message, 5)} // XXX

	go client.write()

	/* TODO Tests */
	client.out <- &msg.Message{}
	receiveMsg(t, clientConn, message{"", []byte(`null`)})

	client.out <- &msg.Message{"test", []byte(`"TestWrite"`)}
	receiveMsg(t, clientConn, message{"test", []byte(`"TestWrite"`)})
}

func TestClose(t *testing.T) {
	conn, _ := net.Pipe()
	client := newClient(conn)

	// Join rooms
	for _, name := range []string{"Buzz", "Rex", "Bo"} {
		newRoom(name).join <- &join{client, nil}
	}

	client.close() // First
	client.close() // Second

	// Verify rooms are left
	for _, room := range rooms {
		if _, ok := room.clients[client]; ok {
			t.Errorf("client still in a room")
			break
		}
	}

	// XXX Verify output channel is closed
	/*if _, ok := <-client.out; ok {
		t.Error("output channel not closed")
	}*/

	// Verify connection is closed
	if _, err := conn.Read(nil); err != io.ErrClosedPipe {
		t.Error(err)
	}
}
