package main

import (
	"io"
	"net"
	"testing"
)

func TestClose(t *testing.T) {
	conn, _ := net.Pipe()
	client := newClient(conn)

	client.close() // First
	client.close() // Second

	// Verify output channel is closed
	if _, ok := <-client.out; ok {
		t.Error("channel not closed")
	}

	// Verify connection is closed
	if _, err := conn.Read(nil); err != io.ErrClosedPipe {
		t.Error(err)
	}
}
