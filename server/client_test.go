package main

import (
	"io"
	"net"
	"testing"
)

func TestClose(t *testing.T) {
	conn, _ := net.Pipe()
	client := newClient(conn)

	room := newRoom("test")
	room.join <- &join{client, nil}

	client.close() // First
	client.close() // Second

	// Verify rooms are left
	room.leave <- nil // TODO
	if _, ok := room.clients[client]; ok {
		t.Fail() // TODO
	}

	// Verify output channel is closed
	if _, ok := <-client.out; ok {
		t.Fail() // TODO
	}

	// Verify connection is closed
	if _, err := conn.Read(nil); err != io.ErrClosedPipe {
		t.Error(err)
	}
}
